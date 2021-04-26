### compare two time from clickhouse query result
#!/bin/bash
#fisrt time is millisecond， cut to seconds and format
cat $1 | awk -F'│' '{print $2}' | sed 's/ //g' | awk '{print substr($1,0,10)}' | awk '{print strftime("%Y-%m-%d %H:%M:%S",$1)}'> cbtime
#second time is format time
cat rd | awk -F'│' '{print $3}' | awk -F'"' '{print $ 52}' > msgtime
#compare two time
paste cbtime msgtime

rm cbtime
rm msgtime