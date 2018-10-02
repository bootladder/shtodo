@test "Given a todo.txt exists, Given a /etc/shtodo.conf exists, Given shtodo.conf specifies todopath=/opt/todo.txt , Then, When running shtodo, It prints Todo" {
  echo blah > /opt/todo.txt
  cp /opt/tests/test_shtodo.conf /etc/shtodo.conf
  run /opt/shtodo
  echo "/opt/todo.txt =" $(cat /opt/todo.txt)
  echo "status = ${status}"
  echo "output = ${output}"
  [ "$output" = "blah" ]
}
