CREATE TABLE orders (
    id integer not null GENERATED ALWAYS AS IDENTITY,
	OrderUID         varchar,
	TrackNumber       varchar,
	Entry   varchar,	
	Locale            varchar,	
	InternalSignature varchar,	
	CustomerID        varchar,	
	DeliveryService   varchar,	
	ShardKey         varchar,	
	SMID              integer,
	DateCreated       timestamp without time zone,
	OOFShard    varchar,
	PRIMARY KEY(id),
	CONSTRAINT	TrackNumberunique UNIQUE (TrackNumber),
	CONSTRAINT	orderUIDunique UNIQUE (orderUID)
	);

CREATE TABLE orderitems (
    id integer not null GENERATED ALWAYS AS IDENTITY,
	ChrtID         varchar,
	TrackNumber       varchar,
	Price integer,
	RID   varchar,	
	Name            varchar,
	Sale integer,
	Size varchar,	
	TotalPrice integer,
	NMID integer,
	Brand        varchar,	
	Status integer,
	PRIMARY KEY(id),
	CONSTRAINT fk_tracknumber
   		FOREIGN KEY(TrackNumber) 
      		REFERENCES orders(TrackNumber)
	);

CREATE TABLE payment (
    id integer not null GENERATED ALWAYS AS IDENTITY,
	Transaction         varchar,
	RequestID       varchar,
	Currency varchar,
	Provider   varchar,	
	Amount integer,
	PaymentDT integer,	
	Bank        varchar,	
	DeliveryCost integer,
	GoodsTotal integer,
	CustomFee integer,	
	PRIMARY KEY(id),
	CONSTRAINT fk_transaction
   		FOREIGN KEY(Transaction) 
      		REFERENCES orders(orderuid)
	);

CREATE TABLE delivery (
    id integer not null GENERATED ALWAYS AS IDENTITY,
	orderUID         varchar,
	Name       varchar,
	Phone   varchar,	
	Zip            varchar,	
	City varchar,	
	Address        varchar,	
	Region   varchar,	
	Email         varchar,	
	PRIMARY KEY(id),
	CONSTRAINT fk_orderuidunique
   		FOREIGN KEY(orderUID) 
      		REFERENCES orders(orderUID)
	);

