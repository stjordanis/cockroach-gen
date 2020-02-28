// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 1148
	`ALTER`: {
		//line sql.y: 1149
		Category: hGroup,
		//line sql.y: 1150
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER SEQUENCE, ALTER DATABASE, ALTER USER, ALTER ROLE
`,
	},
	//line sql.y: 1166
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 1167
		Category: hDDL,
		//line sql.y: 1168
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
		//line sql.y: 1206
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1220
	`ALTER PARTITION`: {
		ShortDescription: `apply zone configurations to a partition`,
		//line sql.y: 1221
		Category: hDDL,
		//line sql.y: 1222
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
		//line sql.y: 1241
		SeeAlso: `WEBDOCS/configure-zone.html
`,
	},
	//line sql.y: 1246
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1247
		Category: hDDL,
		//line sql.y: 1248
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1250
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1257
	`ALTER SEQUENCE`: {
		ShortDescription: `change the definition of a sequence`,
		//line sql.y: 1258
		Category: hDDL,
		//line sql.y: 1259
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
	//line sql.y: 1282
	`ALTER USER`: {
		ShortDescription: `change user properties`,
		//line sql.y: 1283
		Category: hPriv,
		//line sql.y: 1284
		Text: `
ALTER USER [IF EXISTS] <name> WITH PASSWORD <password>
`,
		//line sql.y: 1286
		SeeAlso: `CREATE USER
`,
	},
	//line sql.y: 1291
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1292
		Category: hDDL,
		//line sql.y: 1293
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1295
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1303
	`ALTER RANGE`: {
		ShortDescription: `change the parameters of a range`,
		//line sql.y: 1304
		Category: hDDL,
		//line sql.y: 1305
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
		//line sql.y: 1317
		SeeAlso: `ALTER TABLE
`,
	},
	//line sql.y: 1322
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1323
		Category: hDDL,
		//line sql.y: 1324
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
		//line sql.y: 1340
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1842
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1843
		Category: hCCL,
		//line sql.y: 1844
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
		//line sql.y: 1861
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 1873
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1874
		Category: hCCL,
		//line sql.y: 1875
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
		//line sql.y: 1891
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 1929
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1930
		Category: hCCL,
		//line sql.y: 1931
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
		//line sql.y: 1959
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 2003
	`EXPORT`: {
		ShortDescription: `export data to file in a distributed manner`,
		//line sql.y: 2004
		Category: hCCL,
		//line sql.y: 2005
		Text: `
EXPORT INTO <format> <datafile> [WITH <option> [= value] [,...]] FROM <query>

Formats:
   CSV

Options:
   delimiter = '...'   [CSV-specific]

`,
		//line sql.y: 2014
		SeeAlso: `SELECT
`,
	},
	//line sql.y: 2108
	`CANCEL`: {
		//line sql.y: 2109
		Category: hGroup,
		//line sql.y: 2110
		Text: `CANCEL JOBS, CANCEL QUERIES, CANCEL SESSIONS
`,
	},
	//line sql.y: 2117
	`CANCEL JOBS`: {
		ShortDescription: `cancel background jobs`,
		//line sql.y: 2118
		Category: hMisc,
		//line sql.y: 2119
		Text: `
CANCEL JOBS <selectclause>
CANCEL JOB <jobid>
`,
		//line sql.y: 2122
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 2140
	`CANCEL QUERIES`: {
		ShortDescription: `cancel running queries`,
		//line sql.y: 2141
		Category: hMisc,
		//line sql.y: 2142
		Text: `
CANCEL QUERIES [IF EXISTS] <selectclause>
CANCEL QUERY [IF EXISTS] <expr>
`,
		//line sql.y: 2145
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 2176
	`CANCEL SESSIONS`: {
		ShortDescription: `cancel open sessions`,
		//line sql.y: 2177
		Category: hMisc,
		//line sql.y: 2178
		Text: `
CANCEL SESSIONS [IF EXISTS] <selectclause>
CANCEL SESSION [IF EXISTS] <sessionid>
`,
		//line sql.y: 2181
		SeeAlso: `SHOW SESSIONS
`,
	},
	//line sql.y: 2251
	`CREATE`: {
		//line sql.y: 2252
		Category: hGroup,
		//line sql.y: 2253
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW, CREATE SEQUENCE, CREATE STATISTICS,
CREATE ROLE
`,
	},
	//line sql.y: 2334
	`CREATE STATISTICS`: {
		ShortDescription: `create a new table statistic`,
		//line sql.y: 2335
		Category: hMisc,
		//line sql.y: 2336
		Text: `
CREATE STATISTICS <statisticname>
  [ON <colname> [, ...]]
  FROM <tablename> [AS OF SYSTEM TIME <expr>]
`,
	},
	//line sql.y: 2479
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 2480
		Category: hDML,
		//line sql.y: 2481
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [ORDER BY <exprs...>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 2485
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 2505
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 2506
		Category: hCfg,
		//line sql.y: 2507
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 2519
	`DROP`: {
		//line sql.y: 2520
		Category: hGroup,
		//line sql.y: 2521
		Text: `
DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP SEQUENCE,
DROP USER, DROP ROLE
`,
	},
	//line sql.y: 2538
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 2539
		Category: hDDL,
		//line sql.y: 2540
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2541
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2553
	`DROP SEQUENCE`: {
		ShortDescription: `remove a sequence`,
		//line sql.y: 2554
		Category: hDDL,
		//line sql.y: 2555
		Text: `DROP SEQUENCE [IF EXISTS] <sequenceName> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2556
		SeeAlso: `DROP
`,
	},
	//line sql.y: 2568
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 2569
		Category: hDDL,
		//line sql.y: 2570
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2571
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 2583
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 2584
		Category: hDDL,
		//line sql.y: 2585
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2586
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2606
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 2607
		Category: hDDL,
		//line sql.y: 2608
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 2609
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 2629
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 2630
		Category: hPriv,
		//line sql.y: 2631
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 2632
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 2644
	`DROP ROLE`: {
		ShortDescription: `remove a role`,
		//line sql.y: 2645
		Category: hPriv,
		//line sql.y: 2646
		Text: `DROP ROLE [IF EXISTS] <role> [, ...]
`,
		//line sql.y: 2647
		SeeAlso: `CREATE ROLE, ALTER ROLE, SHOW ROLES
`,
	},
	//line sql.y: 2671
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 2672
		Category: hMisc,
		//line sql.y: 2673
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
		//line sql.y: 2686
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 2769
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 2770
		Category: hMisc,
		//line sql.y: 2771
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 2772
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2803
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 2804
		Category: hMisc,
		//line sql.y: 2805
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 2806
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2836
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 2837
		Category: hMisc,
		//line sql.y: 2838
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 2839
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 2859
	`GRANT`: {
		ShortDescription: `define access privileges and role memberships`,
		//line sql.y: 2860
		Category: hPriv,
		//line sql.y: 2861
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
		//line sql.y: 2874
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 2890
	`REVOKE`: {
		ShortDescription: `remove access privileges and role memberships`,
		//line sql.y: 2891
		Category: hPriv,
		//line sql.y: 2892
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
		//line sql.y: 2905
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 2959
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 2960
		Category: hCfg,
		//line sql.y: 2961
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 2962
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2974
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 2975
		Category: hCfg,
		//line sql.y: 2976
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 2977
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 2986
	`USE`: {
		ShortDescription: `set the current database`,
		//line sql.y: 2987
		Category: hCfg,
		//line sql.y: 2988
		Text: `USE <dbname>

"USE <dbname>" is an alias for "SET [SESSION] database = <dbname>".
`,
		//line sql.y: 2991
		SeeAlso: `SET SESSION, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 3012
	`SCRUB`: {
		ShortDescription: `run checks against databases or tables`,
		//line sql.y: 3013
		Category: hExperimental,
		//line sql.y: 3014
		Text: `
EXPERIMENTAL SCRUB TABLE <table> ...
EXPERIMENTAL SCRUB DATABASE <database>

The various checks that ca be run with SCRUB includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 3022
		SeeAlso: `SCRUB TABLE, SCRUB DATABASE
`,
	},
	//line sql.y: 3028
	`SCRUB DATABASE`: {
		ShortDescription: `run scrub checks on a database`,
		//line sql.y: 3029
		Category: hExperimental,
		//line sql.y: 3030
		Text: `
EXPERIMENTAL SCRUB DATABASE <database>
                            [AS OF SYSTEM TIME <expr>]

All scrub checks will be run on the database. This includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 3038
		SeeAlso: `SCRUB TABLE, SCRUB
`,
	},
	//line sql.y: 3046
	`SCRUB TABLE`: {
		ShortDescription: `run scrub checks on a table`,
		//line sql.y: 3047
		Category: hExperimental,
		//line sql.y: 3048
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
		//line sql.y: 3059
		SeeAlso: `SCRUB DATABASE, SRUB
`,
	},
	//line sql.y: 3114
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 3115
		Category: hCfg,
		//line sql.y: 3116
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 3117
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3138
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 3139
		Category: hCfg,
		//line sql.y: 3140
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
SET [SESSION] TRACING { TO | = } { on | off | cluster | local | kv | results } [,...]

`,
		//line sql.y: 3146
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 3163
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 3164
		Category: hTxn,
		//line sql.y: 3165
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 3172
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 3364
	`SHOW`: {
		//line sql.y: 3365
		Category: hGroup,
		//line sql.y: 3366
		Text: `
SHOW BACKUP, SHOW CLUSTER SETTING, SHOW COLUMNS, SHOW CONSTRAINTS,
SHOW CREATE, SHOW DATABASES, SHOW HISTOGRAM, SHOW INDEXES, SHOW
PARTITIONS, SHOW JOBS, SHOW QUERIES, SHOW RANGE, SHOW RANGES,
SHOW ROLES, SHOW SCHEMAS, SHOW SEQUENCES, SHOW SESSION, SHOW SESSIONS,
SHOW STATISTICS, SHOW SYNTAX, SHOW TABLES, SHOW TRACE SHOW TRANSACTION, SHOW USERS
`,
	},
	//line sql.y: 3402
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 3403
		Category: hCfg,
		//line sql.y: 3404
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 3405
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 3426
	`SHOW STATISTICS`: {
		ShortDescription: `display table statistics (experimental)`,
		//line sql.y: 3427
		Category: hExperimental,
		//line sql.y: 3428
		Text: `SHOW STATISTICS [USING JSON] FOR TABLE <table_name>

Returns the available statistics for a table.
The statistics can include a histogram ID, which can
be used with SHOW HISTOGRAM.
If USING JSON is specified, the statistics and histograms
are encoded in JSON format.
`,
		//line sql.y: 3435
		SeeAlso: `SHOW HISTOGRAM
`,
	},
	//line sql.y: 3448
	`SHOW HISTOGRAM`: {
		ShortDescription: `display histogram (experimental)`,
		//line sql.y: 3449
		Category: hExperimental,
		//line sql.y: 3450
		Text: `SHOW HISTOGRAM <histogram_id>

Returns the data in the histogram with the
given ID (as returned by SHOW STATISTICS).
`,
		//line sql.y: 3454
		SeeAlso: `SHOW STATISTICS
`,
	},
	//line sql.y: 3467
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 3468
		Category: hCCL,
		//line sql.y: 3469
		Text: `SHOW BACKUP [SCHEMAS|FILES|RANGES] <location>
`,
		//line sql.y: 3470
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 3509
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 3510
		Category: hCfg,
		//line sql.y: 3511
		Text: `
SHOW CLUSTER SETTING <var>
SHOW [ PUBLIC | ALL ] CLUSTER SETTINGS
`,
		//line sql.y: 3514
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3540
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 3541
		Category: hDDL,
		//line sql.y: 3542
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 3543
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 3551
	`SHOW PARTITIONS`: {
		ShortDescription: `list partition information`,
		//line sql.y: 3552
		Category: hDDL,
		//line sql.y: 3553
		Text: `SHOW PARTITIONS FROM { TABLE <table> | INDEX <index> | DATABASE <database> }
`,
		//line sql.y: 3554
		SeeAlso: `WEBDOCS/show-partitions.html
`,
	},
	//line sql.y: 3574
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 3575
		Category: hDDL,
		//line sql.y: 3576
		Text: `SHOW DATABASES
`,
		//line sql.y: 3577
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 3585
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 3586
		Category: hPriv,
		//line sql.y: 3587
		Text: `
Show privilege grants:
  SHOW GRANTS [ON <targets...>] [FOR <users...>]
Show role grants:
  SHOW GRANTS ON ROLE [<roles...>] [FOR <grantees...>]

`,
		//line sql.y: 3593
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 3606
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 3607
		Category: hDDL,
		//line sql.y: 3608
		Text: `SHOW INDEXES FROM { <tablename> | DATABASE <database_name> } [WITH COMMENT]
`,
		//line sql.y: 3609
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 3639
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 3640
		Category: hDDL,
		//line sql.y: 3641
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 3642
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 3655
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 3656
		Category: hMisc,
		//line sql.y: 3657
		Text: `SHOW [ALL] [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 3658
		SeeAlso: `CANCEL QUERIES
`,
	},
	//line sql.y: 3679
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 3680
		Category: hMisc,
		//line sql.y: 3681
		Text: `
SHOW [AUTOMATIC] JOBS
SHOW JOB <jobid>
`,
		//line sql.y: 3684
		SeeAlso: `CANCEL JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3724
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 3725
		Category: hMisc,
		//line sql.y: 3726
		Text: `
SHOW [COMPACT] [KV] TRACE FOR SESSION
`,
		//line sql.y: 3728
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 3751
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 3752
		Category: hMisc,
		//line sql.y: 3753
		Text: `SHOW [ALL] [CLUSTER | LOCAL] SESSIONS
`,
		//line sql.y: 3754
		SeeAlso: `CANCEL SESSIONS
`,
	},
	//line sql.y: 3767
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 3768
		Category: hDDL,
		//line sql.y: 3769
		Text: `SHOW TABLES [FROM <databasename> [ . <schemaname> ] ] [WITH COMMENT]
`,
		//line sql.y: 3770
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 3802
	`SHOW SCHEMAS`: {
		ShortDescription: `list schemas`,
		//line sql.y: 3803
		Category: hDDL,
		//line sql.y: 3804
		Text: `SHOW SCHEMAS [FROM <databasename> ]
`,
	},
	//line sql.y: 3816
	`SHOW SEQUENCES`: {
		ShortDescription: `list sequences`,
		//line sql.y: 3817
		Category: hDDL,
		//line sql.y: 3818
		Text: `SHOW SEQUENCES [FROM <databasename> ]
`,
	},
	//line sql.y: 3830
	`SHOW SYNTAX`: {
		ShortDescription: `analyze SQL syntax`,
		//line sql.y: 3831
		Category: hMisc,
		//line sql.y: 3832
		Text: `SHOW SYNTAX <string>
`,
	},
	//line sql.y: 3841
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 3842
		Category: hCfg,
		//line sql.y: 3843
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 3844
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 3863
	`SHOW CREATE`: {
		ShortDescription: `display the CREATE statement for a table, sequence or view`,
		//line sql.y: 3864
		Category: hDDL,
		//line sql.y: 3865
		Text: `SHOW CREATE [ TABLE | SEQUENCE | VIEW ] <tablename>
`,
		//line sql.y: 3866
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 3884
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 3885
		Category: hPriv,
		//line sql.y: 3886
		Text: `SHOW USERS
`,
		//line sql.y: 3887
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 3895
	`SHOW ROLES`: {
		ShortDescription: `list defined roles`,
		//line sql.y: 3896
		Category: hPriv,
		//line sql.y: 3897
		Text: `SHOW ROLES
`,
		//line sql.y: 3898
		SeeAlso: `CREATE ROLE, ALTER ROLE, DROP ROLE
`,
	},
	//line sql.y: 3954
	`SHOW RANGE`: {
		ShortDescription: `show range information for a row`,
		//line sql.y: 3955
		Category: hMisc,
		//line sql.y: 3956
		Text: `
SHOW RANGE FROM TABLE <tablename> FOR ROW (row, value, ...)
SHOW RANGE FROM INDEX [ <tablename> @ ] <indexname> FOR ROW (row, value, ...)
`,
	},
	//line sql.y: 3977
	`SHOW RANGES`: {
		ShortDescription: `list ranges`,
		//line sql.y: 3978
		Category: hMisc,
		//line sql.y: 3979
		Text: `
SHOW RANGES FROM TABLE <tablename>
SHOW RANGES FROM INDEX [ <tablename> @ ] <indexname>
`,
	},
	//line sql.y: 4216
	`PAUSE JOBS`: {
		ShortDescription: `pause background jobs`,
		//line sql.y: 4217
		Category: hMisc,
		//line sql.y: 4218
		Text: `
PAUSE JOBS <selectclause>
PAUSE JOB <jobid>
`,
		//line sql.y: 4221
		SeeAlso: `SHOW JOBS, CANCEL JOBS, RESUME JOBS
`,
	},
	//line sql.y: 4238
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 4239
		Category: hDDL,
		//line sql.y: 4240
		Text: `
CREATE [[GLOBAL | LOCAL] {TEMPORARY | TEMP}] TABLE [IF NOT EXISTS] <tablename> ( <elements...> ) [<interleave>]
CREATE [[GLOBAL | LOCAL] {TEMPORARY | TEMP}] TABLE [IF NOT EXISTS] <tablename> [( <colnames...> )] AS <source>

Table elements:
   <name> <type> [<qualifiers...>]
   [UNIQUE | INVERTED] INDEX [<name>] ( <colname> [ASC | DESC] [, ...] )
                           [USING HASH WITH BUCKET_COUNT = <shard_buckets>] [STORING ( <colnames...> )] [<interleave>]
   FAMILY [<name>] ( <colnames...> )
   [CONSTRAINT <name>] <constraint>

Table constraints:
   PRIMARY KEY ( <colnames...> ) [USING HASH WITH BUCKET_COUNT = <shard_buckets>]
   FOREIGN KEY ( <colnames...> ) REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
   UNIQUE ( <colnames... ) [STORING ( <colnames...> )] [<interleave>]
   CHECK ( <expr> )

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )] [ON DELETE {NO ACTION | RESTRICT}] [ON UPDATE {NO ACTION | RESTRICT}]
  COLLATE <collationname>
  AS ( <expr> ) STORED

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 4267
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 5012
	`CREATE SEQUENCE`: {
		ShortDescription: `create a new sequence`,
		//line sql.y: 5013
		Category: hDDL,
		//line sql.y: 5014
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
		//line sql.y: 5024
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 5089
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 5090
		Category: hDML,
		//line sql.y: 5091
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 5092
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 5100
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 5101
		Category: hPriv,
		//line sql.y: 5102
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] PASSWORD <passwd> ]
`,
		//line sql.y: 5103
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 5132
	`CREATE ROLE`: {
		ShortDescription: `define a new role`,
		//line sql.y: 5133
		Category: hPriv,
		//line sql.y: 5134
		Text: `CREATE ROLE [IF NOT EXISTS] <name> [WITH] <OPTIONS...>
