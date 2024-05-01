/* eslint-disable react/destructuring-assignment */
import Popup from '~/components/molecule/popup';
import { toast } from '~/components/molecule/toast';
import CommonPopupHandle from '~/console/components/common-popup-handle';
import { NameIdView } from '~/console/components/name-id-view';
import { IDialogBase } from '~/console/components/types.d';
import { useConsoleApi } from '~/console/server/gql/api-provider';
import { IEnvironments } from '~/console/server/gql/queries/environment-queries';
import {
  ExtractNodeType,
  parseName,
  parseNodes,
} from '~/console/server/r-utils/common';
import { useReload } from '~/root/lib/client/helpers/reloader';
import useForm, { dummyEvent } from '~/root/lib/client/hooks/use-form';
import Yup from '~/root/lib/server/helpers/yup';
import { handleError } from '~/root/lib/utils/common';
import useCustomSwr from '~/root/lib/client/hooks/use-custom-swr';
import Select from '~/components/atoms/select';

type IDialog = IDialogBase<ExtractNodeType<IEnvironments>>;

const Root = (props: IDialog) => {
  const { isUpdate, setVisible } = props;
  const api = useConsoleApi();
  const reloadPage = useReload();

  const { data: clustersData, isLoading: cIsLoading } = useCustomSwr(
    'clusters',
    async () => api.listClusters({}),
    true
  );

  const { values, errors, handleChange, handleSubmit, resetValues, isLoading } =
    useForm({
      initialValues: {
        name: '',
        displayName: '',
        environmentRoutingMode: false,
        isNameError: false,
        clusterName: '',
      },
      validationSchema: Yup.object({
        name: Yup.string().required('Name is required.'),
        displayName: Yup.string().required(),
        clusterName: Yup.string().required(),
      }),
      onSubmit: async (val) => {
        if (isUpdate) {
          try {
            const { errors: e } = await api.cloneEnvironment({
              displayName: val.displayName,
              environmentRoutingMode: val.environmentRoutingMode
                ? 'public'
                : 'private',
              destinationEnvName: val.name,

              sourceEnvName: parseName(props.data),
            });
            if (e) {
              throw e[0];
            }
            resetValues();
            toast.success('Environment cloned successfully');
            setVisible(false);
            reloadPage();
          } catch (err) {
            handleError(err);
          }
        }
      },
    });
  return (
    <Popup.Form
      onSubmit={(e) => {
        if (!values.isNameError) {
          handleSubmit(e);
        } else {
          e.preventDefault();
        }
      }}
    >
      <Popup.Content>
        <div className="flex flex-col gap-3xl">
          <NameIdView
            displayName={values.displayName}
            name={values.name}
            resType="environment"
            errors={errors.name}
            label="Name"
            placeholder="Environment name"
            handleChange={handleChange}
            nameErrorLabel="isNameError"
          />

          <Select
            label="Select Cluster"
            size="lg"
            value={values.clusterName}
            disabled={cIsLoading}
            placeholder="Select a Cluster"
            options={async () => [
              ...((clustersData &&
                parseNodes(clustersData)
                  .filter((d) => {
                    return d.status?.isReady;
                  })
                  .map((d) => ({
                    label: `${d.displayName} - ${parseName(d)}`,
                    value: parseName(d),
                  }))) ||
                []),
            ]}
            onChange={({ value }) => {
              handleChange('clusterName')(dummyEvent(value));
            }}
            error={!!errors.clusterName}
            message={errors.clusterName}
          />

          {/* <Checkbox */}
          {/*   label="Public" */}
          {/*   checked={values.environmentRoutingMode} */}
          {/*   onChange={(val) => { */}
          {/*     handleChange('environmentRoutingMode')(dummyEvent(val)); */}
          {/*   }} */}
          {/* /> */}
          {/* <Banner */}
          {/*   type="info" */}
          {/*   body={ */}
          {/*     <span> */}
          {/*       Public environments will expose services to the public internet. */}
          {/*       Private environments will be accessible when Kloudlite VPN is */}
          {/*       active. */}
          {/*     </span> */}
          {/*   } */}
          {/* /> */}
        </div>
      </Popup.Content>
      <Popup.Footer>
        <Popup.Button closable content="Cancel" variant="basic" />
        <Popup.Button
          type="submit"
          content="Clone"
          variant="primary"
          loading={isLoading}
        />
      </Popup.Footer>
    </Popup.Form>
  );
};

const CloneEnvironment = (props: IDialog) => {
  return (
    <CommonPopupHandle
      {...props}
      root={Root}
      updateTitle="Clone Environment"
      createTitle="Clone Environment"
    />
  );
};

export default CloneEnvironment;
