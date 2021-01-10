pkgname: model
output: internal/adapter/infrastructure/model
add-global-variants: 'true'
add-panic-variants: 'true'
mysql:
  dbname: ${DB_NAME}
  host: ${DB_HOST}
  port: 3306
  user: ${DB_USER}
  pass: ${DB_PASS}
  sslmode: 'false'
  skipSQLCmd: 'true'
  blacklist:
    - gorp_migrations
