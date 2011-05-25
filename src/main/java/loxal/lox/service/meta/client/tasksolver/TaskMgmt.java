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

package loxal.lox.service.meta.client.tasksolver;

import com.google.gwt.cell.client.ButtonCell;
import com.google.gwt.cell.client.FieldUpdater;
import com.google.gwt.core.client.GWT;
import com.google.gwt.event.dom.client.*;
import com.google.gwt.event.logical.shared.SelectionEvent;
import com.google.gwt.event.logical.shared.SelectionHandler;
import com.google.gwt.event.logical.shared.ValueChangeEvent;
import com.google.gwt.event.logical.shared.ValueChangeHandler;
import com.google.gwt.uibinder.client.UiBinder;
import com.google.gwt.uibinder.client.UiField;
import com.google.gwt.uibinder.client.UiHandler;
import com.google.gwt.user.cellview.client.*;
import com.google.gwt.user.client.Command;
import com.google.gwt.user.client.rpc.AsyncCallback;
import com.google.gwt.user.client.ui.*;
import com.google.gwt.view.client.ListDataProvider;
import com.google.gwt.view.client.MultiSelectionModel;
import com.google.gwt.view.client.SelectionModel;
import loxal.lox.service.meta.client.dto.Task;
import loxal.lox.service.meta.client.meta.layout.Header;

import java.util.ArrayList;
import java.util.List;

/**
 * Task UI logic
 */
public class TaskMgmt extends Composite {
    private TaskSvcAsync taskSvcAsync = GWT.create(TaskSvc.class);

    interface Binder extends UiBinder<Widget, TaskMgmt> {
    }

    @UiField
    TextBox name;
    @UiField
    TextArea description;
    @UiField
    Button createTask;
    @UiField
    TextBox category;
    @UiField
    TextBox priority;
    @UiField
    TabLayoutPanel tabPanel;
    @UiField
    CellTable<Task> taskPager;
    @UiField
    VerticalPanel control;
    @UiField
    MenuBar menu;
    @UiField
    MenuItem taskItem;
    @UiField
    MenuItem deleteTask;
    @UiField
    VerticalPanel processTaskPanel;
    @UiField
    MenuItem closeTask;
    @UiField
    TextBox nameUpdate;
    @UiField
    TextBox categoryUpdate;
    @UiField
    TextBox priorityUpdate;
    @UiField
    TextArea descriptionUpdate;
    @UiField
    Button updateTask;
    @UiField
    Label taskId;
    @UiField
    Button showAllTasks;
    @UiField
    MenuItem placeholder;
    @UiField
    PageSizePager taskPageSizePager;
    @UiField
    SimplePager taskSimplePager;

    public TaskMgmt() {
        Binder binder = GWT.create(Binder.class);
        initWidget(binder.createAndBindUi(this));

        createTask.setAccessKey('C');
        updateTask.setAccessKey('U');
        tabPanel.selectTab(0);

        getTasks();

        closeTask.setCommand(new Command() {
            @Override
            public void execute() {
                tabPanel.selectTab(1);
            }
        });

        tabPanel.addSelectionHandler(new SelectionHandler<Integer>() {
            @Override
            public void onSelection(
                    SelectionEvent<Integer> integerSelectionEvent) {
            }
        });

    }

    private void loadTask(String taskId) {
        taskSvcAsync.getTask(taskId, new AsyncCallback<Task>() {
            @Override
            public void onFailure(Throwable caught) {
            }

            @Override
            public void onSuccess(Task task) {
                displayTask(task);
            }
        });
    }

    private void displayTask(Task task) {
        nameUpdate.setText(task.getName());
        categoryUpdate.setText(task.getCategory());
        priorityUpdate.setText(String.valueOf(task.getPriority()));
        descriptionUpdate.setText(task.getDescription());
        taskId.setText(task.getId());
    }

