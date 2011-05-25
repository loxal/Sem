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

package loxal.lox.meta.server;

import com.google.appengine.api.users.User;
import com.google.appengine.api.users.UserService;
import com.google.appengine.api.users.UserServiceFactory;
import com.google.gwt.user.server.rpc.RemoteServiceServlet;
import loxal.lox.service.meta.client.meta.authentication.AuthInfo;
import loxal.lox.service.meta.client.meta.authentication.AuthInfoSvc;

/**
 * @author Alexander Orlov <alexander.orlov@loxal.net>
 */
public class AuthInfoSvcImpl extends RemoteServiceServlet implements AuthInfoSvc {
    @Override
    public AuthInfo getAuthInfo(String requestUri) {
        UserService userService = UserServiceFactory.getUserService();
        User user = userService.getCurrentUser();
        AuthInfo auth = new AuthInfo();

        if (user != null) {
            auth.setLoggedIn(true);
            auth.setAdmin(userService.isUserAdmin());
            auth.setEmail(user.getEmail());
            auth.setNickname(user.getNickname());
            auth.setLogoutUrl(userService.createLogoutURL(requestUri));
        } else {
            auth.setLoggedIn(false);
            auth.setLoginURL(userService.createLoginURL(requestUri));
        }
        return auth;
    }
}
