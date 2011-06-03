/*
 * Copyright 2011 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package loxal.sem.widget.commander.client;

import com.google.gwt.cell.client.TextCell;
import com.google.gwt.core.client.GWT;
import com.google.gwt.event.dom.client.ClickEvent;
import com.google.gwt.http.client.*;
import com.google.gwt.json.client.*;
import com.google.gwt.uibinder.client.UiBinder;
import com.google.gwt.uibinder.client.UiField;
import com.google.gwt.uibinder.client.UiHandler;
import com.google.gwt.user.cellview.client.CellList;
import com.google.gwt.user.client.ui.*;
import com.google.gwt.xhr.client.XMLHttpRequest;

import java.util.Arrays;
import java.util.List;

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
    @UiField
    Label content;

    interface Binder extends UiBinder<Widget, Commander> {
    }

    public Commander() {
        Binder binder = GWT.create(Binder.class);
        initWidget(binder.createAndBindUi(this));

        XMLHttpRequest xmlHttpRequest;

        create.setAccessKey('C');

//        http://code.google.com/p/google-web-toolkit-doc-1-5/wiki/GettingStartedJSON
//        container.add(formm);

        RequestBuilder requestBuilder = new RequestBuilder(RequestBuilder.GET, "/cmd/list.json");
        try {
            Request request = requestBuilder.sendRequest(null, new RequestCallback() {
                @Override
                public void onResponseReceived(Request request, Response response) {
                    List<String> DAYS = Arrays.asList(response.getText().split("},"));
                    // Create a cell to render each value in the list.
                    TextCell textCell = new TextCell();
                    // Create a CellList that uses the cell.
                    CellList<String> cellList = new CellList<String>(textCell);
                    // Set the total row count. This isn't strictly necessary, but it affects
                    // paging calculations, so its good habit to keep the row count up to date.
                    cellList.setRowCount(DAYS.size(), true);
                    // Push the data into the widget.
                    cellList.setRowData(0, DAYS);
                    container.add(cellList);

                    {
                        getJSONValue(response.getText(), "Name");
                    }
                }

                @Override
                public void onError(Request request, Throwable exception) {
                    GWT.log("RequestBuilder: error" + exception);
                }
            });
        } catch (RequestException e) {
            e.printStackTrace();
        }
    }

    private String getJSONValue(String jsonString, String valueName) {
        JSONValue jsonValue = JSONParser.parseStrict(jsonString);
        JSONObject jsonObject = jsonValue.isObject();
        JSONValue jsonValueCmds = jsonObject.get("cmds");
        JSONArray jsonArray = jsonValueCmds.isArray();
        JSONValue jsonValueCmdsRow = jsonArray.get(2);
        JSONObject jsonObjectRow = jsonValueCmdsRow.isObject();
        JSONValue jsonValueCmdsRowValue = jsonObjectRow.get(valueName);
        JSONString jsonValueString = jsonValueCmdsRowValue.isString();
        content.setText(jsonValueString.stringValue());

        return jsonValueString.stringValue();
    }

    public void cmdCreation() {
        {


        }
    }

    @UiHandler("create")
    void cmdCreate(ClickEvent event) {
        cmdCreation();
    }
}