    private void getTasks() {
        taskSvcAsync.getTasks(new AsyncCallback<ArrayList<Task>>() {
            @Override
            public void onFailure(Throwable caught) {
            }

            @Override
            public void onSuccess(ArrayList<Task> tasks) {
                displayTasks(tasks);

                { // SuggestBox / Oracle
                    MultiWordSuggestOracle oracle = new MultiWordSuggestOracle();
                    for (Task task : tasks) {
                        oracle.add(task.getName());
                    }

                    tabPanel.remove(3);
                    TextBox searchBox = new TextBox();
                    final SuggestBox search = new SuggestBox(oracle, searchBox); // UiBinder variant didn't work; also using the PROVIDED attribute

                    searchBox.addFocusHandler(new FocusHandler() {
                        @Override
                        public void onFocus(FocusEvent focusEvent) {
                            search.setText("");
                        }
                    });

                    searchBox.addBlurHandler(new BlurHandler() {
                        @Override
                        public void onBlur(BlurEvent blurEvent) {
                            search.setText("Search For a Task Name");
                        }
                    });

                    search.setText("Search For a Task Name");
                    search.setWidth("12em");
                    search.setAccessKey('O');
                    search.setTitle("[Access Key: O]");
                    search.setFocus(true);

                    search.addValueChangeHandler(new ValueChangeHandler<String>() {
                        @Override
                        public void onValueChange(
                                ValueChangeEvent<String> stringValueChangeEvent) {
                            taskSvcAsync.searchTasksWithName(
                                    stringValueChangeEvent.getValue(),
                                    new AsyncCallback<ArrayList<Task>>() {
                                        @Override
                                        public void onFailure(
                                                Throwable caught) {
                                        }

                                        @Override
                                        public void onSuccess(
                                                ArrayList<Task> tasks) {
                                            displayTasks(tasks);
                                            tabPanel.selectTab(1);
                                            showAllTasks.setVisible(true);
                                        }
                                    });
                        }
                    });

                    tabPanel.add(new HTML(), search);
                }
            }
        });
    }

    private boolean initConstruction; // TODO re-engineer this HACK: make this var

    private void displayTasks(ArrayList<Task> tasks) {
        ListDataProvider<Task> listDataProvider = new ListDataProvider<Task>();
        listDataProvider.addDataDisplay(taskPager);
        taskPageSizePager.setDisplay(taskPager);
        taskSimplePager.setDisplay(taskPager);
        SelectionModel<Task> selectionModel = new MultiSelectionModel<Task>(); // TODO not yet working (because the API isn't ready?)
        taskPager.setSelectionModel(selectionModel);

        ArrayList<Task> taskDTOs = new ArrayList<Task>();
        for (Task task : tasks) {
            taskDTOs.add(task);
        }
        listDataProvider.setList(taskDTOs);

        if (!initConstruction) {
            initTableColumnsOfTasks();
            initConstruction = true;
        }
    }

