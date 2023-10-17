begin;

create type public.completed_status as enum ('yes', 'no', 'cancelled');

comment on type public.completed_status is 'Completeness status';

create table if not exists public.task (
    id serial primary key,
    title text not null,
    text text,
    completed public.completed_status not null,
    counter_value numeric,
    counter_max_value numeric,
    counter_scale text
);

comment on table public.task is 'Generic task table';
comment on column public.task.title is 'Title of the task';
comment on column public.task.text is 'Text description of the task';
comment on column public.task.completed is 'Completeness of the task';
comment on column public.task.counter_value is 'Value of the counter of the task';
comment on column public.task.counter_max_value is 'Max value of the counter of the task';
comment on column public.task.counter_scale is 'Scale of the counter of the task';

commit;