do
$$
    begin
        execute 'ALTER DATABASE ' || current_database() || ' SET timezone = ''UTC''';
    end;
$$;

create table token
(
    value       text     not null
        primary key,
    created_at  timestamptz not null default now(),
    usr_id      bigint      not null,
    platform_id smallint    not null default 0
);

create index token_usr_id_idx
    on token (usr_id);
