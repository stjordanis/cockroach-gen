// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 1193
	`ALTER`: {
		//line sql.y: 1194
		Category: hGroup,
		//line sql.y: 1195
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER SEQUENCE, ALTER DATABASE, ALTER USER, ALTER ROLE
`,
	},
	//line sql.y: 1211
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 1212
		Category: hDDL,
		//line sql.y: 1213
		Text: `
ALTER TABLE [IF EXISTS] <tablename> <command> [, ...]

Commands:
  ALTER TABLE ... ADD [COLUMN] [IF NOT EXISTS] <colname> <type> [<qualifiers...>]
  ALTER TABLE ... ADD <constraint>
  ALTER TABLE ... DROP [COLUMN] [IF EXISTS] <colname> [RESTRICT | CASCADE]
  ALTER TABLE ... DROP CONSTRAINT [IF EXISTS] <constraintname> [RESTRICT | CASCADE]
  ALTER TABLE ... ALTER [COLUMN] <colname> {SET DEFAULT <expr> | DROP DEFAULT}
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP NOT NULL
  ALTER TABLE ... ALTER [COLUMN] <colname> DROP STORED
  ALTER TABLE ... ALTER [COLUMN] <colname> [SET DATA] TYPE <type> [COLLATE <collation>]
  ALTER TABLE ... ALTER PRIMARY KEY USING INDEX <name>
  ALTER TABLE ... RENAME TO <newname>
  ALTER TABLE ... RENAME [COLUMN] <colname> TO <newname>
  ALTER TABLE ... VALIDATE CONSTRAINT <constraintname>
  ALTER TABLE ... SPLIT AT <selectclause> [WITH EXPIRATION <expr>]
  ALTER TABLE ... UNSPLIT AT <selectclause>
  ALTER TABLE ... UNSPLIT ALL
  ALTER TABLE ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]
  ALTER TABLE ... INJECT STATISTICS ...  (experimental)
  ALTER TABLE ... PARTITION BY RANGE ( <name...> ) ( <rangespec> )
  ALTER TABLE ... PARTITION BY LIST ( <name...> ) ( <listspec> )
  ALTER TABLE ... PARTITION BY NOTHING
  ALTER TABLE ... CONFIGURE ZONE <zoneconfig>

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )]
  COLLATE <collationname>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1251
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1265
	`ALTER PARTITION`: {
		ShortDescription: `apply zone configurations to a partition`,
		//line sql.y: 1266
		Category: hDDL,
		//line sql.y: 1267
		Text: `
ALTER PARTITION <name> <command>

Commands:
  -- Alter a single partition which exists on any of a table's indexes.
  ALTER PARTITION <partition> OF TABLE <tablename> CONFIGURE ZONE <zoneconfig>

  -- Alter a partition of a specific index.
  ALTER PARTITION <partition> OF INDEX <tablename>@<indexname> CONFIGURE ZONE <zoneconfig>

  -- Alter all partitions with the same name across a table's indexes.
  ALTER PARTITION <partition> OF INDEX <tablename>@* CONFIGURE ZONE <zoneconfig>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1286
		SeeAlso: `WEBDOCS/configure-zone.html
`,
	},
	//line sql.y: 1291
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1292
		Category: hDDL,
		//line sql.y: 1293
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1295
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1302
	`ALTER SEQUENCE`: {
		ShortDescription: `change the definition of a sequence`,
		//line sql.y: 1303
		Category: hDDL,
		//line sql.y: 1304
		Text: `
ALTER SEQUENCE [IF EXISTS] <name>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START <start>]
  [[NO] CYCLE]
ALTER SEQUENCE [IF EXISTS] <name> RENAME TO <newname>
`,
	},
	//line sql.y: 1327
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1328
		Category: hDDL,
		//line sql.y: 1329
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1331
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1339
	`ALTER RANGE`: {
		ShortDescription: `change the parameters of a range`,
		//line sql.y: 1340
		Category: hDDL,
		//line sql.y: 1341
		Text: `
ALTER RANGE <zonename> <command>

