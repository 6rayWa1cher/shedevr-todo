begin;

alter table todo.task
    add if not exists counter_exist bool not null default false;

comment on column todo.task.counter_exist is 'Is counter present for this task';

update todo.task
set counter_exist = true
where counter_scale is not null
   or counter_max_value is not null
   or counter_value is not null;

commit;