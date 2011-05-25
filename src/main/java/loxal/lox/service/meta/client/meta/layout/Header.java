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

import com.google.gwt.core.client.GWT;
import com.google.gwt.event.dom.client.ChangeEvent;
import com.google.gwt.http.client.UrlBuilder;
import com.google.gwt.i18n.client.LocaleInfo;
import com.google.gwt.uibinder.client.UiBinder;
import com.google.gwt.uibinder.client.UiField;
import com.google.gwt.uibinder.client.UiHandler;
import com.google.gwt.user.client.Window;
import com.google.gwt.user.client.ui.*;
import loxal.lox.service.meta.client.meta.authentication.AuthInfo;

public class Header extends Composite {
    interface Binder extends UiBinder<Widget, Header> {
    }

    public void setAuthenticationInfo(AuthInfo authInfo) {
        this.authInfo = authInfo;
    }

    private AuthInfo authInfo;

    protected Header() {
        Binder binder = GWT.create(Binder.class);
        initWidget(binder.createAndBindUi(this));
        localeSwitch();
    }

    @UiField
    ListBox localeSwitch;
    @UiField
    Anchor authenticationLink;
    @UiField
    Frame sponsor;
    @UiField
    Frame script;
    @UiField
    AbsolutePanel test;
    @UiField
    static DecoratedPopupPanel actionResult;

    private I18n i18n = GWT.create(I18n.class);

    public static void displayActionResult(String msg, boolean success) {
        actionResult.clear();
        actionResult.add(new HTML(msg));
        actionResult.setStyleName(success ? "success" : "failure", true);
        actionResult.center();
        actionResult.show();
    }

    void localeSwitch() {
        localeSwitch.setAccessKey('L');
        localeSwitch.setTabIndex(0);
        localeSwitch.setTitle("Choose your language" + " [Access Key: L]");
        localeSwitch.setFocus(true);
        String currentLocale = LocaleInfo.getCurrentLocale().getLocaleName().equals("default") ? "en" : LocaleInfo.getCurrentLocale().getLocaleName();
        String[] localeNames = LocaleInfo.getAvailableLocaleNames();
        localeSwitch.addItem("English", "en");
        for (String localeName : localeNames) {
            if (!localeName.equals("default")) {
                String localeNative = LocaleInfo.getLocaleNativeDisplayName(localeName);
                localeSwitch.addItem(localeNative, localeName);
                if (localeName.equals(currentLocale)) {
                    localeSwitch.setSelectedIndex(localeSwitch.getItemCount() - 1);
                }
            }
        }
    }

    @UiHandler("localeSwitch")
    void onChange(ChangeEvent event) {
        String localeName = localeSwitch.getValue(localeSwitch.getSelectedIndex());
        UrlBuilder builder = Window.Location.createUrlBuilder().setParameter("locale",
                localeName);
        Window.Location.replace(builder.buildString());
    }

    void authentication() {
        authenticationLink.setAccessKey('A');
        authenticationLink.setTitle("[Access Key: A]");
        authenticationLink.setTabIndex(1);
        if (authInfo.isLoggedIn()) {
            authenticationLink.setText(i18n.signOut() + ": " + authInfo.getEmail());
            authenticationLink.setHref(authInfo.getLogoutURL());
            authenticationLink.setTitle("Nickname: " + authInfo.getNickname() + (authInfo.isAdmin() ? " (Admin)" : "") + " [Access Key: A]");
        } else {
            authenticationLink.setHref(authInfo.getLoginURL());
            authenticationLink.setText(i18n.signIn());
            authenticationLink.setFocus(true);
        }
    }
}