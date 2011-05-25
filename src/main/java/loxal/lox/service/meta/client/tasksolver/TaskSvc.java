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

import com.google.gwt.user.client.rpc.RemoteService;
import com.google.gwt.user.client.rpc.RemoteServiceRelativePath;
import loxal.lox.service.meta.client.dto.Task;

import java.util.ArrayList;

@RemoteServiceRelativePath("Task")
public interface TaskSvc extends RemoteService {
    void putTask(Task task);

    void deleteTasks(ArrayList<String> entityKeys);

    ArrayList<Task> getTasks();

    Task getTask(String taskId);

    void updateTask(Task task);

    ArrayList<Task> searchTasksWithName(String taskName);
}