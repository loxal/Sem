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

package loxal.sem.widget.commander.client;

import com.google.gwt.core.client.GWT;
import com.google.gwt.event.dom.client.ClickEvent;
import com.google.gwt.event.dom.client.ClickHandler;
import com.google.gwt.http.client.*;
import com.google.gwt.json.client.JSONArray;
import com.google.gwt.json.client.JSONObject;
import com.google.gwt.json.client.JSONParser;
import com.google.gwt.json.client.JSONValue;
import com.google.gwt.uibinder.client.UiBinder;
import com.google.gwt.uibinder.client.UiField;
import com.google.gwt.user.client.ui.*;
import com.google.gwt.xhr.client.XMLHttpRequest;

/**
 * Commander UI Logic
 */
public class Commander extends Composite {
    @UiField
    TextBox name;
    //    @UiField
//    TextArea desc;
    @UiField
    SubmitButton create;
    @UiField
    TabLayoutPanel tabPanel;
    @UiField
    FormPanel cmdCreator;
    @UiField
    VerticalPanel formContainer;
    @UiField
    VerticalPanel container;
    @UiField
    TextArea desc;
    @UiField
    TextBox restCall;

    interface Binder extends UiBinder<Widget, Commander> {
    }

    public Commander() {
        Binder binder = GWT.create(Binder.class);
        initWidget(binder.createAndBindUi(this));

        XMLHttpRequest xmlHttpRequest;
//        xmlHttpRequest.open("PUT", "http://localhost:8080/create?name=gwtMUMMMM&desc=gwturl&restCall=gwtrest");
//        xmlHttpRequest.open("GET", "http://localhost:8080/cmdList");
//        GWT.log(xmlHttpRequest.getStatusText());
//        GWT.log(xmlHttpRequest.getAllResponseHeaders());

        create.setAccessKey('C');

//        http://code.google.com/p/google-web-toolkit-doc-1-5/wiki/GettingStartedJSON

//
//        container.add(formm);
    }

    public static final String jsonUrl = "http://localhost:8080/cmdList?name=";

    public void cmdCreate() {
        {
            String url = URL.encode(jsonUrl);

            // parse the response text into JSON
            JSONValue jsonValue = JSONParser.parseStrict("{\"blu\": \"blab\"}");
            JSONValue jsonValue1 = JSONParser.parseStrict("{\"blu\": \"blab\"}");
            JSONArray jsonArray = jsonValue.isArray();
            JSONObject jsonObject = new JSONObject();


            RequestBuilder requestBuilder = new RequestBuilder(RequestBuilder.GET, "url");
            try {
                Request request = requestBuilder.sendRequest(null, new RequestCallback() {
                    @Override
                    public void onResponseReceived(Request request, Response response) {
                    }

                    @Override
                    public void onError(Request request, Throwable exception) {
                    }
                });
            } catch (RequestException e) {
                e.printStackTrace();
            }


            jsonObject.put("myKey", jsonValue);

            GWT.log(jsonObject.toString());
            GWT.log(jsonObject.get("myKey").toString());
//        GWT.log(jsonValue1.isString().toString());
//        GWT.log(jsonValue1.isArray().toString());
            GWT.log(jsonValue1.isObject().toString());

        }
    }
}
