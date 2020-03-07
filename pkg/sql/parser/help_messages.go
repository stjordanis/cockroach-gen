// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 1140
	`ALTER`: {
		//line sql.y: 1141
		Category: hGroup,
		//line sql.y: 1142
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER SEQUENCE, ALTER DATABASE, ALTER USER, ALTER ROLE
`,
	},
	//line sql.y: 1158
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 1159
		Category: hDDL,
		//line sql.y: 1160
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
		//line sql.y: 1198
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1212
	`ALTER PARTITION`: {
		ShortDescription: `apply zone configurations to a partition`,
		//line sql.y: 1213
		Category: hDDL,
		//line sql.y: 1214
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
		//line sql.y: 1233
		SeeAlso: `WEBDOCS/configure-zone.html
`,
	},
	//line sql.y: 1238
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1239
		Category: hDDL,
		//line sql.y: 1240
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1242
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1249
	`ALTER SEQUENCE`: {
		ShortDescription: `change the definition of a sequence`,
		//line sql.y: 1250
		Category: hDDL,
		//line sql.y: 1251
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
	//line sql.y: 1274
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1275
		Category: hDDL,
		//line sql.y: 1276
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1278
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1286
	`ALTER RANGE`: {
		ShortDescription: `change the parameters of a range`,
		//line sql.y: 1287
		Category: hDDL,
		//line sql.y: 1288
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
		//line sql.y: 1300
		SeeAlso: `ALTER TABLE
`,
	},
	//line sql.y: 1305
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1306
		Category: hDDL,
		//line sql.y: 1307
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
		//line sql.y: 1323
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1825
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1826
		Category: hCCL,
		//line sql.y: 1827
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
		//line sql.y: 1844
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 1856
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1857
		Category: hCCL,
		//line sql.y: 1858
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
		//line sql.y: 1874
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 1912
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1913
		Category: hCCL,
		//line sql.y: 1914
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
		//line sql.y: 1942
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 1986
	`EXPORT`: {
		ShortDescription: `export data to file in a distributed manner`,
		//line sql.y: 1987
		Category: hCCL,
		//line sql.y: 1988
		Text: `
EXPORT INTO <format> <datafile> [WITH <option> [= value] [,...]] FROM <query>

Formats:
   CSV

Options:
   delimiter = '...'   [CSV-specific]

`,
		//line sql.y: 1997
		SeeAlso: `SELECT
`,
	},
	//line sql.y: 2091
	`CANCEL`: {
		//line sql.y: 2092
		Category: hGroup,
		//line sql.y: 2093
		Text: `CANCEL JOBS, CANCEL QUERIES, CANCEL SESSIONS
`,
	},
	//line sql.y: 2100
	`CANCEL JOBS`: {
		ShortDescription: `cancel background jobs`,
		//line sql.y: 2101
		Category: hMisc,
		//line sql.y: 2102
		Text: `
CANCEL JOBS <selectclause>
CANCEL JOB <jobid>
`,
		//line sql.y: 2105
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 2123
	`CANCEL QUERIES`: {
		ShortDescription: `cancel running queries`,
		//line sql.y: 2124
		Category: hMisc,
		//line sql.y: 2125
		Text: `
CANCEL QUERIES [IF EXISTS] <selectclause>
CANCEL QUERY [IF EXISTS] <expr>
`,
		//line sql.y: 2128
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 2159
	`CANCEL SESSIONS`: {
		ShortDescription: `cancel open sessions`,
		//line sql.y: 2160
		Category: hMisc,
		//line sql.y: 2161
		Text: `
CANCEL SESSIONS [IF EXISTS] <selectclause>
CANCEL SESSION [IF EXISTS] <sessionid>
`,
		//line sql.y: 2164
		SeeAlso: `SHOW SESSIONS
`,
	},
	//line sql.y: 2234
	`CREATE`: {
		//line sql.y: 2235
		Category: hGroup,
		//line sql.y: 2236
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW, CREATE SEQUENCE, CREATE STATISTICS,
CREATE ROLE
`,
	},
	//line sql.y: 2317
	`CREATE STATISTICS`: {
		ShortDescription: `create a new table statistic`,
		//line sql.y: 2318
		Category: hMisc,
		//line sql.y: 2319
		Text: `
CREATE STATISTICS <statisticname>
  [ON <colname> [, ...]]
  FROM <tablename> [AS OF SYSTEM TIME <expr>]
`,
	},
	//line sql.y: 2462
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 2463
		Category: hDML,
		//line sql.y: 2464
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [ORDER BY <exprs...>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 2468
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 2488
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 2489
		Category: hCfg,
		//line sql.y: 2490
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 2502
	`DROP`: {
		//line sql.y: 2503
		Category: hGroup,
		//line sql.y: 2504
		Text: `
DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP SEQUENCE,
DROP USER, DROP ROLE
`,
	},
	//line sql.y: 2521
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 2522
		Category: hDDL,
		//line sql.y: 2523
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2524
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2536
	`DROP SEQUENCE`: {
		ShortDescription: `remove a sequence`,
		//line sql.y: 2537
		Category: hDDL,
		//line sql.y: 2538
		Text: `DROP SEQUENCE [IF EXISTS] <sequenceName> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2539
		SeeAlso: `DROP
`,
	},
	//line sql.y: 2551
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 2552
		Category: hDDL,
		//line sql.y: 2553
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2554
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 2566
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 2567
		Category: hDDL,
		//line sql.y: 2568
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2569
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2589
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 2590
		Category: hDDL,
		//line sql.y: 2591
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 2592
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 2612
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 2613
		Category: hPriv,
		//line sql.y: 2614
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 2615
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 2627
	`DROP ROLE`: {
		ShortDescription: `remove a user`,
		//line sql.y: 2628
		Category: hPriv,
		//line sql.y: 2629
		Text: `DROP ROLE [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 2630
		SeeAlso: `CREATE ROLE, SHOW ROLE
`,
	},
	//line sql.y: 2654
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 2655
		Category: hMisc,
		//line sql.y: 2656
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
		//line sql.y: 2669
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 2756
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 2757
		Category: hMisc,
		//line sql.y: 2758
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 2759
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2790
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 2791
		Category: hMisc,
		//line sql.y: 2792
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 2793
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2823
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 2824
		Category: hMisc,
		//line sql.y: 2825
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 2826
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 2846
	`GRANT`: {
		ShortDescription: `define access privileges and role memberships`,
		//line sql.y: 2847
		Category: hPriv,
		//line sql.y: 2848
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
		//line sql.y: 2861
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 2877
	`REVOKE`: {
		ShortDescription: `remove access privileges and role memberships`,
		//line sql.y: 2878
		Category: hPriv,
		//line sql.y: 2879
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
		//line sql.y: 2892
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 2946
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 2947
		Category: hCfg,
		//line sql.y: 2948
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 2949
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2961
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 2962
		Category: hCfg,
		//line sql.y: 2963
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 2964
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 2973
	`USE`: {
		ShortDescription: `set the current database`,
		//line sql.y: 2974
		Category: hCfg,
		//line sql.y: 2975
		Text: `USE <dbname>

"USE <dbname>" is an alias for "SET [SESSION] database = <dbname>".
`,
		//line sql.y: 2978
		SeeAlso: `SET SESSION, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2999
	`SCRUB`: {
		ShortDescription: `run checks against databases or tables`,
		//line sql.y: 3000
		Category: hExperimental,
		//line sql.y: 3001
		Text: `
EXPERIMENTAL SCRUB TABLE <table> ...
EXPERIMENTAL SCRUB DATABASE <database>

The various checks that ca be run with SCRUB includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 3009
		SeeAlso: `SCRUB TABLE, SCRUB DATABASE
`,
	},
	//line sql.y: 3015
	`SCRUB DATABASE`: {
		ShortDescription: `run scrub checks on a database`,
		//line sql.y: 3016
		Category: hExperimental,
		//line sql.y: 3017
		Text: `
EXPERIMENTAL SCRUB DATABASE <database>
                            [AS OF SYSTEM TIME <expr>]

All scrub checks will be run on the database. This includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 3025
		SeeAlso: `SCRUB TABLE, SCRUB
`,
	},
	//line sql.y: 3033
	`SCRUB TABLE`: {
		ShortDescription: `run scrub checks on a table`,
		//line sql.y: 3034
		Category: hExperimental,
		//line sql.y: 3035
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
		//line sql.y: 3046
		SeeAlso: `SCRUB DATABASE, SRUB
`,
	},
	//line sql.y: 3101
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 3102
		Category: hCfg,
		//line sql.y: 3103
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 3104
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3125
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 3126
		Category: hCfg,
		//line sql.y: 3127
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
SET [SESSION] TRACING { TO | = } { on | off | cluster | local | kv | results } [,...]

`,
		//line sql.y: 3133
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 3150
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 3151
		Category: hTxn,
		//line sql.y: 3152
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 3159
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 3351
	`SHOW`: {
		//line sql.y: 3352
		Category: hGroup,
		//line sql.y: 3353
		Text: `
SHOW BACKUP, SHOW CLUSTER SETTING, SHOW COLUMNS, SHOW CONSTRAINTS,
SHOW CREATE, SHOW DATABASES, SHOW HISTOGRAM, SHOW INDEXES, SHOW
PARTITIONS, SHOW JOBS, SHOW QUERIES, SHOW RANGE, SHOW RANGES,
SHOW ROLES, SHOW SCHEMAS, SHOW SEQUENCES, SHOW SESSION, SHOW SESSIONS,
SHOW STATISTICS, SHOW SYNTAX, SHOW TABLES, SHOW TRACE SHOW TRANSACTION, SHOW USERS
`,
	},
	//line sql.y: 3390
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 3391
		Category: hCfg,
		//line sql.y: 3392
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 3393
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 3414
	`SHOW STATISTICS`: {
		ShortDescription: `display table statistics (experimental)`,
		//line sql.y: 3415
		Category: hExperimental,
		//line sql.y: 3416
		Text: `SHOW STATISTICS [USING JSON] FOR TABLE <table_name>

Returns the available statistics for a table.
The statistics can include a histogram ID, which can
be used with SHOW HISTOGRAM.
If USING JSON is specified, the statistics and histograms
are encoded in JSON format.
`,
		//line sql.y: 3423
		SeeAlso: `SHOW HISTOGRAM
`,
	},
	//line sql.y: 3436
	`SHOW HISTOGRAM`: {
		ShortDescription: `display histogram (experimental)`,
		//line sql.y: 3437
		Category: hExperimental,
		//line sql.y: 3438
		Text: `SHOW HISTOGRAM <histogram_id>

Returns the data in the histogram with the
given ID (as returned by SHOW STATISTICS).
`,
		//line sql.y: 3442
		SeeAlso: `SHOW STATISTICS
`,
	},
	//line sql.y: 3455
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 3456
		Category: hCCL,
		//line sql.y: 3457
		Text: `SHOW BACKUP [SCHEMAS|FILES|RANGES] <location>
`,
		//line sql.y: 3458
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 3497
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 3498
		Category: hCfg,
		//line sql.y: 3499
		Text: `
SHOW CLUSTER SETTING <var>
SHOW [ PUBLIC | ALL ] CLUSTER SETTINGS
`,
		//line sql.y: 3502
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3528
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 3529
		Category: hDDL,
		//line sql.y: 3530
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 3531
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 3539
	`SHOW PARTITIONS`: {
		ShortDescription: `list partition information`,
		//line sql.y: 3540
		Category: hDDL,
		//line sql.y: 3541
		Text: `SHOW PARTITIONS FROM { TABLE <table> | INDEX <index> | DATABASE <database> }
`,
		//line sql.y: 3542
		SeeAlso: `WEBDOCS/show-partitions.html
`,
	},
	//line sql.y: 3562
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 3563
		Category: hDDL,
		//line sql.y: 3564
		Text: `SHOW DATABASES
`,
		//line sql.y: 3565
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 3573
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 3574
		Category: hPriv,
		//line sql.y: 3575
		Text: `
Show privilege grants:
  SHOW GRANTS [ON <targets...>] [FOR <users...>]
Show role grants:
  SHOW GRANTS ON ROLE [<roles...>] [FOR <grantees...>]

`,
		//line sql.y: 3581
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 3594
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 3595
		Category: hDDL,
		//line sql.y: 3596
		Text: `SHOW INDEXES FROM { <tablename> | DATABASE <database_name> } [WITH COMMENT]
`,
		//line sql.y: 3597
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 3627
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 3628
		Category: hDDL,
		//line sql.y: 3629
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 3630
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 3643
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 3644
		Category: hMisc,
		//line sql.y: 3645
		Text: `SHOW [ALL] [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 3646
		SeeAlso: `CANCEL QUERIES
`,
	},
	//line sql.y: 3667
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 3668
		Category: hMisc,
		//line sql.y: 3669
		Text: `
SHOW [AUTOMATIC] JOBS
SHOW JOB <jobid>
`,
		//line sql.y: 3672
		SeeAlso: `CANCEL JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3712
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 3713
		Category: hMisc,
		//line sql.y: 3714
		Text: `
SHOW [COMPACT] [KV] TRACE FOR SESSION
`,
		//line sql.y: 3716
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 3739
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 3740
		Category: hMisc,
		//line sql.y: 3741
		Text: `SHOW [ALL] [CLUSTER | LOCAL] SESSIONS
`,
		//line sql.y: 3742
		SeeAlso: `CANCEL SESSIONS
`,
	},
	//line sql.y: 3755
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 3756
		Category: hDDL,
		//line sql.y: 3757
		Text: `SHOW TABLES [FROM <databasename> [ . <schemaname> ] ] [WITH COMMENT]
`,
		//line sql.y: 3758
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 3790
	`SHOW SCHEMAS`: {
		ShortDescription: `list schemas`,
		//line sql.y: 3791
		Category: hDDL,
		//line sql.y: 3792
		Text: `SHOW SCHEMAS [FROM <databasename> ]
`,
	},
	//line sql.y: 3804
	`SHOW SEQUENCES`: {
		ShortDescription: `list sequences`,
		//line sql.y: 3805
		Category: hDDL,
		//line sql.y: 3806
		Text: `SHOW SEQUENCES [FROM <databasename> ]
`,
	},
	//line sql.y: 3818
	`SHOW SYNTAX`: {
		ShortDescription: `analyze SQL syntax`,
		//line sql.y: 3819
		Category: hMisc,
		//line sql.y: 3820
		Text: `SHOW SYNTAX <string>
`,
	},
	//line sql.y: 3829
	`SHOW SAVEPOINT`: {
		ShortDescription: `display current savepoint properties`,
		//line sql.y: 3830
		Category: hCfg,
		//line sql.y: 3831
		Text: `SHOW SAVEPOINT STATUS
`,
	},
	//line sql.y: 3839
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 3840
		Category: hCfg,
		//line sql.y: 3841
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 3842
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 3861
	`SHOW CREATE`: {
		ShortDescription: `display the CREATE statement for a table, sequence or view`,
		//line sql.y: 3862
		Category: hDDL,
		//line sql.y: 3863
		Text: `SHOW CREATE [ TABLE | SEQUENCE | VIEW ] <tablename>
`,
		//line sql.y: 3864
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 3882
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 3883
		Category: hPriv,
		//line sql.y: 3884
		Text: `SHOW USERS
`,
		//line sql.y: 3885
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 3893
	`SHOW ROLES`: {
		ShortDescription: `list defined roles`,
		//line sql.y: 3894
		Category: hPriv,
		//line sql.y: 3895
		Text: `SHOW ROLES
`,
		//line sql.y: 3896
		SeeAlso: `CREATE ROLE, ALTER ROLE, DROP ROLE
`,
	},
	//line sql.y: 3952
	`SHOW RANGE`: {
		ShortDescription: `show range information for a row`,
		//line sql.y: 3953
		Category: hMisc,
		//line sql.y: 3954
		Text: `
SHOW RANGE FROM TABLE <tablename> FOR ROW (row, value, ...)
SHOW RANGE FROM INDEX [ <tablename> @ ] <indexname> FOR ROW (row, value, ...)
`,
	},
	//line sql.y: 3975
	`SHOW RANGES`: {
		ShortDescription: `list ranges`,
		//line sql.y: 3976
		Category: hMisc,
		//line sql.y: 3977
		Text: `
SHOW RANGES FROM TABLE <tablename>
SHOW RANGES FROM INDEX [ <tablename> @ ] <indexname>
`,
	},
	//line sql.y: 4214
	`PAUSE JOBS`: {
		ShortDescription: `pause background jobs`,
		//line sql.y: 4215
		Category: hMisc,
		//line sql.y: 4216
		Text: `
PAUSE JOBS <selectclause>
PAUSE JOB <jobid>
`,
		//line sql.y: 4219
		SeeAlso: `SHOW JOBS, CANCEL JOBS, RESUME JOBS
`,
	},
	//line sql.y: 4236
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 4237
		Category: hDDL,
		//line sql.y: 4238
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
		//line sql.y: 4265
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 5010
	`CREATE SEQUENCE`: {
		ShortDescription: `create a new sequence`,
		//line sql.y: 5011
		Category: hDDL,
		//line sql.y: 5012
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
		//line sql.y: 5022
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 5087
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 5088
		Category: hDML,
		//line sql.y: 5089
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 5090
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 5108
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 5109
		Category: hPriv,
		//line sql.y: 5110
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] <OPTIONS...> ]
`,
		//line sql.y: 5111
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 5123
	`CREATE ROLE`: {
		ShortDescription: `define a new role`,
		//line sql.y: 5124
		Category: hPriv,
		//line sql.y: 5125
		Text: `CREATE ROLE [IF NOT EXISTS] <name> [ [WITH] <OPTIONS...> ]
`,
		//line sql.y: 5126
		SeeAlso: `ALTER ROLE, DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 5138
	`ALTER USER`: {
		ShortDescription: `alter a user`,
		//line sql.y: 5139
		Category: hPriv,
		//line sql.y: 5140
		Text: `ALTER USER <name> [WITH] <options...>
`,
		//line sql.y: 5141
		SeeAlso: `CREATE USER, DROP USER, SHOW USERS
`,
	},
	//line sql.y: 5153
	`ALTER ROLE`: {
		ShortDescription: `alter a role`,
		//line sql.y: 5154
		Category: hPriv,
		//line sql.y: 5155
		Text: `ALTER ROLE <name> [WITH] <options...>
`,
		//line sql.y: 5156
		SeeAlso: `CREATE ROLE, DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 5174
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 5175
		Category: hDDL,
		//line sql.y: 5176
		Text: `CREATE [TEMPORARY | TEMP] VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 5177
		SeeAlso: `CREATE TABLE, SHOW CREATE, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 5274
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 5275
		Category: hDDL,
		//line sql.y: 5276
		Text: `
CREATE [UNIQUE | INVERTED] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [USING HASH WITH BUCKET_COUNT = <shard_buckets>] [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 5284
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 5504
	`RELEASE`: {
		ShortDescription: `complete a sub-transaction`,
		//line sql.y: 5505
		Category: hTxn,
		//line sql.y: 5506
		Text: `RELEASE [SAVEPOINT] <savepoint name>
`,
		//line sql.y: 5507
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5515
	`RESUME JOBS`: {
		ShortDescription: `resume background jobs`,
		//line sql.y: 5516
		Category: hMisc,
		//line sql.y: 5517
		Text: `
RESUME JOBS <selectclause>
RESUME JOB <jobid>
`,
		//line sql.y: 5520
		SeeAlso: `SHOW JOBS, CANCEL JOBS, PAUSE JOBS
`,
	},
	//line sql.y: 5537
	`SAVEPOINT`: {
		ShortDescription: `start a sub-transaction`,
		//line sql.y: 5538
		Category: hTxn,
		//line sql.y: 5539
		Text: `SAVEPOINT <savepoint name>
`,
		//line sql.y: 5540
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5555
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 5556
		Category: hTxn,
		//line sql.y: 5557
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 5565
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 5578
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 5579
		Category: hTxn,
		//line sql.y: 5580
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 5583
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 5607
	`ROLLBACK`: {
		ShortDescription: `abort the current (sub-)transaction`,
		//line sql.y: 5608
		Category: hTxn,
		//line sql.y: 5609
		Text: `
ROLLBACK [TRANSACTION]
ROLLBACK [TRANSACTION] TO [SAVEPOINT] <savepoint name>
`,
		//line sql.y: 5612
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 5712
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 5713
		Category: hDDL,
		//line sql.y: 5714
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 5715
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 5784
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 5785
		Category: hDML,
		//line sql.y: 5786
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5791
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 5810
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 5811
		Category: hDML,
		//line sql.y: 5812
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 5816
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 5927
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 5928
		Category: hDML,
		//line sql.y: 5929
		Text: `
UPDATE <tablename> [[AS] <name>]
       SET ...
       [WHERE <expr>]
       [ORDER BY <exprs...>]
       [LIMIT <expr>]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5936
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 6161
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 6162
		Category: hDML,
		//line sql.y: 6163
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 6174
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 6175
		Category: hDML,
		//line sql.y: 6176
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
		//line sql.y: 6188
		SeeAlso: `WEBDOCS/select-clause.html
`,
	},
	//line sql.y: 6263
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 6264
		Category: hDML,
		//line sql.y: 6265
		Text: `TABLE <tablename>
`,
		//line sql.y: 6266
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 6588
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 6589
		Category: hDML,
		//line sql.y: 6590
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 6591
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 6700
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 6701
		Category: hDML,
		//line sql.y: 6702
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
		//line sql.y: 6724
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
