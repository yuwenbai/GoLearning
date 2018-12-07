echo init database UpdateServerDB begin!
echo ----------------------------
osql -E -i UpdateServerDB_Init.sql
echo ----------------------------
echo init database UpdateServerDB ok!
pause