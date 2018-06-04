#! /bin/bash
count=1

while read line; do
  if [ ${count} -eq 1 ]; then
    FROM=$line
  fi
  if [ ${count} -eq 2 ]; then
    TO=$line
  fi
  count=$(($count + 1))
done < address.txt

./Backbone Send -from $FROM -to $TO -amount 4 -mine
