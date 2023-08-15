import { Button } from '~/components/atoms/button';
import { ArrowLeft, ArrowRight } from '@jengaicons/react';
import { TextInput } from '~/components/atoms/input';
import { BrandLogo } from '~/components/branding/brand-logo';
import { ProgressTracker } from '~/components/organisms/progress-tracker';
import { useState } from 'react';
import {
  useParams,
  useLoaderData,
  useOutletContext,
  useNavigate,
} from '@remix-run/react';
import SelectInput from '~/components/atoms/select';
import useForm from '~/root/lib/client/hooks/use-form';
import Yup from '~/root/lib/server/helpers/yup';
import logger from '~/root/lib/client/helpers/log';
import { toast } from '~/components/molecule/toast';
import { useAPIClient } from '~/root/lib/client/hooks/api-provider';
import { Select } from '~/components/atoms/select-new';
import { Badge } from '~/components/atoms/badge';
import { cn } from '~/components/utils';
import { IdSelector, idTypes } from '../components/id-selector';
import {
  getCluster,
  getClusterSepc,
  getCredentialsRef,
} from '../server/r-urils/cluster';
import {
  getMetadata,
  getPagination,
  parseName,
} from '../server/r-urils/common';
import { GQLServerHandler } from '../server/gql/saved-queries';
import { keyconstants } from '../server/r-urils/key-constants';
import { constDatas } from '../dummy/consts';
import { ensureAccountSet } from '../server/utils/auth-utils';
import AlertDialog from '../components/alert-dialog';
import RawWrapper from '../components/raw-wrapper';

