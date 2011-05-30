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
import com.google.gwt.http.client.*;
import com.google.gwt.json.client.JSONArray;
import com.google.gwt.json.client.JSONObject;
import com.google.gwt.json.client.JSONParser;
import com.google.gwt.json.client.JSONValue;
import com.google.gwt.uibinder.client.UiBinder;
import com.google.gwt.uibinder.client.UiField;
import com.google.gwt.uibinder.client.UiHandler;
import com.google.gwt.user.client.ui.*;
import com.google.gwt.xhr.client.XMLHttpRequest;

/**
 * Commander UI Logic
 */
public class Commander extends Composite {
    @UiField
    TextBox name;
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
//        xmlHttpRequest.open("GET", "http://localhost:8080/cmdList?json=true");
//        GWT.log(xmlHttpRequest.getStatusText());
//        GWT.log(xmlHttpRequest.getAllResponseHeaders());

        create.setAccessKey('C');

//        http://code.google.com/p/google-web-toolkit-doc-1-5/wiki/GettingStartedJSON

//
//        container.add(formm);


        RequestBuilder requestBuilder = new RequestBuilder(RequestBuilder.GET, "http://localhost:8080/cmdList?json=true");
//        RequestBuilder requestBuilder = new RequestBuilder(RequestBuilder.GET, "http://127.0.0.1:8889/commander/index.html?gwt.codesvr=127.0.0.1:9998");
        try {
            Request request = requestBuilder.sendRequest(null, new RequestCallback() {
                @Override
                public void onResponseReceived(Request request, Response response) {
                    GWT.log(response.getText());
                    GWT.log(response.getStatusText());
                    GWT.log(String.valueOf(response.getStatusCode()));
                    GWT.log(response.getText());
                }

                @Override
                public void onError(Request request, Throwable exception) {
                    GWT.log("RequestBuilder: error");
                }
            });
        } catch (RequestException e) {
            e.printStackTrace();
        }

//        try {
//            Request r = requestBuilder.send();
////            GWT.log(String.valueOf(r.isPending()));
//        } catch (RequestException e) {
//            e.printStackTrace();
//        }

    }

    public static final String jsonUrl = "http://localhost:8080/cmdList?json=true";
    public static final String jsonUrl1 = GWT.getModuleBaseURL() + "cmdList?name=";
    public static final String jsonUrl2 = GWT.getHostPageBaseURL() + "cmdList?name=";

    public void cmdCreation() {
        {
            String url = URL.encode(jsonUrl);
            GWT.log(jsonUrl2);

            // parse the response text into JSON
            JSONValue jsonValue = JSONParser.parseStrict("{\"blu\": \"blab\"}");
//            JSONValue jsonValue1 = JSONParser.parseStrict("{\"blab\"}");
            JSONArray jsonArray = jsonValue.isArray();
            JSONObject jsonObject = new JSONObject();


            GWT.log(jsonObject.toString());
//            GWT.log(jsonObject.get("myKey").toString());
//        GWT.log(jsonValue1.isString().toString());
//        GWT.log(jsonValue1.isArray().toString());
//            GWT.log(jsonValue1.isObject().toString());

        }
    }

    @UiHandler("create")
    void cmdCreate(ClickEvent event) {
        cmdCreation();
    }
}
