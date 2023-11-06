alter table todo.task
    add column if not exists user_id text not null default '';

alter table todo.task
    alter column user_id drop default;

comment on column todo.task.user_id is 'Task creator user id';

create index if not exists ix_todo_task_user_id on todo.task (user_id);
