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

import loxal.lox.service.meta.client.dto.Task;

import java.util.ArrayList;

import com.google.gwt.user.client.rpc.AsyncCallback;

public interface TaskSvcAsync {
    void putTask(Task task, AsyncCallback<Void> callback);

    void deleteTasks(ArrayList<String> entityKeys, AsyncCallback<Void> callback);

    void getTasks(AsyncCallback<ArrayList<Task>> callback);

    void getTask(String taskId, AsyncCallback<Task> callback);

    void updateTask(Task task, AsyncCallback<Void> callback);

    void searchTasksWithName(String taskName, AsyncCallback<ArrayList<Task>> callback);
}