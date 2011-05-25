/*
 * Copyright 2011 Alexander Orlov <alexander.orlov@loxal.net>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package loxal.lox.service.meta.client.meta.layout;

import com.google.gwt.core.client.EntryPoint;
import com.google.gwt.core.client.GWT;
import com.google.gwt.uibinder.client.UiBinder;
import com.google.gwt.uibinder.client.UiField;
import com.google.gwt.user.client.Window;
import com.google.gwt.user.client.rpc.AsyncCallback;
import com.google.gwt.user.client.ui.DockLayoutPanel;
import com.google.gwt.user.client.ui.RootLayoutPanel;
import com.google.gwt.user.client.ui.TabLayoutPanel;
import com.google.gwt.user.client.ui.Widget;
import loxal.lox.service.meta.client.meta.authentication.AuthInfo;
import loxal.lox.service.meta.client.meta.authentication.AuthInfoSvc;
import loxal.lox.service.meta.client.meta.authentication.AuthInfoSvcAsync;
import loxal.lox.service.meta.client.tasksolver.TaskMgmt;

public class Entry implements EntryPoint {
    interface Binder extends UiBinder<Widget, Entry> {
    }

    private Binder binder = GWT.create(Binder.class);
    private I18n i18n = GWT.create(I18n.class);
    private AuthInfoSvcAsync authInfoSvcAsync = GWT.create(AuthInfoSvc.class);

    @UiField
    TaskMgmt taskMgmt;
    @UiField
    Footer footer;
    @UiField
    Header header;
    @UiField
    TabLayoutPanel taskTab;

    @Override
    public void onModuleLoad() {
        Widget app = binder.createAndBindUi(this);
        RootLayoutPanel.get().add(app);
        Window.setTitle(i18n.appTitle());

        authInfoSvcAsync.getAuthInfo(GWT.getHostPageBaseURL(),
                new AsyncCallback<AuthInfo>() {
                    @Override
                    public void onFailure(Throwable caught) {
                    }

                    @Override
                    public void onSuccess(AuthInfo authInfo) {
                        header.setAuthenticationInfo(authInfo);
                        header.authentication();
                        taskTab.selectTab(0);
                    }
                });
    }
}
