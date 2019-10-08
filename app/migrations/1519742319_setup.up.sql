

-- DROP TABLE IF EXISTS public.income CASCADE;
CREATE TABLE public.tasks(
	id serial NOT NULL,
    name varchar(30) NOT NULL,
    time timestamp NOT NULL DEFAULT LOCALTIMESTAMP,
	CONSTRAINT pk_task_id PRIMARY KEY (id)

);


