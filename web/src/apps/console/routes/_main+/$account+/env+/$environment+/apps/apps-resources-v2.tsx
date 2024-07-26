import {
  GearSix,
  LinkBreak,
  Link as LinkIcon,
  Repeat,
} from '~/console/components/icons';
import { Link, useOutletContext, useParams } from '@remix-run/react';
import { generateKey, titleCase } from '~/components/utils';
import {
  ListItem,
  ListItemV2,
  ListTitle,
  ListTitleV2,
  listClass,
} from '~/console/components/console-list-components';
import Grid from '~/console/components/grid';
import ListGridView from '~/console/components/list-grid-view';
import ResourceExtraAction, {
  IResourceExtraItem,
} from '~/console/components/resource-extra-action';
import { useConsoleApi } from '~/console/server/gql/api-provider';
import { IApps } from '~/console/server/gql/queries/app-queries';
import {
  ExtractNodeType,
  parseName,
  parseName as pn,
  parseUpdateOrCreatedBy,
  parseUpdateOrCreatedOn,
} from '~/console/server/r-utils/common';
import { handleError } from '~/lib/utils/common';
import { toast } from '~/components/molecule/toast';
import { useReload } from '~/lib/client/helpers/reloader';
import { SyncStatusV2 } from '~/console/components/sync-status';
import { useWatchReload } from '~/lib/client/helpers/socket/useWatch';
import ListV2 from '~/console/components/listV2';
import { useState } from 'react';
import { Badge } from '~/components/atoms/badge';
import { CopyContentToClipboard } from '~/console/components/common-console-components';
import Tooltip from '~/components/atoms/tooltip';
import { NN } from '~/root/lib/types/common';
import { getClusterStatus } from '~/console/utils/commons';
import HandleIntercept from './handle-intercept';
import { IEnvironmentContext } from '../_layout';
import TooltipV2 from '~/components/atoms/tooltipV2';

const RESOURCE_NAME = 'app';
type BaseType = ExtractNodeType<IApps>;

const parseItem = (item: ExtractNodeType<IApps>) => {
  return {
    name: item.displayName,
    id: pn(item),
    intercept: item.spec.intercept,
    updateInfo: {
      author: `Updated by ${titleCase(parseUpdateOrCreatedBy(item))}`,
      time: parseUpdateOrCreatedOn(item),
    },
  };
};

type OnAction = ({
  action,
  item,
}: {
  action: 'delete' | 'edit' | 'intercept' | 'remove_intercept' | 'restart';
  item: BaseType;
}) => void;

type IExtraButton = {
  onAction: OnAction;
  item: BaseType;
};

const InterceptPortView = ({
  ports = [],
  devName = '',
}: {
  ports: NN<ExtractNodeType<IApps>['spec']['intercept']>['portMappings'];
  devName: string;
}) => {
  return (
    <div className="flex flex-row items-center gap-md pulsable">
      <TooltipV2
        content={
          <div>
            <span className="bodyMd-medium text-text-soft">
              Intercepted to{' '}
              <span className="bodyMd-medium text-text-strong">{devName}</span>
            </span>
            <div className="flex flex-row gap-md py-md">
              {ports?.map((d) => {
                return (
                  <Badge className="shrink-0" key={d.appPort}>
                    <div>
                      {d.appPort} → {d.devicePort}
                    </div>
                  </Badge>
                );
              })}
            </div>
          </div>
        }
      >
        <div className="bodyMd-medium text-text-strong w-fit truncate">
          {ports?.length === 1 ? (
            <span>{ports.length} port</span>
          ) : (
            <span>{ports.length} ports</span>
          )}
          <span className="text-text-soft">
            {' '}
            intercepted to{' '}
            <span className="bodyMd-medium text-text-strong truncate">
              {devName}
            </span>
          </span>
        </div>
      </TooltipV2>
    </div>
  );
};

const ExtraButton = ({ onAction, item }: IExtraButton) => {
  const { account, environment } = useParams();
  const iconSize = 16;
  let options: IResourceExtraItem[] = [
    {
      label: 'Settings',
      icon: <GearSix size={iconSize} />,
      type: 'item',
      to: `/${account}/env/${environment}/app/${parseName(
        item,
      )}/settings/general`,
      key: 'settings',
    },
  ];

  if (item.spec.intercept && item.spec.intercept.enabled) {
    options = [
      {
        label: 'Remove intercept',
        icon: <LinkBreak size={iconSize} />,
        type: 'item',
        onClick: () => onAction({ action: 'remove_intercept', item }),
        key: 'remove-intercept',
      },
      ...options,
    ];
  } else {
    options = [
      {
        label: 'Intercept',
        icon: <Repeat size={iconSize} />,
        type: 'item',
        onClick: () => onAction({ action: 'intercept', item }),
        key: 'intercept',
      },
      ...options,
    ];
  }

  options = [
    {
      label: 'Restart',
      icon: <LinkIcon size={iconSize} />,
      type: 'item',
      onClick: () => onAction({ action: 'restart', item }),
      key: 'restart',
    },
    ...options,
  ];

  return <ResourceExtraAction options={options} />;
};

interface IResource {
  items: BaseType[];
  onAction: OnAction;
}

const AppServiceView = ({ service }: { service: string }) => {
  return (
    <CopyContentToClipboard
      toolTip
      content={service}
      toastMessage="App service url copied successfully."
    />
  );
};