    private void initTableColumnsOfTasks() {
        taskPager.addColumn(new TextColumn<Task>() {
                    @Override
                    public String getValue(Task object) {
                        return object.getId();
                    }
                }, "ID");

        taskPager.addColumn(new TextColumn<Task>() {
                    @Override
                    public String getValue(Task object) {
                        return object.getName();
                    }
                }, "Name");

        taskPager.addColumn(new TextColumn<Task>() {
                    @Override
                    public String getValue(Task object) {
                        return object.getCategory();
                    }
                }, "Category");

        taskPager.addColumn(new TextColumn<Task>() {
                    @Override
                    public String getValue(Task object) {
                        return String.valueOf(object.getPriority());
                    }
                }, "Priority");

        taskPager.addColumn(new TextColumn<Task>() {
                    @Override
                    public String getValue(Task object) {
                        return object.getDescription();
                    }
                }, "Description");

        Column<Task, String> edit = new Column<Task, String>(
                new ButtonCell()) {
            @Override
            public String getValue(Task object) {
                return "Edit";
            }
        };
        edit.setFieldUpdater(new FieldUpdater<Task, String>() {
            @Override
            public void update(int index, final Task object,
                               String value) {
                tabPanel.selectTab(3);
                loadTask(object.getId());
                taskItem.setHTML("Task " + object.getId());
                deleteTask.setCommand(new Command() {
                    @Override
                    public void execute() {
                        ArrayList<String> taskIds = new ArrayList<String>();
                        taskIds.add(object.getId());
                        deleteTasks(taskIds);
                        tabPanel.selectTab(1);
                    }
                });
            }
        });
        taskPager.addColumn(edit);

        Column<Task, String> removeButton = new Column<Task, String>(
                new ButtonCell()) {
            @Override
            public String getValue(Task object) {
                return "X";
            }
        };
        removeButton.setFieldUpdater(new FieldUpdater<Task, String>() {
            @Override
            public void update(int index, Task object,
                               String value) {
                List<String> selectedTaskIds = new ArrayList<String>();
                selectedTaskIds.add(object.getId());
                deleteTasks(selectedTaskIds);
                // deleteTasks(selectedTaskIds); // TODO native Longs caused
                // compilation errors in Scala implemented server-side class
            }
        });
        taskPager.addColumn(removeButton);

    }

    private Task declareTask() {
        Task task = new Task();
        task.setName(name.getValue());
        task.setDescription(description.getValue());
        task.setCategory(category.getValue());
        if (priority.getValue().matches("\\d+"))
            task.setPriority(Integer.parseInt(priority.getValue()));

        return task;
    }

    private Task declareTaskUpdate() {
        Task task = new Task();
        task.setName(nameUpdate.getValue());
        task.setDescription(descriptionUpdate.getValue());
        task.setCategory(categoryUpdate.getValue());
        if (priorityUpdate.getValue().matches("\\d+"))
            task.setPriority(Integer.parseInt(priorityUpdate.getValue()));
        task.setId(taskId.getText());

        return task;
    }

    private void putTask(Task task) {
        taskSvcAsync.putTask(task, new AsyncCallback<Void>() {
            @Override
            public void onFailure(Throwable caught) {
                Header.displayActionResult(
                        "Error Creating Task: " + caught.getMessage(), false);
            }

            @Override
            public void onSuccess(Void ignore) {
                Header.displayActionResult("Task Created Successfully", true);
                getTasks();
            }
        });
    }

    private void deleteTasks(List<String> selectedTaskIds) {
        ArrayList<String> tasks = new ArrayList<String>(); // TODO
        // optimize
        // > don't
        // create
        // this
        // redundant
        // object,
        // use the
        // Task IDs
        // only
        for (String taskId : selectedTaskIds) {
            tasks.add(taskId);
        }

        taskSvcAsync.deleteTasks(tasks, new AsyncCallback<Void>() {
            @Override
            public void onFailure(Throwable caught) {
            }

            @Override
            public void onSuccess(Void result) {
                getTasks();
            }
        });
    }

    private void updateTask(Task task) {
        taskSvcAsync.updateTask(task, new AsyncCallback<Void>() {
            @Override
            public void onFailure(Throwable caught) {
                Header.displayActionResult(
                        "Error Updating Task: " + caught.getMessage(), false);
            }

            @Override
            public void onSuccess(Void ignore) {
                Header.displayActionResult("Task Updated Successfully", true);
                getTasks();
            }
        });
    }

    @UiHandler("createTask")
    void addTask(ClickEvent event) {
        putTask(declareTask());
        name.setFocus(true);
        tabPanel.selectTab(1);
    }

    @UiHandler("updateTask")
    void addTaskUpdate(ClickEvent event) {
        updateTask(declareTaskUpdate());
        tabPanel.selectTab(1);
    }

    @UiHandler("showAllTasks")
    void showAllTasks(ClickEvent event) {
        showAllTasks.setVisible(false);
        getTasks();
    }
}