const NewCluster = () => {
  const [isOnboarding, setIsOnboarding] = useState(false);
  const [showUnsavedChanges, setShowUnsavedChanges] = useState(false);
  const api = useAPIClient();

  const { providerSecrets } = useLoaderData();
  const cloudProviders = providerSecrets?.edges?.map(({ node }) => node) || [];

  const { a: account } = useParams();
  // @ts-ignore
  const { user } = useOutletContext();

  const navigate = useNavigate();

  const [selectedProvider, setSelectedProvider] = useState();

  const { values, errors, handleSubmit, handleChange, isLoading } = useForm({
    initialValues: {
      vpc: '',
      name: '',
      region: 'ap-south-1',
      cloudProvider: '',
      credentialsRef: '',
      availabilityMode: '',
      displayName: '',
    },
    validationSchema: Yup.object({
      vpc: Yup.string(),
      region: Yup.string().trim().required('region is required'),
      cloudProvider: Yup.string().trim().required('cloud provider is required'),
      name: Yup.string().trim().required('id is required'),
      displayName: Yup.string().trim().required('name is required'),
      credentialsRef: Yup.string().required(),
      availabilityMode: Yup.string()
        .trim()
        .required('availability is required'),
    }),
    onSubmit: async (val) => {
      try {
        const { errors: e } = await api.createCluster({
          cluster: getCluster({
            spec: getClusterSepc({
              accountName: account,
              vpc: val.vpc || undefined,
              region: val.region,
              cloudProvider: val.cloudProvider,
              credentialsRef: getCredentialsRef({
                name: val.credentialsRef,
              }),
              availabilityMode: val.availabilityMode,
            }),
            metadata: getMetadata({
              name: val.name,
              annotations: {
                [keyconstants.displayName]: val.displayName,
                [keyconstants.author]: user.name,
              },
            }),
          }),
        });
        if (e) {
          throw e[0];
        }
        toast.success('cluster created successfully');
        navigate(`/${account}/clusters`);
      } catch (err) {
        toast.error(err.message);
      }
    },
  });

  return (
    <>
      <RawWrapper
        leftChildren={
          <>
            <BrandLogo detailed={false} size={48} />
            <div
              className={cn('flex flex-col', {
                'gap-4xl': isOnboarding,
                'gap-8xl': !isOnboarding,
              })}
            >
              <div className="flex flex-col gap-3xl">
                <div className="text-text-default heading4xl">
                  {isOnboarding
                    ? "Unleash Data's Full Potential!"
                    : 'Let’s create new cluster.'}
                </div>
                <div className="text-text-default bodyMd">
                  {isOnboarding
                    ? 'Kloudlite will help you to develop and deploy cloud native applications easily.'
                    : 'Create your cluster under to production effortlessly'}
                </div>
                {!isOnboarding && (
                  <div className="flex flex-row gap-md items-center">
                    <Badge>
                      <span className="text-text-strong">Account:</span>
                      <span className="bodySm-semibold text-text-default">
                        Kloudlite Labs Pvt Ltd
                      </span>
                    </Badge>
                  </div>
                )}
              </div>
              <ProgressTracker
                items={
                  isOnboarding
                    ? [
                        { label: 'Create Team', active: true, id: 1 },
                        {
                          label: 'Invite your Team Members',
                          active: true,
                          id: 2,
                        },
                        {
                          label: 'Add your Cloud Provider',
                          active: true,
                          id: 3,
                        },
                        { label: 'Setup First Cluster', active: true, id: 4 },
                        { label: 'Create your project', active: false, id: 5 },
                      ]
                    : [
                        { label: 'Configure cluster', active: true, id: 1 },
                        { label: 'Review', active: false, id: 2 },
                      ]
                }
              />
            </div>
            {isOnboarding && (
              <Button variant="outline" content="Skip" size="lg" />
            )}
            {!isOnboarding && (
              <Button
                variant="outline"
                content="Cancel"
                size="lg"
                onClick={() => setShowUnsavedChanges(true)}
              />
            )}
          </>
        }
        rightChildren={
          <form onSubmit={handleSubmit} className="flex flex-col gap-3xl">
            <div className="text-text-soft headingLg">Cluster details</div>
            {Object.keys(JSON.parse(JSON.stringify(errors || '{}')) || {})
              .length > 0 && (
              <pre className="text-xs text-surface-warning-default">
                <code>{JSON.stringify(errors, null, 2)}</code>
              </pre>
            )}
            <TextInput
              label="Cluster name"
              onChange={handleChange('displayName')}
              value={values.displayName}
              error={!!errors.displayName}
              message={errors.displayName}
              size="lg"
            />
            <IdSelector
              resType={idTypes.cluster}
              name={values.displayName}
              onChange={(v) => {
                handleChange('name')({ target: { value: v } });
              }}
            />

            <Select
              id="cloudprovider-select"
              label="Cloud Provider"
              size="lg"
              value={{
                value: parseName(selectedProvider),
                label: parseName(selectedProvider),
                provider: selectedProvider,
              }}
              options={(cloudProviders || []).map((provider) => ({
                value: parseName(provider),
                label: parseName(provider),
                provider,
              }))}
              onChange={({ provider }) => {
                handleChange('credentialsRef')({
                  target: { value: parseName(provider) },
                });
                handleChange('cloudProvider')({
                  target: { value: provider?.cloudProviderName },
                });
                setSelectedProvider(provider);
              }}
            />

            <SelectInput.Root
              label="Region"
              value={values.region}
              size="lg"
              onChange={(v) => {
                handleChange('region')({ target: { value: v } });
              }}
            >
              <SelectInput.Option> -- not-selected -- </SelectInput.Option>
              {constDatas.regions.map(({ name, value }) => {
                return (
                  <SelectInput.Option key={value} value={value}>
                    {name}
                  </SelectInput.Option>
                );
              })}
            </SelectInput.Root>

            <SelectInput.Root
              label="Availabilty"
              size="lg"
              value={values.availabilityMode}
              onChange={(v) => {
                handleChange('availabilityMode')({ target: { value: v } });
              }}
            >
              <SelectInput.Option> -- not-selected -- </SelectInput.Option>
              {constDatas.availabilityModes.map(({ name, value }) => {
                return (
                  <SelectInput.Option key={value} value={value}>
                    {name}
                  </SelectInput.Option>
                );
              })}
            </SelectInput.Root>

            <TextInput
              label="VPC"
              size="lg"
              onChange={handleChange('vpc')}
              value={values.vpc}
              error={!!errors.vpc}
              message={errors.vpc}
            />
            {isOnboarding ? (
              <div className="flex flex-row gap-xl justify-end">
                <Button
                  variant="outline"
                  content="Back"
                  prefix={ArrowLeft}
                  size="lg"
                />
                <Button
                  variant="primary"
                  content="Continue"
                  suffix={ArrowRight}
                  size="lg"
                  type="submit"
                />
              </div>
            ) : (
              <div className="flex flex-row justify-end">
                <Button
                  loading={isLoading}
                  variant="primary"
                  content="Create"
                  suffix={ArrowRight}
                  type="submit"
                  size="lg"
                />
              </div>
            )}
          </form>
        }
      />
      <AlertDialog
        title="Leave page with unsaved changes?"
        message="Leaving this page will delete all unsaved changes."
        okText="Delete"
        type="critical"
        show={showUnsavedChanges}
        setShow={setShowUnsavedChanges}
        onSubmit={() => {
          setShowUnsavedChanges(false);
          navigate(`/${account}/clusters`);
        }}
      />
    </>
  );
};

export const loader = async (ctx) => {
  ensureAccountSet(ctx);
  const { data, errors } = await GQLServerHandler(
    ctx.request
  ).listProviderSecrets({
    pagination: getPagination(ctx),
  });

  if (errors) {
    logger.error(errors);
  }

  return {
    providerSecrets: data,
  };
};

export default NewCluster;
