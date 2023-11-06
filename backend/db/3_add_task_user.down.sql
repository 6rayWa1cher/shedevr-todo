drop index if exists todo.task.ix_todo_task_user_id;

alter table todo.task
    drop column if exists user_id;