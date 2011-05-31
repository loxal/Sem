/*
 * Copyright 2011 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
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

        RequestBuilder requestBuilder = new RequestBuilder(RequestBuilder.GET, "/cmd/list.json");
        try {
            Request request = requestBuilder.sendRequest(null, new RequestCallback() {
                @Override
                public void onResponseReceived(Request request, Response response) {
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

    public static final String jsonUrl = "http://localhost:8080/cmd/list.json";
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
