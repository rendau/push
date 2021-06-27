alter table usr_push_token
add column platform_id integer not null default 0;

alter table usr_push_token
rename usr_id to usr_id_old;

alter table usr_push_token
add usr_id bigint;

update usr_push_token
set usr_id = usr_id_old::bigint;

alter table usr_push_token
alter column usr_id set not null;

alter table usr_push_token
drop column if exists usr_id_old cascade;

create index usr_push_token_usr_id_idx on usr_push_token (usr_id);

create index usr_push_token_platform_idx on usr_push_token (platform_id);