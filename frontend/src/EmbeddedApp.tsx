/* Use of this software is governed by the Business Source License
* included in the file https://github.com/redpanda-data/redpanda/blob/dev/licenses/bsl.md
*
* As of the Change Date specified in that file, in accordance with
* the Business Source License, use of this software will be governed
* by the Apache License, Version 2.0
*/

import { useEffect } from 'react';
import {
    BrowserRouter,
} from 'react-router-dom';

/* start global stylesheets */

import 'antd/dist/antd.variable.min.css';
import './index.scss';
import './index-cloud-integration.scss';

/* end global stylesheet */


import { appGlobal } from './state/appGlobal';

import { SetConfigArguments, setup, embeddedAvailableRoutes } from './config';
import HistorySetter from './components/misc/HistorySetter';
import RequireAuth from './components/RequireAuth';
import AppContent from './components/layout/Content';
import { observer } from 'mobx-react';


export interface EmbeddedProps extends SetConfigArguments {
    // This is the base url that is used:
    //   - when making api requests
    //   - to setup the 'basename' in react-router
    //
    // In the simplest case this would be the exact url where the host is running,
    // for example "http://localhost:3001/"
    //
    // When running in cloud-ui the base most likely need to include a few more
    // things like cluster id, etc...
    // So the base would probably be "https://cloud.redpanda.com/NAMESPACE/CLUSTER/"
    //
    basePath?: string;
};

function EmbeddedApp({basePath, ...p}: EmbeddedProps) {
    useEffect(
        () => {
            const shellNavigationHandler = (event: Event) => {
                const pathname = (event as CustomEvent<string>).detail;
                const { pathname: currentPathname } = appGlobal.history.location;
                if (currentPathname === pathname || !embeddedAvailableRoutes.some((r) => r.path === pathname )) {
                    return;
                }

                appGlobal.history.push(pathname);
            };

            window.addEventListener(
                '[shell] navigated',
                shellNavigationHandler
            );

            return () => {
                window.removeEventListener(
                    '[shell] navigated',
                    shellNavigationHandler
                );
            };
        },[]);


    setup(p);

    return (
        <BrowserRouter basename={basePath}>
            <HistorySetter />
            <RequireAuth>
                <AppContent/>
            </RequireAuth>
        </BrowserRouter>
    );
}

export default observer(EmbeddedApp);
