/*
 * Copyright 2011 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

package loxal.sem.widget.commander.client;

import com.google.gwt.core.client.EntryPoint;
import com.google.gwt.core.client.GWT;
import com.google.gwt.uibinder.client.UiBinder;
import com.google.gwt.uibinder.client.UiField;
import com.google.gwt.user.client.ui.RootLayoutPanel;
import com.google.gwt.user.client.ui.Widget;

public class Entry implements EntryPoint {
    interface Binder extends UiBinder<Widget, Entry> {
    }

    private Binder binder = GWT.create(Binder.class);

    @UiField
    Commander commander;

    @Override
    public void onModuleLoad() {
        Widget app = binder.createAndBindUi(this);
        RootLayoutPanel.get().add(app);
    }
}