`,
		//line sql.y: 5135
		SeeAlso: `ALTER ROLE, DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 5163
	`ALTER ROLE`: {
		ShortDescription: `alter a role`,
		//line sql.y: 5164
		Category: hPriv,
		//line sql.y: 5165
		Text: `ALTER ROLE <name> [WITH] <options...>
`,
		//line sql.y: 5166
		SeeAlso: `CREATE ROLE, DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 5184
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 5185
		Category: hDDL,
		//line sql.y: 5186
		Text: `CREATE [TEMPORARY | TEMP] VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 5187
		SeeAlso: `CREATE TABLE, SHOW CREATE, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 5258
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 5259
		Category: hDDL,
		//line sql.y: 5260
		Text: `
CREATE [UNIQUE | INVERTED] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [USING HASH WITH BUCKET_COUNT = <shard_buckets>] [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 5268
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 5499
	`RELEASE`: {
		ShortDescription: `complete a retryable block`,
		//line sql.y: 5500
		Category: hTxn,
		//line sql.y: 5501
		Text: `RELEASE [SAVEPOINT] cockroach_restart
`,
		//line sql.y: 5502
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5510
	`RESUME JOBS`: {
		ShortDescription: `resume background jobs`,
		//line sql.y: 5511
		Category: hMisc,
		//line sql.y: 5512
		Text: `
RESUME JOBS <selectclause>
RESUME JOB <jobid>
`,
		//line sql.y: 5515
		SeeAlso: `SHOW JOBS, CANCEL JOBS, PAUSE JOBS
`,
	},
	//line sql.y: 5532
	`SAVEPOINT`: {
		ShortDescription: `start a retryable block`,
		//line sql.y: 5533
		Category: hTxn,
		//line sql.y: 5534
		Text: `SAVEPOINT cockroach_restart
`,
		//line sql.y: 5535
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5550
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 5551
		Category: hTxn,
		//line sql.y: 5552
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 5560
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 5573
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 5574
		Category: hTxn,
		//line sql.y: 5575
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 5578
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 5602
	`ROLLBACK`: {
		ShortDescription: `abort the current transaction`,
		//line sql.y: 5603
		Category: hTxn,
		//line sql.y: 5604
		Text: `ROLLBACK [TRANSACTION] [TO [SAVEPOINT] cockroach_restart]
`,
		//line sql.y: 5605
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 5723
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 5724
		Category: hDDL,
		//line sql.y: 5725
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 5726
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 5795
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 5796
		Category: hDML,
		//line sql.y: 5797
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5802
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 5821
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 5822
		Category: hDML,
		//line sql.y: 5823
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 5827
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 5938
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 5939
		Category: hDML,
		//line sql.y: 5940
		Text: `
UPDATE <tablename> [[AS] <name>]
       SET ...
       [WHERE <expr>]
       [ORDER BY <exprs...>]
       [LIMIT <expr>]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5947
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 6172
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 6173
		Category: hDML,
		//line sql.y: 6174
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 6185
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 6186
		Category: hDML,
		//line sql.y: 6187
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
		//line sql.y: 6199
		SeeAlso: `WEBDOCS/select-clause.html
`,
	},
	//line sql.y: 6274
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 6275
		Category: hDML,
		//line sql.y: 6276
		Text: `TABLE <tablename>
`,
		//line sql.y: 6277
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 6599
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 6600
		Category: hDML,
		//line sql.y: 6601
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 6602
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 6711
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 6712
		Category: hDML,
		//line sql.y: 6713
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
		//line sql.y: 6735
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