Commands:
  ALTER RANGE ... CONFIGURE ZONE <zoneconfig>

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1353
		SeeAlso: `ALTER TABLE
`,
	},
	//line sql.y: 1358
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1359
		Category: hDDL,
		//line sql.y: 1360
		Text: `
ALTER INDEX [IF EXISTS] <idxname> <command>

Commands:
  ALTER INDEX ... RENAME TO <newname>
  ALTER INDEX ... SPLIT AT <selectclause> [WITH EXPIRATION <expr>]
  ALTER INDEX ... UNSPLIT AT <selectclause>
  ALTER INDEX ... UNSPLIT ALL
  ALTER INDEX ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

Zone configurations:
  DISCARD
  USING <var> = <expr> [, ...]
  USING <var> = COPY FROM PARENT [, ...]
  { TO | = } <expr>

`,
		//line sql.y: 1376
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1878
	`ALTER TYPE`: {
		ShortDescription: `change the definition of a type.`,
		//line sql.y: 1879
		Category: hDDL,
		//line sql.y: 1880
		Text: `ALTER TYPE <typename> <command>

Commands:
  ALTER TYPE ... ADD VALUE [IF NOT EXISTS] <value> [ { BEFORE | AFTER } <value> ]
  ALTER TYPE ... RENAME VALUE <oldname> TO <newname>
  ALTER TYPE ... RENAME TO <newname>
  ALTER TYPE ... SET SCHEMA <newschemaname>
  ALTER TYPE ... OWNER TO {<newowner> | CURRENT_USER | SESSION_USER }
  ALTER TYPE ... RENAME ATTRIBUTE <oldname> TO <newname> [ CASCADE | RESTRICT ]
  ALTER TYPE ... <attributeaction> [, ... ]

Attribute action:
  ADD ATTRIBUTE <name> <type> [ COLLATE <collation> ] [ CASCADE | RESTRICT ]
  DROP ATTRIBUTE [IF EXISTS] <name> [ CASCADE | RESTRICT ]
  ALTER ATTRIBUTE <name> [ SET DATA ] TYPE <type> [ COLLATE <collation> ] [ CASCADE | RESTRICT ]

`,
		//line sql.y: 1896
		SeeAlso: `WEBDOCS/alter-type.html
`,
	},
	//line sql.y: 1998
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1999
		Category: hCCL,
		//line sql.y: 2000
		Text: `
BACKUP <targets...> TO <location...>
       [ AS OF SYSTEM TIME <expr> ]
       [ INCREMENTAL FROM <location...> ]
       [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Location:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 2017
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 2029
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 2030
		Category: hCCL,
		//line sql.y: 2031
		Text: `
RESTORE <targets...> FROM <location...>
        [ AS OF SYSTEM TIME <expr> ]
        [ WITH <option> [= <value>] [, ...] ]

Targets:
   TABLE <pattern> [, ...]
   DATABASE <databasename> [, ...]

Locations:
   "[scheme]://[host]/[path to backup]?[parameters]"

Options:
   INTO_DB
   SKIP_MISSING_FOREIGN_KEYS

`,
		//line sql.y: 2047
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 2085
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 2086
		Category: hCCL,
		//line sql.y: 2087
		Text: `
-- Import both schema and table data:
IMPORT [ TABLE <tablename> FROM ]
       <format> <datafile>
       [ WITH <option> [= <value>] [, ...] ]

-- Import using specific schema, use only table data from external file:
IMPORT TABLE <tablename>
       { ( <elements> ) | CREATE USING <schemafile> }
       <format>
       DATA ( <datafile> [, ...] )
       [ WITH <option> [= <value>] [, ...] ]

Formats:
   CSV
   DELIMITED
   MYSQLDUMP
   PGCOPY
   PGDUMP

Options:
   distributed = '...'
   sstsize = '...'
   temp = '...'
   delimiter = '...'      [CSV, PGCOPY-specific]
   nullif = '...'         [CSV, PGCOPY-specific]
   comment = '...'        [CSV-specific]

`,
		//line sql.y: 2115
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 2159
	`EXPORT`: {
		ShortDescription: `export data to file in a distributed manner`,
		//line sql.y: 2160
		Category: hCCL,
		//line sql.y: 2161
		Text: `
EXPORT INTO <format> <datafile> [WITH <option> [= value] [,...]] FROM <query>

Formats:
   CSV

Options:
   delimiter = '...'   [CSV-specific]

