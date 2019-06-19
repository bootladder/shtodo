@test "Given a todo.txt exists, Given a /etc/shtodo.conf exists, Given shtodo.conf specifies todopath=/tmp/todo.txt , Then, When running shtodo, It prints Todo" {
  echo blah > /tmp/todo.txt
  cp /opt/tests/test_shtodo.conf /etc/shtodo.conf

  echo "status = ${status}"
  echo "output = ${output}"

  out=$(/opt/shtodo)
  echo output is $out
  echo $out | grep blah 
}
