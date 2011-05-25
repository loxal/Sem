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

package loxal.lox.service.meta.server;

import com.google.appengine.api.datastore.*;
import com.google.appengine.api.users.User;
import com.google.appengine.api.users.UserServiceFactory;
import com.google.gwt.user.server.rpc.RemoteServiceServlet;
import loxal.lox.service.meta.client.dto.Task;
import loxal.lox.service.meta.client.tasksolver.TaskSvc;

import java.util.ArrayList;
import java.util.Iterator;

/**
 * @author Alexander Orlov <alexander.orlov@loxal.net>
 */
public class TaskSvcImpl extends RemoteServiceServlet implements TaskSvc {
    private DatastoreService datastoreService = DatastoreServiceFactory.getDatastoreService();
    private static String entityName = "Task";

    @Override
    public void putTask(Task task) {
        Entity entity = new Entity(entityName);
        entity.setProperty("name", task.getName());
        entity.setProperty("category", task.getCategory());
        entity.setProperty("description", task.getDescription());
        entity.setProperty("priority", task.getPriority());
        User user = UserServiceFactory.getUserService().getCurrentUser();
        if (user != null) entity.setProperty("userEmail", user.getEmail());
        else entity.setProperty("userEmail", null); // null must be set explicitly

        datastoreService.put(entity);
    }

    @Override
    public void deleteTasks(ArrayList<String> entityKeys) { // TODO check if these tasks actually belong to the user
        ArrayList<Key> keys = new ArrayList<Key>();

        for (String key : entityKeys) {
            keys.add(KeyFactory.createKey(entityName, Long.parseLong(key)));
        }

        datastoreService.delete(keys);
    }

    @Override
    public void updateTask(Task task) { // TODO check if this task actually belongs to the user
        Entity taskEntity;
        try {
            taskEntity = datastoreService.get(KeyFactory.createKey(entityName, task.getId()));

            taskEntity.setProperty("name", task.getName());
            taskEntity.setProperty("category", task.getCategory());
            taskEntity.setProperty("description", task.getDescription());
            taskEntity.setProperty("priority", task.getPriority());

            datastoreService.put(taskEntity);
        } catch (EntityNotFoundException e) {
            e.printStackTrace();
        }
    }

    @Override
    public ArrayList<Task> searchTasksWithName(String taskName) {
        ArrayList<Task> tasks = new ArrayList<Task>();

        Query taskQuery = new Query(entityName);
        String userEmail =
                UserServiceFactory.getUserService().getCurrentUser() == null ? null : UserServiceFactory.getUserService().getCurrentUser().getEmail();

        taskQuery.addFilter("userEmail", Query.FilterOperator.EQUAL, userEmail);
        taskQuery.addFilter("name", Query.FilterOperator.EQUAL, taskName);
        Iterator<Entity> taskQueryResult = datastoreService.prepare(taskQuery).asIterator();
        while (taskQueryResult.hasNext()) {
            Entity taskEntity = taskQueryResult.next();
            Task task = new Task();

            task.setName(String.valueOf(taskEntity.getProperty("name")));
            task.setCategory(String.valueOf(taskEntity.getProperty("category")));
            task.setDescription(String.valueOf(taskEntity.getProperty("description")));
            task.setPriority(Integer.parseInt(taskEntity.getProperty("priority").toString()));
            task.setUserEmail(String.valueOf(taskEntity.getProperty("userEmail")));
            task.setId(String.valueOf(taskEntity.getKey().getId()));

            tasks.add(task);
        }
        return tasks;
    }

    @Override
    public Task getTask(String taskId) {
        try {
            Entity taskEntity = datastoreService.get(KeyFactory.createKey(entityName, Long.parseLong(taskId)));

            Task task = new Task();
            String userEmail = UserServiceFactory.getUserService().getCurrentUser() == null ? null : UserServiceFactory.getUserService().getCurrentUser().getEmail();
            if (taskEntity.getProperty("userEmail") == null || taskEntity.getProperty("userEmail").equals(userEmail)) { // check for NPE first 
                task.setUserEmail(String.valueOf(taskEntity.getProperty("userEmail")));
                task.setName(String.valueOf(taskEntity.getProperty("name")));
                task.setCategory(String.valueOf(taskEntity.getProperty("category")));
                task.setDescription(String.valueOf(taskEntity.getProperty("description")));
                task.setPriority(Integer.parseInt(taskEntity.getProperty("priority").toString()));
                task.setId(String.valueOf(taskEntity.getKey().getId()));
            }

            return task;
        } catch (EntityNotFoundException e) {
            e.printStackTrace();
        }
        return null;
    }

    @Override
    public ArrayList<Task> getTasks() {
        DatastoreService datastoreService = DatastoreServiceFactory.getDatastoreService();
        ArrayList<Task> tasks = new ArrayList<Task>();
        Query taskQuery = new Query(entityName);

        String userEmail = UserServiceFactory.getUserService().getCurrentUser() == null ? null : UserServiceFactory.getUserService().getCurrentUser().getEmail();
        taskQuery.addFilter("userEmail", Query.FilterOperator.EQUAL, userEmail); // to assure that only user-owned entities are fetched
        Iterator<Entity> taskQueryResult = datastoreService.prepare(taskQuery).asIterator();
        while (taskQueryResult.hasNext()) {
            Entity taskEntity = taskQueryResult.next();

            Task task = new Task();
            task.setName(String.valueOf(taskEntity.getProperty("name")));
            task.setCategory(String.valueOf(taskEntity.getProperty("category")));
            task.setDescription(String.valueOf(taskEntity.getProperty("description")));
            task.setPriority(Integer.parseInt(taskEntity.getProperty("priority").toString()));
            task.setUserEmail(String.valueOf(taskEntity.getProperty("userEmail")));
            task.setId(String.valueOf(taskEntity.getKey().getId()));

            tasks.add(task);
        }

        return tasks;
    }

}
