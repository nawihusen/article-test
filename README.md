# Service Dapil Caleg

## How to

### Checkout

```bash
git clone https://mnc-repo.mncdigital.com/perindo/saksi/service-dapil-caleg
make init
```

### Run from Source

```bash
make run
```

### Build

```bash
make build
```

### Run

To run Service Dapil Caleg with default configuration use

```bash
./be-service-dapil-caleg
```

To run Service Dapil Caleg with configuration use

```bash
./be-service-dapil-caleg -c config.yaml
```

To run Service Dapil Caleg with environment variable

```bash
SERVER_PORT=8080 ./be-service-dapil-caleg
```

### Config and Environment Variable

You can run use YAML config file or environment variable. Here are the parameters.

| Config File | Environment Variable | Type | Default Value | Description |
|-------------|----------------------|------|---------------|-------------|
| server.port | SERVER_PORT | String | 8555 | Local machine TCP Port to bind the HTTP Server to |
| server.prefork | SERVER_PREFORK | Boolean | false | Prefork will spawn multiple Go processes listening on the same port |
| server.strict_routing | SERVER_STRICT_ROUTING | Boolean | false | When enabled, the router treats /foo and /foo/ as different |
| server.case_sensitive | SERVER_CASE_SENSITIVE | Boolean | false | When enabled, /Foo and /foo are different routes |
| server.body_limit | SERVER_BODY_LIMIT | Integer | 4194304 | Sets the maximum allowed size for a request body |
| server.concurrency | SERVER_CONCURRENCY | Integer | 262144 | Concurrency maximum number of concurrent connections |
| server.timeout.read | SERVER_TIMEOUT_READ | Integer | 5 | The amount of time to wait until an HTTP server read operation is cancelled |
| server.timeout.write | SERVER_TIMEOUT_WRITE | Integer | 10 | The amount of time to wait until an HTTP server write operation is cancelled |
| server.timeout.idle | SERVER_TIMEOUT_IDLE | Integer | 120 | The amount of time to wait until an IDLE HTTP session is closed |
| server.log_level | SERVER_LOG_LEVEL | String | debug | Log level, available value: `error`, `warning`, `info`, `debug` |
| redis.host | REDIS_HOST | String | localhost | The Redis IP Address to connect to |
| redis.port | REDIS_PORT | String | 6379 | The Redis Port to connect to |
| redis.max_connection | REDIS_MAX_CONNECTION | Integer | 80 | Redis maximum connection |
| redis.username | REDIS_USERNAME | String | | Redis username |
| redis.password | REDIS_PASSWORD | String | | Redis password |
| redis.database | REDIS_DATABASE | Integer | 0 | Redis database number |
| middleware.allows_origin | MIDDLEWARE_ALLOWS_ORIGIN | String | * | List of origins that allow for CORS |

### Sequence Diagram


### Entity Relationship Diagram (ERD)

```erd
[level] {bgcolor: "#e0e0e0"}
  *id {label: "int(11), auto_increment"}
  name {label: "varchar(20), not null"}

[partai] {bgcolor: "#f0b0b0"}
  *id {label: "int(11), auto_increment"}
  name {label: "varchar(50), not null"}
  full_name {label: "varchar(255), not null"}

[dapil] {bgcolor: "#f0d0b0"}
  *id {label: "int(11), auto_increment"}
  +level_id {label: "int(11), not null"}
  name {label: "varchar(255), not null"}

[dapil_map] {bgcolor: "#b0d0f0"}
  *id {label: "int(11), auto_increment"}
  +dapil_id {label: "int(11), not null"}
  *prov_code {label: "int(11), not null"}
  *kab_code {label: "int(11)"}
  *kec_code {label: "int(11)"}
  *kel_code {label: "int(11)"}

[caleg] {bgcolor: "#d0f0d0"}
  *id {label: "bigint(20), auto_increment"}
  +partai_id {label: "int(11), not null"}
  +dapil_id {label: "int(11), not null"}
  *seq_no {label: "int(11), not null"}
  name {label: "text, not null"}

dapil *--1 level {label: "level_id <--- id"}
dapil_map *--1 dapil {label: "dapil_id <--- id"}
caleg *--1 partai {label: "partai_id <--- id"}
caleg *--1 dapil {label: "dapil_id <--- id"}
```

*Example*

**level**
| *id | *name   |
|-----|---------|
|   1 | DPR RI  |
|   2 | DPRD I  |
|   3 | DPRD II |

