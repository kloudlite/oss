import { useState } from 'react';
import { Plus, PlusFill } from '@jengaicons/react';
import { Button } from '~/components/atoms/button.jsx';
import Wrapper from '~/console/components/wrapper';
import { LoadingComp, pWrapper } from '~/console/components/loading-component';
import { useParams, useLoaderData, Link } from '@remix-run/react';
import { defer } from '@remix-run/node';
import logger from '~/root/lib/client/helpers/log';
import { GQLServerHandler } from '~/console/server/gql/saved-queries';
import {
  ensureAccountSet,
  ensureClusterSet,
} from '~/console/server/utils/auth-utils';
import {
  getPagination,
  getScopeAndProjectQuery,
  getSearch,
  parseName,
  parseNodes,
} from '~/console/server/r-urils/common';
import { parseError } from '~/root/lib/utils/common';
import ResourceList from '../../components/resource-list';
import Resources from '../_.$account.projects._index/resources';
import Tools from './tools';

const Apps = () => {
  const [viewMode, setViewMode] = useState('list');

  const { account } = useParams();
  const { promise } = useLoaderData();

  return (
    <LoadingComp data={promise}>
      {({ appsData }) => {
        const apps = parseNodes(appsData);
        if (!apps) {
          return null;
        }
        return (
          <Wrapper
            header={{
              title: 'Apps',
              action: apps.length > 0 && (
                <Button
                  variant="primary"
                  content="Create new app"
                  prefix={<PlusFill />}
                  to="../new-app"
                  LinkComponent={Link}
                />
              ),
            }}
            empty={{
              is: apps.length === 0,
              title: 'This is where you’ll manage your Apps.',
              content: (
                <p>You can create a new app and manage the listed app.</p>
              ),
              action: {
                content: 'Create new app',
                prefix: <Plus />,
                LinkComponent: Link,
                to: '../new-app',
              },
            }}
          >
            <Tools viewMode={viewMode} setViewMode={setViewMode} />
            {/* @ts-ignore */}
            <ResourceList mode={viewMode} linkComponent={Link} prefetchLink>
              {/* @ts-ignore */}
              {apps.map((app) => {
                return (
                  <ResourceList.ResourceItem
                    to={`/${account}/${app.clusterName}/${parseName(
                      app
                    )}/workspaces`}
                    key={parseName(app)}
                    textValue={parseName(app)}
                  >
                    {/* @ts-ignore */}
                    <Resources item={app} />
                  </ResourceList.ResourceItem>
                );
              })}
            </ResourceList>
          </Wrapper>
        );
      }}
    </LoadingComp>
  );
};

// @ts-ignore
export const loader = async (ctx) => {
  ensureAccountSet(ctx);
  ensureClusterSet(ctx);

  const promise = pWrapper(async () => {
    try {
      const { data, errors } = await GQLServerHandler(ctx.request).listApps({
        ...getScopeAndProjectQuery(ctx),
        pagination: getPagination(ctx),
        search: getSearch(ctx),
      });
      if (errors) {
        throw errors[0];
      }
      return { data };
    } catch (err) {
      logger.error(err);
      return { error: parseError(err).message };
    }
  });

  return defer({ promise });
};

export default Apps;
