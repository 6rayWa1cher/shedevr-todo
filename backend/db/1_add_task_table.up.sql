begin;

create schema if not exists todo;

create type todo.completed_status as enum ('yes', 'no', 'cancelled');

comment on type todo.completed_status is 'Completeness status';

create table if not exists todo.task
(
    id serial primary key,
    title text not null,
    text text,
    completed todo.completed_status not null,
    counter_value numeric,
    counter_max_value numeric,
    counter_scale text
);

comment on table todo.task is 'Generic task table';
comment on column todo.task.title is 'Title of the task';
comment on column todo.task.text is 'Text description of the task';
comment on column todo.task.completed is 'Completeness of the task';
comment on column todo.task.counter_value is 'Value of the counter of the task';
comment on column todo.task.counter_max_value is 'Max value of the counter of the task';
comment on column todo.task.counter_scale is 'Scale of the counter of the task';

commit;