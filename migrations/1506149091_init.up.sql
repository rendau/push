do $$
begin
  execute 'ALTER DATABASE ' || current_database() || ' SET timezone = ''UTC''';
end;
$$;

create table usr_push_token (
  value      varchar     not null,
  usr_id     varchar     not null,
  created_at timestamptz not null default now(),

  constraint usr_push_token_pk primary key (value)
);

create index usr_push_token_usr_id_idx
  on usr_push_token (usr_id);