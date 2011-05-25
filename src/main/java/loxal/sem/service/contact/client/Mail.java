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

package loxal.sem.service.contact.client;

import com.google.gwt.core.client.EntryPoint;
import com.google.gwt.core.client.GWT;
import com.google.gwt.event.dom.client.ClickEvent;
import com.google.gwt.uibinder.client.UiBinder;
import com.google.gwt.uibinder.client.UiField;
import com.google.gwt.uibinder.client.UiHandler;
import com.google.gwt.user.client.rpc.AsyncCallback;
import com.google.gwt.user.client.ui.*;

/**
 * @author Alexander Orlov <alexander.orlov@loxal.net>
 */
public class Mail implements EntryPoint {
    @UiField
    TextBox sender;
    @UiField
    TextBox subject;
    @UiField
    TextArea message;
    @UiField
    Button send;
    @UiField
    Button sendNewMessage;
    @UiField
    DecoratedPopupPanel actionResult;

    @Override
    public void onModuleLoad() {
    }

    interface MailUiBinder extends UiBinder<Widget, Mail> {
    }

    private Mail() {
        MailUiBinder binder = GWT.create(MailUiBinder.class);
        Widget app = binder.createAndBindUi(this);

        send.setAccessKey('S');
        sendNewMessage.setAccessKey('B');
        subject.setFocus(true);

        RootLayoutPanel.get().add(app);
    }

    private void displayActionResult(String msg, boolean success) {
        actionResult.clear();
        actionResult.add(new HTML(msg));
        actionResult.setStyleName(success ? "success" : "failure", true);
        actionResult.center();
        actionResult.show();
    }

    private void sendMail() {
    }

    @UiHandler("sendNewMessage")
    void sendNewMessage(ClickEvent event) {
        sender.setReadOnly(false);
        subject.setReadOnly(false);
        message.setReadOnly(false);
        send.setVisible(true);
        sendNewMessage.setVisible(false);
    }

    @UiHandler("send")
    void send(ClickEvent event) {
        sendMail();
    }
}