**partai**
| *id | *name | *full_name |
|-----|-------|------------|
|   1 | PKB | Partai Kebangkitan Bangsa |
|   2 | Gerindra | Partai Gerakan Indonesia Raya |
|   3 | PDI Perjuangan | Partai Demokrasi Indonesia Perjuangan |
|   4 | Golkar | Partai Golongan Karya |
|   5 | NasDem | Partai NasDem |
|   6 | Partai Buruh | Partai Buruh |
|   7 | Gelora | Partai Gelombang Rakyat Indonesia |
|   8 | PKS | Partai Keadilan Sejahtera |
|   9 | PKN | Partai Kebangkitan Nusantara |
|  10 | Hanura | Partai Hati Nurani Rakyat |
|  11 | Garuda | Partai Garda Perubahan Indonesia |
|  12 | PAN | Partai Amanat Nasional |
|  13 | PBB | Partai Bulan Bintang |
|  14 | Partai Demokrat | Partai Demokrat |
|  15 | PSI | Partai Solidaritas Indonesia |
|  16 | Perindo | Partai Persatuan Indonesia |
|  17 | PPP | Partai Persatuan Pembangunan |

**dapil**
| *id | *level_id | *name |
|-----|-----------|-------|
|   1 |         1 | ACEH I |
|   2 |         1 | ACEH II |
|   3 |         1 | DKI JAKARTA I |
|   4 |         1 | DKI JAKARTA II |
|   5 |         2 | DP DKI JAKARTA 2 |
|   6 |         2 | DP DKI JAKARTA 3 |
|   7 |         2 | DP BALI 3 |
|   8 |         2 | DP SULAWESI BARAT 2 |
|   9 |         3 | DP SULAWESI BARAT 2 |
|  10 |         1 | BENGKULU |
|  11 |         1 | JAMBI |
|  12 |         3 | DP LABUHANBATU SELATAN 3 |
|  13 |         3 | DP LABUHANBATU SELATAN 4 |

**dapil_map**
| *id | *dapil_id | *prov_code | kab_code | kec_code | kel_code |
|-----|-----------|------------|----------|----------|----------|
|   1 |         1 |         11 |        1 |          |          |
|   2 |         1 |         11 |       71 |          |          |
|   3 |         1 |         11 |       75 |          |          |
|   4 |         2 |         11 |        3 |          |          |
|   5 |         2 |         11 |        4 |          |          |
|   6 |         2 |         11 |        8 |          |          |
|   7 |         3 |         31 |       75 |          |          |
|   8 |         4 |         31 |       71 |          |          |
|   9 |         4 |         31 |       74 |          |          |
|  10 |         5 |         31 |       72 |        3 |          |
|  11 |         6 |         31 |       72 |        1 |          |
|  12 |         7 |         51 |        2 |          |          |
|  13 |         8 |         76 |        4 |        4 |          |
|  14 |         9 |         76 |        4 |        4 |          |
|  15 |        10 |         17 |          |          |          |
|  16 |        11 |         15 |          |          |          |
|  17 |        12 |         12 |       22 |        3 |     2004 |
|  18 |        12 |         12 |       22 |        3 |     2005 |
|  19 |        12 |         13 |       22 |        3 |     2002 |


**caleg**
| *id | *partai_id | *dapil_id | *seq_no | *name |
|-----|------------|-----------|---------|-------|
|   1 |         16 |         4 |       1 | SURYADI, SE, MM |
|   2 |         16 |         4 |       2 | DAVID SETIAWAN, SPD |
|   3 |         16 |         2 |       1 | RICO SANGGEL, ST, MT |
|   4 |         10 |         1 |       1 | AMANDA WULAN, SE |
|   5 |         16 |         6 |       1 | TYO ADITIYA, SKOM, MM |

Sample query to get caleg names at DPR RI level from Perindo

```sql
SELECT level.name AS level, dapil.name AS dapil, partai.name AS partai, caleg.seq_no, caleg.name FROM caleg INNER JOIN dapil ON caleg.dapil_id = dapil.id INNER JOIN dapil_map ON dapil.id = dapil_map.dapil_id INNER JOIN partai ON caleg.partai_id = partai.id INNER JOIN level ON dapil.level_id = level.id WHERE dapil.level_id = 1 AND caleg.partai_id = 16 AND dapil_map.prov_code = 31 AND (dapil_map.kab_code = 74 OR dapil_map.kab_code IS NULL) AND (dapil_map.kec_code = 5 OR dapil_map.kec_code IS NULL) AND (dapil_map.kel_code = 7 OR dapil_map.kel_code IS NULL);
```

Result

```
+--------+----------------+---------+--------+---------------------+
| level  | dapil          | partai  | seq_no | name                |
+--------+----------------+---------+--------+---------------------+
| DPR RI | DKI JAKARTA II | Perindo |      1 | SURYADI, SE, MM     |
| DPR RI | DKI JAKARTA II | Perindo |      2 | DAVID SETIAWAN, SPD |
+--------+----------------+---------+--------+---------------------+
```
