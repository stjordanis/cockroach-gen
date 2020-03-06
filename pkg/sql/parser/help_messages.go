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
	//line sql.y: 2752
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 2753
		Category: hMisc,
		//line sql.y: 2754
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 2755
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2786
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 2787
		Category: hMisc,
		//line sql.y: 2788
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 2789
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2819
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 2820
		Category: hMisc,
		//line sql.y: 2821
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 2822
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 2842
	`GRANT`: {
		ShortDescription: `define access privileges and role memberships`,
		//line sql.y: 2843
		Category: hPriv,
		//line sql.y: 2844
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
		//line sql.y: 2857
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 2873
	`REVOKE`: {
		ShortDescription: `remove access privileges and role memberships`,
		//line sql.y: 2874
		Category: hPriv,
		//line sql.y: 2875
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
		//line sql.y: 2888
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 2942
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 2943
		Category: hCfg,
		//line sql.y: 2944
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 2945
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2957
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 2958
		Category: hCfg,
		//line sql.y: 2959
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 2960
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 2969
	`USE`: {
		ShortDescription: `set the current database`,
		//line sql.y: 2970
		Category: hCfg,
		//line sql.y: 2971
		Text: `USE <dbname>

"USE <dbname>" is an alias for "SET [SESSION] database = <dbname>".
`,
		//line sql.y: 2974
		SeeAlso: `SET SESSION, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2995
	`SCRUB`: {
		ShortDescription: `run checks against databases or tables`,
		//line sql.y: 2996
		Category: hExperimental,
		//line sql.y: 2997
		Text: `
EXPERIMENTAL SCRUB TABLE <table> ...
EXPERIMENTAL SCRUB DATABASE <database>

The various checks that ca be run with SCRUB includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 3005
		SeeAlso: `SCRUB TABLE, SCRUB DATABASE
`,
	},
	//line sql.y: 3011
	`SCRUB DATABASE`: {
		ShortDescription: `run scrub checks on a database`,
		//line sql.y: 3012
		Category: hExperimental,
		//line sql.y: 3013
		Text: `
EXPERIMENTAL SCRUB DATABASE <database>
                            [AS OF SYSTEM TIME <expr>]

All scrub checks will be run on the database. This includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 3021
		SeeAlso: `SCRUB TABLE, SCRUB
`,
	},
	//line sql.y: 3029
	`SCRUB TABLE`: {
		ShortDescription: `run scrub checks on a table`,
		//line sql.y: 3030
		Category: hExperimental,
		//line sql.y: 3031
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
		//line sql.y: 3042
		SeeAlso: `SCRUB DATABASE, SRUB
`,
	},
	//line sql.y: 3097
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 3098
		Category: hCfg,
		//line sql.y: 3099
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 3100
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3121
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 3122
		Category: hCfg,
		//line sql.y: 3123
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
SET [SESSION] TRACING { TO | = } { on | off | cluster | local | kv | results } [,...]

`,
		//line sql.y: 3129
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 3146
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 3147
		Category: hTxn,
		//line sql.y: 3148
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 3155
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 3347
	`SHOW`: {
		//line sql.y: 3348
		Category: hGroup,
		//line sql.y: 3349
		Text: `
SHOW BACKUP, SHOW CLUSTER SETTING, SHOW COLUMNS, SHOW CONSTRAINTS,
SHOW CREATE, SHOW DATABASES, SHOW HISTOGRAM, SHOW INDEXES, SHOW
PARTITIONS, SHOW JOBS, SHOW QUERIES, SHOW RANGE, SHOW RANGES,
SHOW ROLES, SHOW SCHEMAS, SHOW SEQUENCES, SHOW SESSION, SHOW SESSIONS,
SHOW STATISTICS, SHOW SYNTAX, SHOW TABLES, SHOW TRACE SHOW TRANSACTION, SHOW USERS
`,
	},
	//line sql.y: 3386
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 3387
		Category: hCfg,
		//line sql.y: 3388
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 3389
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 3410
	`SHOW STATISTICS`: {
		ShortDescription: `display table statistics (experimental)`,
		//line sql.y: 3411
		Category: hExperimental,
		//line sql.y: 3412
		Text: `SHOW STATISTICS [USING JSON] FOR TABLE <table_name>

Returns the available statistics for a table.
The statistics can include a histogram ID, which can
be used with SHOW HISTOGRAM.
If USING JSON is specified, the statistics and histograms
are encoded in JSON format.
`,
		//line sql.y: 3419
		SeeAlso: `SHOW HISTOGRAM
`,
	},
	//line sql.y: 3432
	`SHOW HISTOGRAM`: {
		ShortDescription: `display histogram (experimental)`,
		//line sql.y: 3433
		Category: hExperimental,
		//line sql.y: 3434
		Text: `SHOW HISTOGRAM <histogram_id>

Returns the data in the histogram with the
given ID (as returned by SHOW STATISTICS).
`,
		//line sql.y: 3438
		SeeAlso: `SHOW STATISTICS
`,
	},
	//line sql.y: 3451
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 3452
		Category: hCCL,
		//line sql.y: 3453
		Text: `SHOW BACKUP [SCHEMAS|FILES|RANGES] <location>
`,
		//line sql.y: 3454
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 3493
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 3494
		Category: hCfg,
		//line sql.y: 3495
		Text: `
SHOW CLUSTER SETTING <var>
SHOW [ PUBLIC | ALL ] CLUSTER SETTINGS
`,
		//line sql.y: 3498
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 3524
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 3525
		Category: hDDL,
		//line sql.y: 3526
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 3527
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 3535
	`SHOW PARTITIONS`: {
		ShortDescription: `list partition information`,
		//line sql.y: 3536
		Category: hDDL,
		//line sql.y: 3537
		Text: `SHOW PARTITIONS FROM { TABLE <table> | INDEX <index> | DATABASE <database> }
`,
		//line sql.y: 3538
		SeeAlso: `WEBDOCS/show-partitions.html
`,
	},
	//line sql.y: 3558
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 3559
		Category: hDDL,
		//line sql.y: 3560
		Text: `SHOW DATABASES
`,
		//line sql.y: 3561
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 3569
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 3570
		Category: hPriv,
		//line sql.y: 3571
		Text: `
Show privilege grants:
  SHOW GRANTS [ON <targets...>] [FOR <users...>]
Show role grants:
  SHOW GRANTS ON ROLE [<roles...>] [FOR <grantees...>]

`,
		//line sql.y: 3577
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 3590
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 3591
		Category: hDDL,
		//line sql.y: 3592
		Text: `SHOW INDEXES FROM { <tablename> | DATABASE <database_name> } [WITH COMMENT]
`,
		//line sql.y: 3593
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 3623
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 3624
		Category: hDDL,
		//line sql.y: 3625
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 3626
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 3639
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 3640
		Category: hMisc,
		//line sql.y: 3641
		Text: `SHOW [ALL] [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 3642
		SeeAlso: `CANCEL QUERIES
`,
	},
	//line sql.y: 3663
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 3664
		Category: hMisc,
		//line sql.y: 3665
		Text: `
SHOW [AUTOMATIC] JOBS
SHOW JOB <jobid>
`,
		//line sql.y: 3668
		SeeAlso: `CANCEL JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3708
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 3709
		Category: hMisc,
		//line sql.y: 3710
		Text: `
SHOW [COMPACT] [KV] TRACE FOR SESSION
`,
		//line sql.y: 3712
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 3735
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 3736
		Category: hMisc,
		//line sql.y: 3737
		Text: `SHOW [ALL] [CLUSTER | LOCAL] SESSIONS
`,
		//line sql.y: 3738
		SeeAlso: `CANCEL SESSIONS
`,
	},
	//line sql.y: 3751
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 3752
		Category: hDDL,
		//line sql.y: 3753
		Text: `SHOW TABLES [FROM <databasename> [ . <schemaname> ] ] [WITH COMMENT]
`,
		//line sql.y: 3754
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 3786
	`SHOW SCHEMAS`: {
		ShortDescription: `list schemas`,
		//line sql.y: 3787
		Category: hDDL,
		//line sql.y: 3788
		Text: `SHOW SCHEMAS [FROM <databasename> ]
`,
	},
	//line sql.y: 3800
	`SHOW SEQUENCES`: {
		ShortDescription: `list sequences`,
		//line sql.y: 3801
		Category: hDDL,
		//line sql.y: 3802
		Text: `SHOW SEQUENCES [FROM <databasename> ]
`,
	},
	//line sql.y: 3814
	`SHOW SYNTAX`: {
		ShortDescription: `analyze SQL syntax`,
		//line sql.y: 3815
		Category: hMisc,
		//line sql.y: 3816
		Text: `SHOW SYNTAX <string>
`,
	},
	//line sql.y: 3825
	`SHOW SAVEPOINT`: {
		ShortDescription: `display current savepoint properties`,
		//line sql.y: 3826
		Category: hCfg,
		//line sql.y: 3827
		Text: `SHOW SAVEPOINT STATUS
`,
	},
	//line sql.y: 3836
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 3837
		Category: hCfg,
		//line sql.y: 3838
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 3839
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 3858
	`SHOW CREATE`: {
		ShortDescription: `display the CREATE statement for a table, sequence or view`,
		//line sql.y: 3859
		Category: hDDL,
		//line sql.y: 3860
		Text: `SHOW CREATE [ TABLE | SEQUENCE | VIEW ] <tablename>
`,
		//line sql.y: 3861
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 3879
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 3880
		Category: hPriv,
		//line sql.y: 3881
		Text: `SHOW USERS
`,
		//line sql.y: 3882
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 3890
	`SHOW ROLES`: {
		ShortDescription: `list defined roles`,
		//line sql.y: 3891
		Category: hPriv,
		//line sql.y: 3892
		Text: `SHOW ROLES
`,
		//line sql.y: 3893
		SeeAlso: `CREATE ROLE, ALTER ROLE, DROP ROLE
`,
	},
	//line sql.y: 3949
	`SHOW RANGE`: {
		ShortDescription: `show range information for a row`,
		//line sql.y: 3950
		Category: hMisc,
		//line sql.y: 3951
		Text: `
SHOW RANGE FROM TABLE <tablename> FOR ROW (row, value, ...)
SHOW RANGE FROM INDEX [ <tablename> @ ] <indexname> FOR ROW (row, value, ...)
`,
	},
	//line sql.y: 3972
	`SHOW RANGES`: {
		ShortDescription: `list ranges`,
		//line sql.y: 3973
		Category: hMisc,
		//line sql.y: 3974
		Text: `
SHOW RANGES FROM TABLE <tablename>
SHOW RANGES FROM INDEX [ <tablename> @ ] <indexname>
`,
	},
	//line sql.y: 4211
	`PAUSE JOBS`: {
		ShortDescription: `pause background jobs`,
		//line sql.y: 4212
		Category: hMisc,
		//line sql.y: 4213
		Text: `
PAUSE JOBS <selectclause>
PAUSE JOB <jobid>
`,
		//line sql.y: 4216
		SeeAlso: `SHOW JOBS, CANCEL JOBS, RESUME JOBS
`,
	},
	//line sql.y: 4233
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 4234
		Category: hDDL,
		//line sql.y: 4235
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
		//line sql.y: 4262
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 5007
	`CREATE SEQUENCE`: {
		ShortDescription: `create a new sequence`,
		//line sql.y: 5008
		Category: hDDL,
		//line sql.y: 5009
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
		//line sql.y: 5019
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 5084
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 5085
		Category: hDML,
		//line sql.y: 5086
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 5087
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 5105
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 5106
		Category: hPriv,
		//line sql.y: 5107
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] <OPTIONS...> ]
`,
		//line sql.y: 5108
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 5120
	`CREATE ROLE`: {
		ShortDescription: `define a new role`,
		//line sql.y: 5121
		Category: hPriv,
		//line sql.y: 5122
		Text: `CREATE ROLE [IF NOT EXISTS] <name> [ [WITH] <OPTIONS...> ]
`,
		//line sql.y: 5123
		SeeAlso: `ALTER ROLE, DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 5135
	`ALTER USER`: {
		ShortDescription: `alter a user`,
		//line sql.y: 5136
		Category: hPriv,
		//line sql.y: 5137
		Text: `ALTER USER <name> [WITH] <options...>
`,
		//line sql.y: 5138
		SeeAlso: `CREATE USER, DROP USER, SHOW USERS
`,
	},
	//line sql.y: 5150
	`ALTER ROLE`: {
		ShortDescription: `alter a role`,
		//line sql.y: 5151
		Category: hPriv,
		//line sql.y: 5152
		Text: `ALTER ROLE <name> [WITH] <options...>
`,
		//line sql.y: 5153
		SeeAlso: `CREATE ROLE, DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 5171
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 5172
		Category: hDDL,
		//line sql.y: 5173
		Text: `CREATE [TEMPORARY | TEMP] VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 5174
		SeeAlso: `CREATE TABLE, SHOW CREATE, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 5252
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 5253
		Category: hDDL,
		//line sql.y: 5254
		Text: `
CREATE [UNIQUE | INVERTED] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [USING HASH WITH BUCKET_COUNT = <shard_buckets>] [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 5262
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 5482
	`RELEASE`: {
		ShortDescription: `complete a sub-transaction`,
		//line sql.y: 5483
		Category: hTxn,
		//line sql.y: 5484
		Text: `RELEASE [SAVEPOINT] <savepoint name>
`,
		//line sql.y: 5485
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5493
	`RESUME JOBS`: {
		ShortDescription: `resume background jobs`,
		//line sql.y: 5494
		Category: hMisc,
		//line sql.y: 5495
		Text: `
RESUME JOBS <selectclause>
RESUME JOB <jobid>
`,
		//line sql.y: 5498
		SeeAlso: `SHOW JOBS, CANCEL JOBS, PAUSE JOBS
`,
	},
	//line sql.y: 5515
	`SAVEPOINT`: {
		ShortDescription: `start a sub-transaction`,
		//line sql.y: 5516
		Category: hTxn,
		//line sql.y: 5517
		Text: `SAVEPOINT <savepoint name>
`,
		//line sql.y: 5518
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 5533
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 5534
		Category: hTxn,
		//line sql.y: 5535
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 5543
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 5556
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 5557
		Category: hTxn,
		//line sql.y: 5558
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 5561
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 5585
	`ROLLBACK`: {
		ShortDescription: `abort the current (sub-)transaction`,
		//line sql.y: 5586
		Category: hTxn,
		//line sql.y: 5587
		Text: `
ROLLBACK [TRANSACTION]
ROLLBACK [TRANSACTION] TO [SAVEPOINT] <savepoint name>
`,
		//line sql.y: 5590
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 5690
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 5691
		Category: hDDL,
		//line sql.y: 5692
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 5693
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 5762
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 5763
		Category: hDML,
		//line sql.y: 5764
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5769
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 5788
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 5789
		Category: hDML,
		//line sql.y: 5790
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 5794
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 5905
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 5906
		Category: hDML,
		//line sql.y: 5907
		Text: `
UPDATE <tablename> [[AS] <name>]
       SET ...
       [WHERE <expr>]
       [ORDER BY <exprs...>]
       [LIMIT <expr>]
       [RETURNING <exprs...>]
`,
		//line sql.y: 5914
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 6139
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 6140
		Category: hDML,
		//line sql.y: 6141
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 6152
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 6153
		Category: hDML,
		//line sql.y: 6154
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
		//line sql.y: 6166
		SeeAlso: `WEBDOCS/select-clause.html
`,
	},
	//line sql.y: 6241
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 6242
		Category: hDML,
		//line sql.y: 6243
		Text: `TABLE <tablename>
`,
		//line sql.y: 6244
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 6566
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 6567
		Category: hDML,
		//line sql.y: 6568
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 6569
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 6678
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 6679
		Category: hDML,
		//line sql.y: 6680
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
		//line sql.y: 6702
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
