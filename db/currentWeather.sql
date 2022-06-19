create table location(
   location_id serial primary key,
   sonya_flakeID varchar(20) UNIQUE NOT NULL,
   city_name VARCHAR (100) not null,
   country VARCHAR (100) not null,
   lat numeric (5, 2) not null,
   lon numeric (5, 2) not null,
   local_time timestamp not null
);

create table current_temp(
   city_id int not null,
   sonya_flakeID varchar(20) UNIQUE NOT NULL,
   temp_in_C numeric (5, 2) not null,
   condition_text VARCHAR (50) not null,
   primary key(city_id)
);

create table air_quality(
   city_id int not null,
   sonya_flakeID varchar(20) UNIQUE NOT NULL,
   uv numeric (4, 2) not null,
   co numeric (5, 2) not null,
   no2 numeric (5, 2) not null,
   o3 numeric (5, 2) not null,
   so2 numeric (5, 2) not null,
   pm2_5 numeric (5, 2) not null,
   pm10 numeric (5, 2) not null,
   us_epa_index int not null,
   gb_defra_index int not null,
   primary key(city_id)
);