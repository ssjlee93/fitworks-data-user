-- roles table
-- roles as constants
drop table if exists public.roles;

create table public.roles
(
    role_id serial primary key,
    role    varchar(64)                         not null,
    created timestamp default CURRENT_TIMESTAMP not null,
    updated timestamp default CURRENT_TIMESTAMP not null
);

-- add roles
insert into public.roles (role)
values ('trainer'),
       ('client'),
       ('admin');