`,
		//line sql.y: 2170
		SeeAlso: `SELECT
`,
	},
	//line sql.y: 2264
	`CANCEL`: {
		//line sql.y: 2265
		Category: hGroup,
		//line sql.y: 2266
		Text: `CANCEL JOBS, CANCEL QUERIES, CANCEL SESSIONS
`,
	},
	//line sql.y: 2273
	`CANCEL JOBS`: {
		ShortDescription: `cancel background jobs`,
		//line sql.y: 2274
		Category: hMisc,
		//line sql.y: 2275
		Text: `
CANCEL JOBS <selectclause>
CANCEL JOB <jobid>
`,
		//line sql.y: 2278
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 2296
	`CANCEL QUERIES`: {
		ShortDescription: `cancel running queries`,
		//line sql.y: 2297
		Category: hMisc,
		//line sql.y: 2298
		Text: `
CANCEL QUERIES [IF EXISTS] <selectclause>
CANCEL QUERY [IF EXISTS] <expr>
`,
		//line sql.y: 2301
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 2332
	`CANCEL SESSIONS`: {
		ShortDescription: `cancel open sessions`,
		//line sql.y: 2333
		Category: hMisc,
		//line sql.y: 2334
		Text: `
CANCEL SESSIONS [IF EXISTS] <selectclause>
CANCEL SESSION [IF EXISTS] <sessionid>
`,
		//line sql.y: 2337
		SeeAlso: `SHOW SESSIONS