const GridView = ({ items = [], onAction: _ }: IResource) => {
  const { account, environment } = useParams();
  return (
    <Grid.Root className="!grid-cols-1 md:!grid-cols-3" linkComponent={Link}>
      {items.map((item, index) => {
        const { name, id, updateInfo } = parseItem(item);
        const keyPrefix = `${RESOURCE_NAME}-${id}-${index}`;
        return (
          <Grid.Column
            key={id}
            to={`/${account}/env/${environment}/app/${id}`}
            rows={[
              {
                key: generateKey(keyPrefix, name + id),
                render: () => (
                  <ListTitle
                    title={name}
                    subtitle={id}
                    action={
                      <ResourceExtraAction
                        options={[
                          {
                            key: 'apps-resource-extra-action-1',
                            to: `/${account}/env/${environment}/app/${id}/settings/general`,
                            icon: <GearSix size={16} />,
                            label: 'Settings',
                            type: 'item',
                          },
                        ]}
                      />
                    }
                  />
                ),
              },
              {
                key: generateKey(keyPrefix, updateInfo.author),
                render: () => (
                  <ListItem
                    data={updateInfo.author}
                    subtitle={updateInfo.time}
                  />
                ),
              },
            ]}
          />
        );
      })}
    </Grid.Root>
  );
};

const ListView = ({ items = [], onAction }: IResource) => {
  const { environment, account, cluster } =
    useOutletContext<IEnvironmentContext>();

  return (
    <ListV2.Root
      linkComponent={Link}
      data={{
        headers: [
          {
            render: () => 'Name',
            name: 'name',
            className: listClass.title,
          },
          {
            render: () => '',
            name: 'intercept',
            className: 'w-[250px] truncate',
          },
          {
            render: () => '',
            name: 'flex-pre',
            className: listClass.flex,
          },
          {
            render: () => 'Service',
            name: 'service',
            className: 'w-[240px] flex',
          },
          {
            render: () => '',
            name: 'flex-post',
            className: listClass.flex,
          },
          {
            render: () => 'Status',
            name: 'status',
            className: 'min-w-[120px]',
          },
          {
            render: () => '',
            name: 'flex-post',
            className: listClass.flex,
          },
          {
            render: () => 'Updated',
            name: 'updated',
            className: listClass.updated,
          },
          {
            render: () => '',
            name: 'action',
            className: listClass.action,
          },
        ],
        rows: items.map((i) => {
          const isClusterOnline = getClusterStatus(cluster);

          const { name, id, updateInfo } = parseItem(i);
          return {
            columns: {
              name: {
                render: () => <ListTitleV2 title={name} subtitle={id} />,
              },
              intercept: {
                render: () =>
                  i.spec.intercept?.enabled ? (
                    <div>
                      <InterceptPortView
                        ports={i.spec.intercept.portMappings || []}
                        devName={i.spec.intercept.toDevice || ''}
                      />
                    </div>
                  ) : null,
              },
              service: {
                render: () => <AppServiceView service={i.serviceHost || ''} />,
              },
              status: {
                render: () => (
                  <div className="inline-block">
                    {isClusterOnline ? (
                      <SyncStatusV2 item={i} />
                    ) : (
                      <Badge type="warning">Cluster Offline</Badge>
                    )}
                  </div>
                ),
              },
              updated: {
                render: () => (
                  <ListItemV2
                    data={`${updateInfo.author}`}
                    subtitle={updateInfo.time}
                  />
                ),
              },
              action: {
                render: () => <ExtraButton onAction={onAction} item={i} />,
              },
            },
            to: `/${parseName(account)}/env/${parseName(
              environment,
            )}/app/${id}`,
          };
        }),
      }}
    />
  );
};

const AppsResourcesV2 = ({ items = [] }: Omit<IResource, 'onAction'>) => {
  const api = useConsoleApi();
  const { environment, account } = useOutletContext<IEnvironmentContext>();
  const reload = useReload();

  const [visible, setVisible] = useState(false);
  const [mi, setItem] = useState<ExtractNodeType<IApps>>();

  useWatchReload(
    items.map((i) => {
      return `account:${parseName(account)}.environment:${parseName(
        environment,
      )}.app:${parseName(i)}`;
    }),
  );

  const interceptApp = async (item: BaseType, intercept: boolean) => {
    if (intercept) {
      setItem(item);
      setVisible(true);
      return;
    }

    try {
      const { errors } = await api.interceptApp({
        appname: pn(item),
        deviceName: item.spec.intercept?.toDevice || '',
        envName: pn(environment),
        intercept,
      });

      if (errors) {
        throw errors[0];
      }
      // toast.success('app intercepted successfully');
      toast.success(
        `${
          intercept
            ? 'App Intercepted successfully'
            : 'App Intercept removed successfully'
        }`,
      );
      reload();
    } catch (error) {
      handleError(error);
    }
  };

  const restartApp = async (item: BaseType) => {
    if (!environment) {
      throw new Error('Environment is required!.');
    }

    try {
      const { errors } = await api.restartApp({
        appName: pn(item),
        envName: pn(environment),
      });

      if (errors) {
        throw errors[0];
      }
      toast.success('App restarted successfully');
      // reload();
    } catch (error) {
      handleError(error);
    }
  };

  const props: IResource = {
    items,
    onAction: ({ action, item }) => {
      switch (action) {
        case 'intercept':
          interceptApp(item, true);
          break;
        case 'restart':
          restartApp(item);
          break;
        case 'remove_intercept':
          interceptApp(item, false);
          break;
        default:
      }
    },
  };
  return (
    <>
      <ListGridView
        listView={<ListView {...props} />}
        gridView={<GridView {...props} />}
      />
      <HandleIntercept
        {...{
          visible,
          setVisible,
          app: mi,
        }}
      />
    </>
  );
};

export default AppsResourcesV2;
