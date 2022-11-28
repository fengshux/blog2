-- auto update update_time when update
CREATE OR REPLACE FUNCTION update_modified_column()   
RETURNS TRIGGER AS $$
BEGIN
    NEW.update_time = now();
    RETURN NEW;   
END;
$$ language 'plpgsql';


CREATE TYPE user_role AS ENUM('admin', 'general');

-- Table: public.user

-- DROP TABLE IF EXISTS public."user";

CREATE TABLE IF NOT EXISTS public."user"
(
    id bigint NOT NULL DEFAULT nextval('user_id_seq'::regclass),
    username character varying(16) COLLATE pg_catalog."default" NOT NULL,
    email character varying(64) COLLATE pg_catalog."default",
    nickname character varying COLLATE pg_catalog."default" NOT NULL,
    create_time timestamp(0) with time zone NOT NULL DEFAULT now(),
    update_time timestamp(0) with time zone NOT NULL DEFAULT now(),
    role user_role NOT NULL DEFAULT 'general'::user_role,
    password character varying(128) COLLATE pg_catalog."default",
    gender gender NOT NULL DEFAULT 'unknown'::gender,
    CONSTRAINT user_pk PRIMARY KEY (id),
    CONSTRAINT user_name_unique_idx UNIQUE (username)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public."user"
    OWNER to postgres;

COMMENT ON TABLE public."user"
    IS '用户表，包含管理用户和普通用户';

COMMENT ON COLUMN public."user".id
    IS '用户id';

COMMENT ON COLUMN public."user".email
    IS '邮箱';

COMMENT ON COLUMN public."user".create_time
    IS '创建时间';

COMMENT ON COLUMN public."user".update_time
    IS '修改时间';

COMMENT ON COLUMN public."user".role
    IS '用户角色';

COMMENT ON COLUMN public."user".password
    IS 'password for login';

COMMENT ON COLUMN public."user".gender
    IS '性別';

COMMENT ON CONSTRAINT user_name_unique_idx ON public."user"
    IS 'username needs unique';

-- Trigger: update_user_update_time

-- DROP TRIGGER IF EXISTS update_user_update_time ON public."user";

CREATE TRIGGER update_user_update_time
    BEFORE UPDATE 
    ON public."user"
    FOR EACH ROW
    EXECUTE FUNCTION public.update_modified_column();
