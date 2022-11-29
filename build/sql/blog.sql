-- auto update update_time when update
CREATE OR REPLACE FUNCTION update_modified_column()   
RETURNS TRIGGER AS $$
BEGIN
    NEW.update_time = now();
    RETURN NEW;   
END;
$$ language 'plpgsql';


CREATE TYPE user_role AS ENUM('admin', 'general');
CREATE TYPE public.gender AS ENUM ('unknown', 'female', 'male');
COMMENT ON TYPE public.gender IS '性别';

-- Table: public.user

-- DROP TABLE IF EXISTS public."user";

CREATE TABLE public."user" (
	id bigserial NOT NULL, -- 用户id
	username varchar(16) NOT NULL,
	email varchar(64) NULL, -- 邮箱
	nickname varchar NOT NULL,
	"role" public.user_role NOT NULL DEFAULT 'general'::user_role, -- 用户角色
	gender public.gender NOT NULL DEFAULT 'unknown'::gender, -- 性別
	"password" varchar(128) NULL,
	create_time timestamptz(0) NOT NULL DEFAULT now(), -- 创建时间
	update_time timestamptz(0) NOT NULL DEFAULT now(), -- 修改时间        
	CONSTRAINT user_name_unique_idx UNIQUE (username),
	CONSTRAINT user_pk PRIMARY KEY (id)
);
COMMENT ON TABLE public."user" IS '用户表，包含管理用户和普通用户';

-- Column comments

COMMENT ON COLUMN public."user".id IS '用户id';
COMMENT ON COLUMN public."user".email IS '邮箱';
COMMENT ON COLUMN public."user".create_time IS '创建时间';
COMMENT ON COLUMN public."user".update_time IS '修改时间';
COMMENT ON COLUMN public."user"."role" IS '用户角色';
COMMENT ON COLUMN public."user".gender IS '性別';

-- Constraint comments

COMMENT ON CONSTRAINT user_name_unique_idx ON public."user" IS 'username needs unique';

-- Table Triggers

create trigger update_user_update_time before
update
    on
    public."user" for each row execute function update_modified_column();