`,
	},
	//line sql.y: 2407
	`CREATE`: {
		//line sql.y: 2408
		Category: hGroup,
		//line sql.y: 2409
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW, CREATE SEQUENCE, CREATE STATISTICS,
CREATE ROLE, CREATE TYPE
`,
	},
	//line sql.y: 2488
	`CREATE STATISTICS`: {
		ShortDescription: `create a new table statistic`,
		//line sql.y: 2489
		Category: hMisc,
		//line sql.y: 2490
		Text: `
CREATE STATISTICS <statisticname>
  [ON <colname> [, ...]]
  FROM <tablename> [AS OF SYSTEM TIME <expr>]
`,
	},
	//line sql.y: 2633
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 2634
		Category: hDML,
		//line sql.y: 2635
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [ORDER BY <exprs...>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 2639
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 2659
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 2660
		Category: hCfg,
		//line sql.y: 2661
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 2673
	`DROP`: {
		//line sql.y: 2674
		Category: hGroup,
		//line sql.y: 2675
		Text: `
DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP SEQUENCE,
DROP USER, DROP ROLE, DROP TYPE
`,
	},
	//line sql.y: 2692
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 2693
		Category: hDDL,
		//line sql.y: 2694
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2695
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2707
	`DROP SEQUENCE`: {
		ShortDescription: `remove a sequence`,
		//line sql.y: 2708
		Category: hDDL,
		//line sql.y: 2709
		Text: `DROP SEQUENCE [IF EXISTS] <sequenceName> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2710
		SeeAlso: `DROP
`,
	},
	//line sql.y: 2722
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 2723
		Category: hDDL,
		//line sql.y: 2724
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2725
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 2737
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 2738
		Category: hDDL,
		//line sql.y: 2739
		Text: `DROP INDEX [CONCURRENTLY] [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2740
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2762
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 2763
		Category: hDDL,
		//line sql.y: 2764
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 2765
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 2785
	`DROP TYPE`: {
		ShortDescription: `remove a type`,
		//line sql.y: 2786
		Category: hDDL,
		//line sql.y: 2787
		Text: `DROP TYPE [IF EXISTS] <type_name> [, ...] [CASCASE | RESTRICT]
`,
		//line sql.y: 2788
		SeeAlso: `WEBDOCS/drop-type.html
`,
	},
	//line sql.y: 2820
	`DROP ROLE`: {
		ShortDescription: `remove a user`,
		//line sql.y: 2821
		Category: hPriv,
		//line sql.y: 2822
		Text: `DROP ROLE [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 2823
		SeeAlso: `CREATE ROLE, SHOW ROLE
`,
	},
	//line sql.y: 2847
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 2848
		Category: hMisc,
		//line sql.y: 2849
		Text: `
EXPLAIN <statement>
EXPLAIN ([PLAN ,] <planoptions...> ) <statement>
EXPLAIN [ANALYZE] (DISTSQL) <statement>
EXPLAIN ANALYZE [(DISTSQL)] <statement>

Explainable statements:
    SELECT, CREATE, DROP, ALTER, INSERT, UPSERT, UPDATE, DELETE,
    SHOW, EXPLAIN

Plan options:
    TYPES, VERBOSE, OPT

`,
		//line sql.y: 2862
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 2969
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 2970
		Category: hMisc,
		//line sql.y: 2971
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 2972
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 3003
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 3004
		Category: hMisc,
		//line sql.y: 3005
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 3006
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 3036
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 3037
		Category: hMisc,
		//line sql.y: 3038
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 3039
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 3059
	`GRANT`: {
		ShortDescription: `define access privileges and role memberships`,
		//line sql.y: 3060
		Category: hPriv,
		//line sql.y: 3061
		Text: `
Grant privileges:
  GRANT {ALL | <privileges...> } ON <targets...> TO <grantees...>
Grant role membership (CCL only):
  GRANT <roles...> TO <grantees...> [WITH ADMIN OPTION]

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, ...]
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 3074
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 3090
	`REVOKE`: {
		ShortDescription: `remove access privileges and role memberships`,
		//line sql.y: 3091
		Category: hPriv,
		//line sql.y: 3092
		Text: `
Revoke privileges:
  REVOKE {ALL | <privileges...> } ON <targets...> FROM <grantees...>
Revoke role membership (CCL only):
  REVOKE [ADMIN OPTION FOR] <roles...> FROM <grantees...>

Privileges:
  CREATE, DROP, GRANT, SELECT, INSERT, DELETE, UPDATE

Targets:
  DATABASE <databasename> [, <databasename>]...
  [TABLE] [<databasename> .] { <tablename> | * } [, ...]

`,
		//line sql.y: 3105
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 3159
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 3160
		Category: hCfg,
		//line sql.y: 3161
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 3162
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 3174
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 3175
		Category: hCfg,
		//line sql.y: 3176
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 3177
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 3186
	`USE`: {
		ShortDescription: `set the current database`,
		//line sql.y: 3187
		Category: hCfg,
		//line sql.y: 3188
		Text: `USE <dbname>

"USE <dbname>" is an alias for "SET [SESSION] database = <dbname>".
`,
		//line sql.y: 3191
		SeeAlso: `SET SESSION, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 3212
	`SCRUB`: {
		ShortDescription: `run checks against databases or tables`,
		//line sql.y: 3213
		Category: hExperimental,
		//line sql.y: 3214
		Text: `
EXPERIMENTAL SCRUB TABLE <table> ...
EXPERIMENTAL SCRUB DATABASE <database>

The various checks that ca be run with SCRUB includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 3222
		SeeAlso: `SCRUB TABLE, SCRUB DATABASE
`,
	},
	//line sql.y: 3228
	`SCRUB DATABASE`: {
		ShortDescription: `run scrub checks on a database`,
		//line sql.y: 3229
		Category: hExperimental,
		//line sql.y: 3230
		Text: `
EXPERIMENTAL SCRUB DATABASE <database>
                            [AS OF SYSTEM TIME <expr>]

All scrub checks will be run on the database. This includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 3238
		SeeAlso: `SCRUB TABLE, SCRUB
`,
	},
	//line sql.y: 3246
	`SCRUB TABLE`: {
		ShortDescription: `run scrub checks on a table`,
		//line sql.y: 3247
		Category: hExperimental,
		//line sql.y: 3248
		Text: `
SCRUB TABLE <tablename>
            [AS OF SYSTEM TIME <expr>]
            [WITH OPTIONS <option> [, ...]]

Options:
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS INDEX ALL
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS INDEX (<index>...)
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS CONSTRAINT ALL
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS CONSTRAINT (<constraint>...)
  EXPERIMENTAL SCRUB TABLE ... WITH OPTIONS PHYSICAL
`,
		//line sql.y: 3259
		SeeAlso: `SCRUB DATABASE, SRUB
`,
	},
	//line sql.y: 3314
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 3315
		Category: hCfg,
		//line sql.y: 3316
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 3317
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3338
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 3339
		Category: hCfg,
		//line sql.y: 3340
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
SET [SESSION] TRACING { TO | = } { on | off | cluster | local | kv | results } [,...]

`,
		//line sql.y: 3346
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 3363
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 3364
		Category: hTxn,
		//line sql.y: 3365
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 3372
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 3564
	`SHOW`: {
		//line sql.y: 3565
		Category: hGroup,
		//line sql.y: 3566
		Text: `
SHOW BACKUP, SHOW CLUSTER SETTING, SHOW COLUMNS, SHOW CONSTRAINTS,
SHOW CREATE, SHOW DATABASES, SHOW HISTOGRAM, SHOW INDEXES, SHOW
PARTITIONS, SHOW JOBS, SHOW QUERIES, SHOW RANGE, SHOW RANGES,
SHOW ROLES, SHOW SCHEMAS, SHOW SEQUENCES, SHOW SESSION, SHOW SESSIONS,
SHOW STATISTICS, SHOW SYNTAX, SHOW TABLES, SHOW TRACE SHOW TRANSACTION, SHOW USERS
`,
	},
	//line sql.y: 3634
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 3635
		Category: hCfg,
		//line sql.y: 3636
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 3637
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 3658
	`SHOW STATISTICS`: {
		ShortDescription: `display table statistics (experimental)`,
		//line sql.y: 3659
		Category: hExperimental,
		//line sql.y: 3660
		Text: `SHOW STATISTICS [USING JSON] FOR TABLE <table_name>

Returns the available statistics for a table.
The statistics can include a histogram ID, which can
be used with SHOW HISTOGRAM.
If USING JSON is specified, the statistics and histograms
are encoded in JSON format.
`,
		//line sql.y: 3667
		SeeAlso: `SHOW HISTOGRAM
`,
	},
	//line sql.y: 3680
	`SHOW HISTOGRAM`: {
		ShortDescription: `display histogram (experimental)`,
		//line sql.y: 3681
		Category: hExperimental,
		//line sql.y: 3682
		Text: `SHOW HISTOGRAM <histogram_id>

Returns the data in the histogram with the
given ID (as returned by SHOW STATISTICS).
`,
		//line sql.y: 3686
		SeeAlso: `SHOW STATISTICS
`,
	},
	//line sql.y: 3699
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 3700
		Category: hCCL,
		//line sql.y: 3701
		Text: `SHOW BACKUP [SCHEMAS|FILES|RANGES] <location>
`,
		//line sql.y: 3702
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 3741
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 3742
		Category: hCfg,
		//line sql.y: 3743
		Text: `
SHOW CLUSTER SETTING <var>
SHOW [ PUBLIC | ALL ] CLUSTER SETTINGS
`,
		//line sql.y: 3746
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3772
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 3773
		Category: hDDL,
		//line sql.y: 3774
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 3775
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 3783
	`SHOW PARTITIONS`: {
		ShortDescription: `list partition information`,
		//line sql.y: 3784
		Category: hDDL,
		//line sql.y: 3785
		Text: `SHOW PARTITIONS FROM { TABLE <table> | INDEX <index> | DATABASE <database> }
`,
		//line sql.y: 3786
		SeeAlso: `WEBDOCS/show-partitions.html
`,
	},
	//line sql.y: 3806
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 3807
		Category: hDDL,
		//line sql.y: 3808
		Text: `SHOW DATABASES
`,
		//line sql.y: 3809
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 3817
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 3818
		Category: hPriv,
		//line sql.y: 3819
		Text: `
Show privilege grants:
  SHOW GRANTS [ON <targets...>] [FOR <users...>]
Show role grants:
  SHOW GRANTS ON ROLE [<roles...>] [FOR <grantees...>]

`,
		//line sql.y: 3825
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 3838
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 3839
		Category: hDDL,
		//line sql.y: 3840
		Text: `SHOW INDEXES FROM { <tablename> | DATABASE <database_name> } [WITH COMMENT]
`,
		//line sql.y: 3841
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 3871
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 3872
		Category: hDDL,
		//line sql.y: 3873
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 3874
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 3887
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 3888
		Category: hMisc,
		//line sql.y: 3889
		Text: `SHOW [ALL] [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 3890
		SeeAlso: `CANCEL QUERIES
`,
	},
	//line sql.y: 3911
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 3912
		Category: hMisc,
		//line sql.y: 3913
		Text: `
SHOW [AUTOMATIC] JOBS
SHOW JOB <jobid>
`,
		//line sql.y: 3916
		SeeAlso: `CANCEL JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3956
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 3957
		Category: hMisc,
		//line sql.y: 3958
		Text: `
SHOW [COMPACT] [KV] TRACE FOR SESSION
`,
		//line sql.y: 3960
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 3983
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 3984
		Category: hMisc,
		//line sql.y: 3985
		Text: `SHOW [ALL] [CLUSTER | LOCAL] SESSIONS
`,
		//line sql.y: 3986
		SeeAlso: `CANCEL SESSIONS
`,
	},
	//line sql.y: 3999
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 4000
		Category: hDDL,
		//line sql.y: 4001
		Text: `SHOW TABLES [FROM <databasename> [ . <schemaname> ] ] [WITH COMMENT]
`,
		//line sql.y: 4002
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 4034
	`SHOW SCHEMAS`: {
		ShortDescription: `list schemas`,
		//line sql.y: 4035
		Category: hDDL,
		//line sql.y: 4036
		Text: `SHOW SCHEMAS [FROM <databasename> ]
`,
	},
	//line sql.y: 4048
	`SHOW SEQUENCES`: {
		ShortDescription: `list sequences`,
		//line sql.y: 4049
		Category: hDDL,
		//line sql.y: 4050
		Text: `SHOW SEQUENCES [FROM <databasename> ]
`,
	},
	//line sql.y: 4062
	`SHOW SYNTAX`: {
		ShortDescription: `analyze SQL syntax`,
		//line sql.y: 4063
		Category: hMisc,
		//line sql.y: 4064
		Text: `SHOW SYNTAX <string>
`,
	},
	//line sql.y: 4073
	`SHOW SAVEPOINT`: {
		ShortDescription: `display current savepoint properties`,
		//line sql.y: 4074
		Category: hCfg,
		//line sql.y: 4075
		Text: `SHOW SAVEPOINT STATUS
`,
	},
	//line sql.y: 4083
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 4084
		Category: hCfg,
		//line sql.y: 4085
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 4086
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 4105
	`SHOW CREATE`: {
		ShortDescription: `display the CREATE statement for a table, sequence or view`,
		//line sql.y: 4106
		Category: hDDL,
		//line sql.y: 4107
		Text: `SHOW CREATE [ TABLE | SEQUENCE | VIEW ] <tablename>
`,
		//line sql.y: 4108
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 4126
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 4127
		Category: hPriv,
		//line sql.y: 4128
		Text: `SHOW USERS
`,
		//line sql.y: 4129
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 4137
	`SHOW ROLES`: {
		ShortDescription: `list defined roles`,
		//line sql.y: 4138
		Category: hPriv,
		//line sql.y: 4139
		Text: `SHOW ROLES
`,
		//line sql.y: 4140
		SeeAlso: `CREATE ROLE, ALTER ROLE, DROP ROLE
`,
	},
	//line sql.y: 4196
	`SHOW RANGE`: {
		ShortDescription: `show range information for a row`,
		//line sql.y: 4197
		Category: hMisc,
		//line sql.y: 4198
		Text: `
SHOW RANGE FROM TABLE <tablename> FOR ROW (row, value, ...)
SHOW RANGE FROM INDEX [ <tablename> @ ] <indexname> FOR ROW (row, value, ...)
`,
	},
	//line sql.y: 4219
	`SHOW RANGES`: {
		ShortDescription: `list ranges`,
		//line sql.y: 4220
		Category: hMisc,
		//line sql.y: 4221
		Text: `
SHOW RANGES FROM TABLE <tablename>
SHOW RANGES FROM INDEX [ <tablename> @ ] <indexname>
`,
	},
	//line sql.y: 4458
	`PAUSE JOBS`: {
		ShortDescription: `pause background jobs`,
		//line sql.y: 4459
		Category: hMisc,
		//line sql.y: 4460
		Text: `
PAUSE JOBS <selectclause>
PAUSE JOB <jobid>
`,
		//line sql.y: 4463
		SeeAlso: `SHOW JOBS, CANCEL JOBS, RESUME JOBS
`,
	},
	//line sql.y: 4480
	`CREATE SCHEMA`: {
		ShortDescription: `create a new schema (not yet supported)`,
		//line sql.y: 4481
		Category: hDDL,
		//line sql.y: 4482
		Text: `
CREATE SCHEMA [IF NOT EXISTS] <schemaname>
`,
	},
	//line sql.y: 4500
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 4501
		Category: hDDL,
		//line sql.y: 4502
		Text: `
CREATE [[GLOBAL | LOCAL] {TEMPORARY | TEMP}] TABLE [IF NOT EXISTS] <tablename> ( <elements...> ) [<interleave>] [<on_commit>]
CREATE [[GLOBAL | LOCAL] {TEMPORARY | TEMP}] TABLE [IF NOT EXISTS] <tablename> [( <colnames...> )] AS <source> [<interleave>] [<on commit>]

Table elements:
   <name> <type> [<qualifiers...>]
   [UNIQUE | INVERTED] INDEX [<name>] ( <colname> [ASC | DESC] [, ...] )
                           [USING HASH WITH BUCKET_COUNT = <shard_buckets>] [{STORING | INCLUDE | COVERING} ( <colnames...> )] [<interleave>]
   FAMILY [<name>] ( <colnames...> )
   [CONSTRAINT <name>] <constraint>

Table constraints:
   PRIMARY KEY ( <colnames...> ) [USING HASH WITH BUCKET_COUNT = <shard_buckets>]
   FOREIGN KEY ( <colnames...> ) REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
   UNIQUE ( <colnames... ) [{STORING | INCLUDE | COVERING} ( <colnames...> )] [<interleave>]
   CHECK ( <expr> )

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
  COLLATE <collationname>
  AS ( <expr> ) STORED

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

On commit clause:
   ON COMMIT {PRESERVE ROWS | DROP | DELETE ROWS}

`,
		//line sql.y: 4532
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 5351
	`CREATE SEQUENCE`: {
		ShortDescription: `create a new sequence`,
		//line sql.y: 5352
		Category: hDDL,
		//line sql.y: 5353
		Text: `
CREATE [TEMPORARY | TEMP] SEQUENCE <seqname>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START [WITH] <start>]
  [CACHE <cache>]
  [NO CYCLE]
  [VIRTUAL]

`,
		//line sql.y: 5363
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 5428
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 5429
		Category: hDML,
		//line sql.y: 5430
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 5431
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 5449
	`CREATE ROLE`: {
		ShortDescription: `define a new role`,
		//line sql.y: 5450
		Category: hPriv,
		//line sql.y: 5451
		Text: `CREATE ROLE [IF NOT EXISTS] <name> [ [WITH] <OPTIONS...> ]
`,
		//line sql.y: 5452
		SeeAlso: `ALTER ROLE, DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 5464
	`ALTER ROLE`: {
		ShortDescription: `alter a role`,
		//line sql.y: 5465
		Category: hPriv,
		//line sql.y: 5466
		Text: `ALTER ROLE <name> [WITH] <options...>
`,
		//line sql.y: 5467
		SeeAlso: `CREATE ROLE, DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 5496
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 5497
		Category: hDDL,
		//line sql.y: 5498
		Text: `CREATE [TEMPORARY | TEMP] VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 5499
		SeeAlso: `CREATE TABLE, SHOW CREATE, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 5596
	`CREATE TYPE`: {
		ShortDescription: `- create a type`,
		//line sql.y: 5597
		Category: hDDL,
		//line sql.y: 5598
		Text: `CREATE TYPE <type_name> AS ENUM (...)
`,
		//line sql.y: 5599
		SeeAlso: `WEBDOCS/create-type.html
`,
	},
	//line sql.y: 5642
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 5643
		Category: hDDL,
		//line sql.y: 5644
		Text: `
CREATE [UNIQUE | INVERTED] INDEX [CONCURRENTLY] [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [USING HASH WITH BUCKET_COUNT = <shard_buckets>] [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 5652
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 5886
	`RELEASE`: {
		ShortDescription: `complete a sub-transaction`,
		//line sql.y: 5887
		Category: hTxn,
		//line sql.y: 5888
		Text: `RELEASE [SAVEPOINT] <savepoint name>
`,
		//line sql.y: 5889
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5897
	`RESUME JOBS`: {
		ShortDescription: `resume background jobs`,
		//line sql.y: 5898
		Category: hMisc,
		//line sql.y: 5899
		Text: `
RESUME JOBS <selectclause>
RESUME JOB <jobid>
`,
		//line sql.y: 5902
		SeeAlso: `SHOW JOBS, CANCEL JOBS, PAUSE JOBS
`,
	},
	//line sql.y: 5919
	`SAVEPOINT`: {
		ShortDescription: `start a sub-transaction`,
		//line sql.y: 5920
		Category: hTxn,
		//line sql.y: 5921
		Text: `SAVEPOINT <savepoint name>
`,
		//line sql.y: 5922
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5937
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 5938
		Category: hTxn,
		//line sql.y: 5939
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 5947
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 5960
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 5961
		Category: hTxn,
		//line sql.y: 5962
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 5965
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 5989
	`ROLLBACK`: {
		ShortDescription: `abort the current (sub-)transaction`,
		//line sql.y: 5990
		Category: hTxn,
		//line sql.y: 5991
		Text: `
ROLLBACK [TRANSACTION]
ROLLBACK [TRANSACTION] TO [SAVEPOINT] <savepoint name>
`,
		//line sql.y: 5994
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 6094
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 6095
		Category: hDDL,
		//line sql.y: 6096
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 6097
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 6166
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 6167
		Category: hDML,
		//line sql.y: 6168
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 6173
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 6192
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 6193
		Category: hDML,
		//line sql.y: 6194
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 6198
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 6309
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 6310
		Category: hDML,
		//line sql.y: 6311
		Text: `
UPDATE <tablename> [[AS] <name>]
       SET ...
       [WHERE <expr>]
       [ORDER BY <exprs...>]
       [LIMIT <expr>]
       [RETURNING <exprs...>]
`,
		//line sql.y: 6318
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 6543
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 6544
		Category: hDML,
		//line sql.y: 6545
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 6556
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 6557
		Category: hDML,
		//line sql.y: 6558
		Text: `
SELECT [DISTINCT [ ON ( <expr> [ , ... ] ) ] ]
       { <expr> [[AS] <name>] | [ [<dbname>.] <tablename>. ] * } [, ...]
       [ FROM <source> ]
       [ WHERE <expr> ]
       [ GROUP BY <expr> [ , ... ] ]
       [ HAVING <expr> ]
       [ WINDOW <name> AS ( <definition> ) ]
       [ { UNION | INTERSECT | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
       [ ORDER BY <expr> [ ASC | DESC ] [, ...] ]
       [ LIMIT { <expr> | ALL } ]
       [ OFFSET <expr> [ ROW | ROWS ] ]
`,
		//line sql.y: 6570
		SeeAlso: `WEBDOCS/select-clause.html
`,
	},
	//line sql.y: 6645
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 6646
		Category: hDML,
		//line sql.y: 6647
		Text: `TABLE <tablename>
`,
		//line sql.y: 6648
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 7005
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 7006
		Category: hDML,
		//line sql.y: 7007
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 7008
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 7117
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 7118
		Category: hDML,
		//line sql.y: 7119
		Text: `
Data sources:
  <tablename> [ @ { <idxname> | <indexflags> } ]
  <tablefunc> ( <exprs...> )
  ( { <selectclause> | <source> } )
  <source> [AS] <alias> [( <colnames...> )]
  <source> [ <jointype> ] JOIN <source> ON <expr>
  <source> [ <jointype> ] JOIN <source> USING ( <colnames...> )
  <source> NATURAL [ <jointype> ] JOIN <source>
  <source> CROSS JOIN <source>
  <source> WITH ORDINALITY
  '[' EXPLAIN ... ']'
  '[' SHOW ... ']'

Index flags:
  '{' FORCE_INDEX = <idxname> [, ...] '}'
  '{' NO_INDEX_JOIN [, ...] '}'
  '{' IGNORE_FOREIGN_KEYS [, ...] '}'

Join types:
  { INNER | { LEFT | RIGHT | FULL } [OUTER] } [ { HASH | MERGE | LOOKUP } ]

`,
		//line sql.y: 7141
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
