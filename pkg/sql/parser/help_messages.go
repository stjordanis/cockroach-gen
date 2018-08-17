// Code generated by help.awk. DO NOT EDIT.
// GENERATED FILE DO NOT EDIT

package parser

var helpMessages = map[string]HelpMessageBody{
	//line sql.y: 1090
	`ALTER`: {
		//line sql.y: 1091
		Category: hGroup,
		//line sql.y: 1092
		Text: `ALTER TABLE, ALTER INDEX, ALTER VIEW, ALTER SEQUENCE, ALTER DATABASE, ALTER USER
`,
	},
	//line sql.y: 1106
	`ALTER TABLE`: {
		ShortDescription: `change the definition of a table`,
		//line sql.y: 1107
		Category: hDDL,
		//line sql.y: 1108
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
  ALTER TABLE ... RENAME TO <newname>
  ALTER TABLE ... RENAME [COLUMN] <colname> TO <newname>
  ALTER TABLE ... VALIDATE CONSTRAINT <constraintname>
  ALTER TABLE ... SPLIT AT <selectclause>
  ALTER TABLE ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

Column qualifiers:
  [CONSTRAINT <constraintname>] {NULL | NOT NULL | UNIQUE | PRIMARY KEY | CHECK (<expr>) | DEFAULT <expr>}
  FAMILY <familyname>, CREATE [IF NOT EXISTS] FAMILY [<familyname>]
  REFERENCES <tablename> [( <colnames...> )]
  COLLATE <collationname>

`,
		//line sql.y: 1132
		SeeAlso: `WEBDOCS/alter-table.html
`,
	},
	//line sql.y: 1145
	`ALTER VIEW`: {
		ShortDescription: `change the definition of a view`,
		//line sql.y: 1146
		Category: hDDL,
		//line sql.y: 1147
		Text: `
ALTER VIEW [IF EXISTS] <name> RENAME TO <newname>
`,
		//line sql.y: 1149
		SeeAlso: `WEBDOCS/alter-view.html
`,
	},
	//line sql.y: 1156
	`ALTER SEQUENCE`: {
		ShortDescription: `change the definition of a sequence`,
		//line sql.y: 1157
		Category: hDDL,
		//line sql.y: 1158
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
	//line sql.y: 1181
	`ALTER USER`: {
		ShortDescription: `change user properties`,
		//line sql.y: 1182
		Category: hPriv,
		//line sql.y: 1183
		Text: `
ALTER USER [IF EXISTS] <name> WITH PASSWORD <password>
`,
		//line sql.y: 1185
		SeeAlso: `CREATE USER
`,
	},
	//line sql.y: 1190
	`ALTER DATABASE`: {
		ShortDescription: `change the definition of a database`,
		//line sql.y: 1191
		Category: hDDL,
		//line sql.y: 1192
		Text: `
ALTER DATABASE <name> RENAME TO <newname>
`,
		//line sql.y: 1194
		SeeAlso: `WEBDOCS/alter-database.html
`,
	},
	//line sql.y: 1205
	`ALTER INDEX`: {
		ShortDescription: `change the definition of an index`,
		//line sql.y: 1206
		Category: hDDL,
		//line sql.y: 1207
		Text: `
ALTER INDEX [IF EXISTS] <idxname> <command>

Commands:
  ALTER INDEX ... RENAME TO <newname>
  ALTER INDEX ... SPLIT AT <selectclause>
  ALTER INDEX ... SCATTER [ FROM ( <exprs...> ) TO ( <exprs...> ) ]

`,
		//line sql.y: 1215
		SeeAlso: `WEBDOCS/alter-index.html
`,
	},
	//line sql.y: 1568
	`BACKUP`: {
		ShortDescription: `back up data to external storage`,
		//line sql.y: 1569
		Category: hCCL,
		//line sql.y: 1570
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
		//line sql.y: 1587
		SeeAlso: `RESTORE, WEBDOCS/backup.html
`,
	},
	//line sql.y: 1595
	`RESTORE`: {
		ShortDescription: `restore data from external storage`,
		//line sql.y: 1596
		Category: hCCL,
		//line sql.y: 1597
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
		//line sql.y: 1613
		SeeAlso: `BACKUP, WEBDOCS/restore.html
`,
	},
	//line sql.y: 1631
	`IMPORT`: {
		ShortDescription: `load data from file in a distributed manner`,
		//line sql.y: 1632
		Category: hCCL,
		//line sql.y: 1633
		Text: `
IMPORT [ TABLE <tablename> FROM ]
       <format> ( <datafile> )
       [ WITH <option> [= <value>] [, ...] ]

IMPORT TABLE <tablename>
       { ( <elements> ) | CREATE USING <schemafile> }
       <format>
       DATA ( <datafile> [, ...] )
       [ WITH <option> [= <value>] [, ...] ]

Formats:
   CSV
   MYSQLOUTFILE
   MYSQLDUMP (mysqldump's SQL output)
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
		//line sql.y: 1659
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 1680
	`EXPORT`: {
		ShortDescription: `export data to file in a distributed manner`,
		//line sql.y: 1681
		Category: hCCL,
		//line sql.y: 1682
		Text: `
EXPORT <format> (<datafile> [WITH <option> [= value] [,...]]) FROM <query>

Formats:
   CSV

Options:
   delimiter = '...'   [CSV-specific]

`,
		//line sql.y: 1691
		SeeAlso: `SELECT
`,
	},
	//line sql.y: 1778
	`CANCEL`: {
		//line sql.y: 1779
		Category: hGroup,
		//line sql.y: 1780
		Text: `CANCEL JOBS, CANCEL QUERIES, CANCEL SESSIONS
`,
	},
	//line sql.y: 1787
	`CANCEL JOBS`: {
		ShortDescription: `cancel background jobs`,
		//line sql.y: 1788
		Category: hMisc,
		//line sql.y: 1789
		Text: `
CANCEL JOBS <selectclause>
CANCEL JOB <jobid>
`,
		//line sql.y: 1792
		SeeAlso: `SHOW JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 1810
	`CANCEL QUERIES`: {
		ShortDescription: `cancel running queries`,
		//line sql.y: 1811
		Category: hMisc,
		//line sql.y: 1812
		Text: `
CANCEL QUERIES [IF EXISTS] <selectclause>
CANCEL QUERY [IF EXISTS] <expr>
`,
		//line sql.y: 1815
		SeeAlso: `SHOW QUERIES
`,
	},
	//line sql.y: 1846
	`CANCEL SESSIONS`: {
		ShortDescription: `cancel open sessions`,
		//line sql.y: 1847
		Category: hMisc,
		//line sql.y: 1848
		Text: `
CANCEL SESSIONS [IF EXISTS] <selectclause>
CANCEL SESSION [IF EXISTS] <sessionid>
`,
		//line sql.y: 1851
		SeeAlso: `SHOW SESSIONS
`,
	},
	//line sql.y: 1898
	`CREATE`: {
		//line sql.y: 1899
		Category: hGroup,
		//line sql.y: 1900
		Text: `
CREATE DATABASE, CREATE TABLE, CREATE INDEX, CREATE TABLE AS,
CREATE USER, CREATE VIEW, CREATE SEQUENCE, CREATE STATISTICS,
CREATE ROLE
`,
	},
	//line sql.y: 1923
	`CREATE STATISTICS`: {
		ShortDescription: `create a new table statistic`,
		//line sql.y: 1924
		Category: hMisc,
		//line sql.y: 1925
		Text: `
CREATE STATISTICS <statisticname>
  ON <colname> [, ...]
  FROM <tablename>
`,
	},
	//line sql.y: 1981
	`DELETE`: {
		ShortDescription: `delete rows from a table`,
		//line sql.y: 1982
		Category: hDML,
		//line sql.y: 1983
		Text: `DELETE FROM <tablename> [WHERE <expr>]
              [ORDER BY <exprs...>]
              [LIMIT <expr>]
              [RETURNING <exprs...>]
`,
		//line sql.y: 1987
		SeeAlso: `WEBDOCS/delete.html
`,
	},
	//line sql.y: 2002
	`DISCARD`: {
		ShortDescription: `reset the session to its initial state`,
		//line sql.y: 2003
		Category: hCfg,
		//line sql.y: 2004
		Text: `DISCARD ALL
`,
	},
	//line sql.y: 2016
	`DROP`: {
		//line sql.y: 2017
		Category: hGroup,
		//line sql.y: 2018
		Text: `
DROP DATABASE, DROP INDEX, DROP TABLE, DROP VIEW, DROP SEQUENCE,
DROP USER, DROP ROLE
`,
	},
	//line sql.y: 2034
	`DROP VIEW`: {
		ShortDescription: `remove a view`,
		//line sql.y: 2035
		Category: hDDL,
		//line sql.y: 2036
		Text: `DROP VIEW [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2037
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2049
	`DROP SEQUENCE`: {
		ShortDescription: `remove a sequence`,
		//line sql.y: 2050
		Category: hDDL,
		//line sql.y: 2051
		Text: `DROP SEQUENCE [IF EXISTS] <sequenceName> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2052
		SeeAlso: `DROP
`,
	},
	//line sql.y: 2064
	`DROP TABLE`: {
		ShortDescription: `remove a table`,
		//line sql.y: 2065
		Category: hDDL,
		//line sql.y: 2066
		Text: `DROP TABLE [IF EXISTS] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2067
		SeeAlso: `WEBDOCS/drop-table.html
`,
	},
	//line sql.y: 2079
	`DROP INDEX`: {
		ShortDescription: `remove an index`,
		//line sql.y: 2080
		Category: hDDL,
		//line sql.y: 2081
		Text: `DROP INDEX [IF EXISTS] <idxname> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 2082
		SeeAlso: `WEBDOCS/drop-index.html
`,
	},
	//line sql.y: 2102
	`DROP DATABASE`: {
		ShortDescription: `remove a database`,
		//line sql.y: 2103
		Category: hDDL,
		//line sql.y: 2104
		Text: `DROP DATABASE [IF EXISTS] <databasename> [CASCADE | RESTRICT]
`,
		//line sql.y: 2105
		SeeAlso: `WEBDOCS/drop-database.html
`,
	},
	//line sql.y: 2125
	`DROP USER`: {
		ShortDescription: `remove a user`,
		//line sql.y: 2126
		Category: hPriv,
		//line sql.y: 2127
		Text: `DROP USER [IF EXISTS] <user> [, ...]
`,
		//line sql.y: 2128
		SeeAlso: `CREATE USER, SHOW USERS
`,
	},
	//line sql.y: 2140
	`DROP ROLE`: {
		ShortDescription: `remove a role`,
		//line sql.y: 2141
		Category: hPriv,
		//line sql.y: 2142
		Text: `DROP ROLE [IF EXISTS] <role> [, ...]
`,
		//line sql.y: 2143
		SeeAlso: `CREATE ROLE, SHOW ROLES
`,
	},
	//line sql.y: 2165
	`EXPLAIN`: {
		ShortDescription: `show the logical plan of a query`,
		//line sql.y: 2166
		Category: hMisc,
		//line sql.y: 2167
		Text: `
EXPLAIN <statement>
EXPLAIN ([PLAN ,] <planoptions...> ) <statement>
EXPLAIN [ANALYZE] (DISTSQL) <statement>

Explainable statements:
    SELECT, CREATE, DROP, ALTER, INSERT, UPSERT, UPDATE, DELETE,
    SHOW, EXPLAIN, EXECUTE

Plan options:
    TYPES, VERBOSE

`,
		//line sql.y: 2179
		SeeAlso: `WEBDOCS/explain.html
`,
	},
	//line sql.y: 2244
	`PREPARE`: {
		ShortDescription: `prepare a statement for later execution`,
		//line sql.y: 2245
		Category: hMisc,
		//line sql.y: 2246
		Text: `PREPARE <name> [ ( <types...> ) ] AS <query>
`,
		//line sql.y: 2247
		SeeAlso: `EXECUTE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2269
	`EXECUTE`: {
		ShortDescription: `execute a statement prepared previously`,
		//line sql.y: 2270
		Category: hMisc,
		//line sql.y: 2271
		Text: `EXECUTE <name> [ ( <exprs...> ) ]
`,
		//line sql.y: 2272
		SeeAlso: `PREPARE, DEALLOCATE, DISCARD
`,
	},
	//line sql.y: 2295
	`DEALLOCATE`: {
		ShortDescription: `remove a prepared statement`,
		//line sql.y: 2296
		Category: hMisc,
		//line sql.y: 2297
		Text: `DEALLOCATE [PREPARE] { <name> | ALL }
`,
		//line sql.y: 2298
		SeeAlso: `PREPARE, EXECUTE, DISCARD
`,
	},
	//line sql.y: 2318
	`GRANT`: {
		ShortDescription: `define access privileges and role memberships`,
		//line sql.y: 2319
		Category: hPriv,
		//line sql.y: 2320
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
		//line sql.y: 2333
		SeeAlso: `REVOKE, WEBDOCS/grant.html
`,
	},
	//line sql.y: 2349
	`REVOKE`: {
		ShortDescription: `remove access privileges and role memberships`,
		//line sql.y: 2350
		Category: hPriv,
		//line sql.y: 2351
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
		//line sql.y: 2364
		SeeAlso: `GRANT, WEBDOCS/revoke.html
`,
	},
	//line sql.y: 2419
	`RESET`: {
		ShortDescription: `reset a session variable to its default value`,
		//line sql.y: 2420
		Category: hCfg,
		//line sql.y: 2421
		Text: `RESET [SESSION] <var>
`,
		//line sql.y: 2422
		SeeAlso: `RESET CLUSTER SETTING, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2434
	`RESET CLUSTER SETTING`: {
		ShortDescription: `reset a cluster setting to its default value`,
		//line sql.y: 2435
		Category: hCfg,
		//line sql.y: 2436
		Text: `RESET CLUSTER SETTING <var>
`,
		//line sql.y: 2437
		SeeAlso: `SET CLUSTER SETTING, RESET
`,
	},
	//line sql.y: 2446
	`USE`: {
		ShortDescription: `set the current database`,
		//line sql.y: 2447
		Category: hCfg,
		//line sql.y: 2448
		Text: `USE <dbname>

"USE <dbname>" is an alias for "SET [SESSION] database = <dbname>".
`,
		//line sql.y: 2451
		SeeAlso: `SET SESSION, WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2468
	`SCRUB`: {
		ShortDescription: `run checks against databases or tables`,
		//line sql.y: 2469
		Category: hExperimental,
		//line sql.y: 2470
		Text: `
EXPERIMENTAL SCRUB TABLE <table> ...
EXPERIMENTAL SCRUB DATABASE <database>

The various checks that ca be run with SCRUB includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2478
		SeeAlso: `SCRUB TABLE, SCRUB DATABASE
`,
	},
	//line sql.y: 2484
	`SCRUB DATABASE`: {
		ShortDescription: `run scrub checks on a database`,
		//line sql.y: 2485
		Category: hExperimental,
		//line sql.y: 2486
		Text: `
EXPERIMENTAL SCRUB DATABASE <database>
                            [AS OF SYSTEM TIME <expr>]

All scrub checks will be run on the database. This includes:
  - Physical table data (encoding)
  - Secondary index integrity
  - Constraint integrity (NOT NULL, CHECK, FOREIGN KEY, UNIQUE)
`,
		//line sql.y: 2494
		SeeAlso: `SCRUB TABLE, SCRUB
`,
	},
	//line sql.y: 2502
	`SCRUB TABLE`: {
		ShortDescription: `run scrub checks on a table`,
		//line sql.y: 2503
		Category: hExperimental,
		//line sql.y: 2504
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
		//line sql.y: 2515
		SeeAlso: `SCRUB DATABASE, SRUB
`,
	},
	//line sql.y: 2570
	`SET CLUSTER SETTING`: {
		ShortDescription: `change a cluster setting`,
		//line sql.y: 2571
		Category: hCfg,
		//line sql.y: 2572
		Text: `SET CLUSTER SETTING <var> { TO | = } <value>
`,
		//line sql.y: 2573
		SeeAlso: `SHOW CLUSTER SETTING, RESET CLUSTER SETTING, SET SESSION,
WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2594
	`SET SESSION`: {
		ShortDescription: `change a session variable`,
		//line sql.y: 2595
		Category: hCfg,
		//line sql.y: 2596
		Text: `
SET [SESSION] <var> { TO | = } <values...>
SET [SESSION] TIME ZONE <tz>
SET [SESSION] CHARACTERISTICS AS TRANSACTION ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
SET [SESSION] TRACING { TO | = } { on | off | cluster | local | kv | results } [,...]

`,
		//line sql.y: 2602
		SeeAlso: `SHOW SESSION, RESET, DISCARD, SHOW, SET CLUSTER SETTING, SET TRANSACTION,
WEBDOCS/set-vars.html
`,
	},
	//line sql.y: 2619
	`SET TRANSACTION`: {
		ShortDescription: `configure the transaction settings`,
		//line sql.y: 2620
		Category: hTxn,
		//line sql.y: 2621
		Text: `
SET [SESSION] TRANSACTION <txnparameters...>

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 2628
		SeeAlso: `SHOW TRANSACTION, SET SESSION,
WEBDOCS/set-transaction.html
`,
	},
	//line sql.y: 2805
	`SHOW`: {
		//line sql.y: 2806
		Category: hGroup,
		//line sql.y: 2807
		Text: `
SHOW SESSION, SHOW CLUSTER SETTING, SHOW DATABASES, SHOW TABLES, SHOW COLUMNS, SHOW INDEXES,
SHOW CONSTRAINTS, SHOW CREATE, SHOW USERS,
SHOW TRANSACTION, SHOW BACKUP, SHOW JOBS, SHOW QUERIES, SHOW ROLES, SHOW SESSIONS, SHOW SYNTAX,
SHOW TRACE
`,
	},
	//line sql.y: 2839
	`SHOW SESSION`: {
		ShortDescription: `display session variables`,
		//line sql.y: 2840
		Category: hCfg,
		//line sql.y: 2841
		Text: `SHOW [SESSION] { <var> | ALL }
`,
		//line sql.y: 2842
		SeeAlso: `WEBDOCS/show-vars.html
`,
	},
	//line sql.y: 2863
	`SHOW STATISTICS`: {
		ShortDescription: `display table statistics`,
		//line sql.y: 2864
		Category: hMisc,
		//line sql.y: 2865
		Text: `SHOW STATISTICS [USING JSON] FOR TABLE <table_name>

Returns the available statistics for a table.
The statistics can include a histogram ID, which can
be used with SHOW HISTOGRAM.
If USING JSON is specified, the statistics and histograms
are encoded in JSON format.
`,
		//line sql.y: 2872
		SeeAlso: `SHOW HISTOGRAM
`,
	},
	//line sql.y: 2884
	`SHOW HISTOGRAM`: {
		ShortDescription: `display histogram`,
		//line sql.y: 2885
		Category: hMisc,
		//line sql.y: 2886
		Text: `SHOW HISTOGRAM <histogram_id>

Returns the data in the histogram with the
given ID (as returned by SHOW STATISTICS).
`,
		//line sql.y: 2890
		SeeAlso: `SHOW STATISTICS
`,
	},
	//line sql.y: 2903
	`SHOW BACKUP`: {
		ShortDescription: `list backup contents`,
		//line sql.y: 2904
		Category: hCCL,
		//line sql.y: 2905
		Text: `SHOW BACKUP [FILES|RANGES] <location>
`,
		//line sql.y: 2906
		SeeAlso: `WEBDOCS/show-backup.html
`,
	},
	//line sql.y: 2933
	`SHOW CLUSTER SETTING`: {
		ShortDescription: `display cluster settings`,
		//line sql.y: 2934
		Category: hCfg,
		//line sql.y: 2935
		Text: `
SHOW CLUSTER SETTING <var>
SHOW ALL CLUSTER SETTINGS
`,
		//line sql.y: 2938
		SeeAlso: `WEBDOCS/cluster-settings.html
`,
	},
	//line sql.y: 2955
	`SHOW COLUMNS`: {
		ShortDescription: `list columns in relation`,
		//line sql.y: 2956
		Category: hDDL,
		//line sql.y: 2957
		Text: `SHOW COLUMNS FROM <tablename>
`,
		//line sql.y: 2958
		SeeAlso: `WEBDOCS/show-columns.html
`,
	},
	//line sql.y: 2966
	`SHOW DATABASES`: {
		ShortDescription: `list databases`,
		//line sql.y: 2967
		Category: hDDL,
		//line sql.y: 2968
		Text: `SHOW DATABASES
`,
		//line sql.y: 2969
		SeeAlso: `WEBDOCS/show-databases.html
`,
	},
	//line sql.y: 2977
	`SHOW GRANTS`: {
		ShortDescription: `list grants`,
		//line sql.y: 2978
		Category: hPriv,
		//line sql.y: 2979
		Text: `
Show privilege grants:
  SHOW GRANTS [ON <targets...>] [FOR <users...>]
Show role grants:
  SHOW GRANTS ON ROLE [<roles...>] [FOR <grantees...>]

`,
		//line sql.y: 2985
		SeeAlso: `WEBDOCS/show-grants.html
`,
	},
	//line sql.y: 2998
	`SHOW INDEXES`: {
		ShortDescription: `list indexes`,
		//line sql.y: 2999
		Category: hDDL,
		//line sql.y: 3000
		Text: `SHOW INDEXES FROM <tablename>
`,
		//line sql.y: 3001
		SeeAlso: `WEBDOCS/show-index.html
`,
	},
	//line sql.y: 3019
	`SHOW CONSTRAINTS`: {
		ShortDescription: `list constraints`,
		//line sql.y: 3020
		Category: hDDL,
		//line sql.y: 3021
		Text: `SHOW CONSTRAINTS FROM <tablename>
`,
		//line sql.y: 3022
		SeeAlso: `WEBDOCS/show-constraints.html
`,
	},
	//line sql.y: 3035
	`SHOW QUERIES`: {
		ShortDescription: `list running queries`,
		//line sql.y: 3036
		Category: hMisc,
		//line sql.y: 3037
		Text: `SHOW [CLUSTER | LOCAL] QUERIES
`,
		//line sql.y: 3038
		SeeAlso: `CANCEL QUERIES
`,
	},
	//line sql.y: 3054
	`SHOW JOBS`: {
		ShortDescription: `list background jobs`,
		//line sql.y: 3055
		Category: hMisc,
		//line sql.y: 3056
		Text: `SHOW JOBS
`,
		//line sql.y: 3057
		SeeAlso: `CANCEL JOBS, PAUSE JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3065
	`SHOW TRACE`: {
		ShortDescription: `display an execution trace`,
		//line sql.y: 3066
		Category: hMisc,
		//line sql.y: 3067
		Text: `
SHOW [COMPACT] [KV] TRACE FOR SESSION
`,
		//line sql.y: 3069
		SeeAlso: `EXPLAIN
`,
	},
	//line sql.y: 3092
	`SHOW SESSIONS`: {
		ShortDescription: `list open client sessions`,
		//line sql.y: 3093
		Category: hMisc,
		//line sql.y: 3094
		Text: `SHOW [CLUSTER | LOCAL] SESSIONS
`,
		//line sql.y: 3095
		SeeAlso: `CANCEL SESSIONS
`,
	},
	//line sql.y: 3111
	`SHOW TABLES`: {
		ShortDescription: `list tables`,
		//line sql.y: 3112
		Category: hDDL,
		//line sql.y: 3113
		Text: `SHOW TABLES [FROM <databasename> [ . <schemaname> ] ]
`,
		//line sql.y: 3114
		SeeAlso: `WEBDOCS/show-tables.html
`,
	},
	//line sql.y: 3140
	`SHOW SCHEMAS`: {
		ShortDescription: `list schemas`,
		//line sql.y: 3141
		Category: hDDL,
		//line sql.y: 3142
		Text: `SHOW SCHEMAS [FROM <databasename> ]
`,
	},
	//line sql.y: 3154
	`SHOW SYNTAX`: {
		ShortDescription: `analyze SQL syntax`,
		//line sql.y: 3155
		Category: hMisc,
		//line sql.y: 3156
		Text: `SHOW SYNTAX <string>
`,
	},
	//line sql.y: 3165
	`SHOW TRANSACTION`: {
		ShortDescription: `display current transaction properties`,
		//line sql.y: 3166
		Category: hCfg,
		//line sql.y: 3167
		Text: `SHOW TRANSACTION {ISOLATION LEVEL | PRIORITY | STATUS}
`,
		//line sql.y: 3168
		SeeAlso: `WEBDOCS/show-transaction.html
`,
	},
	//line sql.y: 3187
	`SHOW CREATE`: {
		ShortDescription: `display the CREATE statement for a table, sequence or view`,
		//line sql.y: 3188
		Category: hDDL,
		//line sql.y: 3189
		Text: `SHOW CREATE [ TABLE | SEQUENCE | VIEW ] <tablename>
`,
		//line sql.y: 3190
		SeeAlso: `WEBDOCS/show-create-table.html
`,
	},
	//line sql.y: 3208
	`SHOW USERS`: {
		ShortDescription: `list defined users`,
		//line sql.y: 3209
		Category: hPriv,
		//line sql.y: 3210
		Text: `SHOW USERS
`,
		//line sql.y: 3211
		SeeAlso: `CREATE USER, DROP USER, WEBDOCS/show-users.html
`,
	},
	//line sql.y: 3219
	`SHOW ROLES`: {
		ShortDescription: `list defined roles`,
		//line sql.y: 3220
		Category: hPriv,
		//line sql.y: 3221
		Text: `SHOW ROLES
`,
		//line sql.y: 3222
		SeeAlso: `CREATE ROLE, DROP ROLE
`,
	},
	//line sql.y: 3274
	`SHOW RANGES`: {
		ShortDescription: `list ranges`,
		//line sql.y: 3275
		Category: hMisc,
		//line sql.y: 3276
		Text: `
SHOW EXPERIMENTAL_RANGES FROM TABLE <tablename>
SHOW EXPERIMENTAL_RANGES FROM INDEX [ <tablename> @ ] <indexname>
`,
	},
	//line sql.y: 3512
	`PAUSE JOBS`: {
		ShortDescription: `pause background jobs`,
		//line sql.y: 3513
		Category: hMisc,
		//line sql.y: 3514
		Text: `
PAUSE JOBS <selectclause>
PAUSE JOB <jobid>
`,
		//line sql.y: 3517
		SeeAlso: `SHOW JOBS, CANCEL JOBS, RESUME JOBS
`,
	},
	//line sql.y: 3534
	`CREATE TABLE`: {
		ShortDescription: `create a new table`,
		//line sql.y: 3535
		Category: hDDL,
		//line sql.y: 3536
		Text: `
CREATE TABLE [IF NOT EXISTS] <tablename> ( <elements...> ) [<interleave>]
CREATE TABLE [IF NOT EXISTS] <tablename> [( <colnames...> )] AS <source>

Table elements:
   <name> <type> [<qualifiers...>]
   [UNIQUE | INVERTED] INDEX [<name>] ( <colname> [ASC | DESC] [, ...] )
                           [STORING ( <colnames...> )] [<interleave>]
   FAMILY [<name>] ( <colnames...> )
   [CONSTRAINT <name>] <constraint>

Table constraints:
   PRIMARY KEY ( <colnames...> )
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
		//line sql.y: 3563
		SeeAlso: `SHOW TABLES, CREATE VIEW, SHOW CREATE,
WEBDOCS/create-table.html
WEBDOCS/create-table-as.html
`,
	},
	//line sql.y: 4076
	`CREATE SEQUENCE`: {
		ShortDescription: `create a new sequence`,
		//line sql.y: 4077
		Category: hDDL,
		//line sql.y: 4078
		Text: `
CREATE SEQUENCE <seqname>
  [INCREMENT <increment>]
  [MINVALUE <minvalue> | NO MINVALUE]
  [MAXVALUE <maxvalue> | NO MAXVALUE]
  [START [WITH] <start>]
  [CACHE <cache>]
  [NO CYCLE]
  [VIRTUAL]

`,
		//line sql.y: 4088
		SeeAlso: `CREATE TABLE
`,
	},
	//line sql.y: 4142
	`TRUNCATE`: {
		ShortDescription: `empty one or more tables`,
		//line sql.y: 4143
		Category: hDML,
		//line sql.y: 4144
		Text: `TRUNCATE [TABLE] <tablename> [, ...] [CASCADE | RESTRICT]
`,
		//line sql.y: 4145
		SeeAlso: `WEBDOCS/truncate.html
`,
	},
	//line sql.y: 4153
	`CREATE USER`: {
		ShortDescription: `define a new user`,
		//line sql.y: 4154
		Category: hPriv,
		//line sql.y: 4155
		Text: `CREATE USER [IF NOT EXISTS] <name> [ [WITH] PASSWORD <passwd> ]
`,
		//line sql.y: 4156
		SeeAlso: `DROP USER, SHOW USERS, WEBDOCS/create-user.html
`,
	},
	//line sql.y: 4178
	`CREATE ROLE`: {
		ShortDescription: `define a new role`,
		//line sql.y: 4179
		Category: hPriv,
		//line sql.y: 4180
		Text: `CREATE ROLE [IF NOT EXISTS] <name>
`,
		//line sql.y: 4181
		SeeAlso: `DROP ROLE, SHOW ROLES
`,
	},
	//line sql.y: 4193
	`CREATE VIEW`: {
		ShortDescription: `create a new view`,
		//line sql.y: 4194
		Category: hDDL,
		//line sql.y: 4195
		Text: `CREATE VIEW <viewname> [( <colnames...> )] AS <source>
`,
		//line sql.y: 4196
		SeeAlso: `CREATE TABLE, SHOW CREATE, WEBDOCS/create-view.html
`,
	},
	//line sql.y: 4226
	`CREATE INDEX`: {
		ShortDescription: `create a new index`,
		//line sql.y: 4227
		Category: hDDL,
		//line sql.y: 4228
		Text: `
CREATE [UNIQUE | INVERTED] INDEX [IF NOT EXISTS] [<idxname>]
       ON <tablename> ( <colname> [ASC | DESC] [, ...] )
       [STORING ( <colnames...> )] [<interleave>]

Interleave clause:
   INTERLEAVE IN PARENT <tablename> ( <colnames...> ) [CASCADE | RESTRICT]

`,
		//line sql.y: 4236
		SeeAlso: `CREATE TABLE, SHOW INDEXES, SHOW CREATE,
WEBDOCS/create-index.html
`,
	},
	//line sql.y: 4434
	`RELEASE`: {
		ShortDescription: `complete a retryable block`,
		//line sql.y: 4435
		Category: hTxn,
		//line sql.y: 4436
		Text: `RELEASE [SAVEPOINT] cockroach_restart
`,
		//line sql.y: 4437
		SeeAlso: `SAVEPOINT, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 4445
	`RESUME JOBS`: {
		ShortDescription: `resume background jobs`,
		//line sql.y: 4446
		Category: hMisc,
		//line sql.y: 4447
		Text: `
RESUME JOBS <selectclause>
RESUME JOB <jobid>
`,
		//line sql.y: 4450
		SeeAlso: `SHOW JOBS, CANCEL JOBS, PAUSE JOBS
`,
	},
	//line sql.y: 4467
	`SAVEPOINT`: {
		ShortDescription: `start a retryable block`,
		//line sql.y: 4468
		Category: hTxn,
		//line sql.y: 4469
		Text: `SAVEPOINT cockroach_restart
`,
		//line sql.y: 4470
		SeeAlso: `RELEASE, WEBDOCS/savepoint.html
`,
	},
	//line sql.y: 4485
	`BEGIN`: {
		ShortDescription: `start a transaction`,
		//line sql.y: 4486
		Category: hTxn,
		//line sql.y: 4487
		Text: `
BEGIN [TRANSACTION] [ <txnparameter> [[,] ...] ]
START TRANSACTION [ <txnparameter> [[,] ...] ]

Transaction parameters:
   ISOLATION LEVEL { SNAPSHOT | SERIALIZABLE }
   PRIORITY { LOW | NORMAL | HIGH }

`,
		//line sql.y: 4495
		SeeAlso: `COMMIT, ROLLBACK, WEBDOCS/begin-transaction.html
`,
	},
	//line sql.y: 4508
	`COMMIT`: {
		ShortDescription: `commit the current transaction`,
		//line sql.y: 4509
		Category: hTxn,
		//line sql.y: 4510
		Text: `
COMMIT [TRANSACTION]
END [TRANSACTION]
`,
		//line sql.y: 4513
		SeeAlso: `BEGIN, ROLLBACK, WEBDOCS/commit-transaction.html
`,
	},
	//line sql.y: 4537
	`ROLLBACK`: {
		ShortDescription: `abort the current transaction`,
		//line sql.y: 4538
		Category: hTxn,
		//line sql.y: 4539
		Text: `ROLLBACK [TRANSACTION] [TO [SAVEPOINT] cockroach_restart]
`,
		//line sql.y: 4540
		SeeAlso: `BEGIN, COMMIT, SAVEPOINT, WEBDOCS/rollback-transaction.html
`,
	},
	//line sql.y: 4653
	`CREATE DATABASE`: {
		ShortDescription: `create a new database`,
		//line sql.y: 4654
		Category: hDDL,
		//line sql.y: 4655
		Text: `CREATE DATABASE [IF NOT EXISTS] <name>
`,
		//line sql.y: 4656
		SeeAlso: `WEBDOCS/create-database.html
`,
	},
	//line sql.y: 4725
	`INSERT`: {
		ShortDescription: `create new rows in a table`,
		//line sql.y: 4726
		Category: hDML,
		//line sql.y: 4727
		Text: `
INSERT INTO <tablename> [[AS] <name>] [( <colnames...> )]
       <selectclause>
       [ON CONFLICT [( <colnames...> )] {DO UPDATE SET ... [WHERE <expr>] | DO NOTHING}]
       [RETURNING <exprs...>]
`,
		//line sql.y: 4732
		SeeAlso: `UPSERT, UPDATE, DELETE, WEBDOCS/insert.html
`,
	},
	//line sql.y: 4751
	`UPSERT`: {
		ShortDescription: `create or replace rows in a table`,
		//line sql.y: 4752
		Category: hDML,
		//line sql.y: 4753
		Text: `
UPSERT INTO <tablename> [AS <name>] [( <colnames...> )]
       <selectclause>
       [RETURNING <exprs...>]
`,
		//line sql.y: 4757
		SeeAlso: `INSERT, UPDATE, DELETE, WEBDOCS/upsert.html
`,
	},
	//line sql.y: 4862
	`UPDATE`: {
		ShortDescription: `update rows of a table`,
		//line sql.y: 4863
		Category: hDML,
		//line sql.y: 4864
		Text: `
UPDATE <tablename> [[AS] <name>]
       SET ...
       [WHERE <expr>]
       [ORDER BY <exprs...>]
       [LIMIT <expr>]
       [RETURNING <exprs...>]
`,
		//line sql.y: 4871
		SeeAlso: `INSERT, UPSERT, DELETE, WEBDOCS/update.html
`,
	},
	//line sql.y: 5041
	`<SELECTCLAUSE>`: {
		ShortDescription: `access tabular data`,
		//line sql.y: 5042
		Category: hDML,
		//line sql.y: 5043
		Text: `
Select clause:
  TABLE <tablename>
  VALUES ( <exprs...> ) [ , ... ]
  SELECT ... [ { INTERSECT | UNION | EXCEPT } [ ALL | DISTINCT ] <selectclause> ]
`,
	},
	//line sql.y: 5054
	`SELECT`: {
		ShortDescription: `retrieve rows from a data source and compute a result`,
		//line sql.y: 5055
		Category: hDML,
		//line sql.y: 5056
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
		//line sql.y: 5068
		SeeAlso: `WEBDOCS/select-clause.html
`,
	},
	//line sql.y: 5143
	`TABLE`: {
		ShortDescription: `select an entire table`,
		//line sql.y: 5144
		Category: hDML,
		//line sql.y: 5145
		Text: `TABLE <tablename>
`,
		//line sql.y: 5146
		SeeAlso: `SELECT, VALUES, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 5412
	`VALUES`: {
		ShortDescription: `select a given set of values`,
		//line sql.y: 5413
		Category: hDML,
		//line sql.y: 5414
		Text: `VALUES ( <exprs...> ) [, ...]
`,
		//line sql.y: 5415
		SeeAlso: `SELECT, TABLE, WEBDOCS/table-expressions.html
`,
	},
	//line sql.y: 5505
	`<SOURCE>`: {
		ShortDescription: `define a data source for SELECT`,
		//line sql.y: 5506
		Category: hDML,
		//line sql.y: 5507
		Text: `
Data sources:
  <tablename> [ @ { <idxname> | <indexhint> } ]
  <tablefunc> ( <exprs...> )
  ( { <selectclause> | <source> } )
  <source> [AS] <alias> [( <colnames...> )]
  <source> { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source> ON <expr>
  <source> { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source> USING ( <colnames...> )
  <source> NATURAL { [INNER] | { LEFT | RIGHT | FULL } [OUTER] } JOIN <source>
  <source> CROSS JOIN <source>
  <source> WITH ORDINALITY
  '[' EXPLAIN ... ']'
  '[' SHOW ... ']'

Index flags:
  '{' FORCE_INDEX = <idxname> [, ...] '}'
  '{' NO_INDEX_JOIN [, ...] '}'

`,
		//line sql.y: 5525
		SeeAlso: `WEBDOCS/table-expressions.html
`,
	},
}
