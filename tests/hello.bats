@test "Given a todo.txt exists, Given a /etc/shtodo.conf exists, Given shtodo.conf specifies todopath=/tmp/todo.txt , Then, When running shtodo, It prints Todo" {
  echo blah > /tmp/todo.txt
  cp /tmp/tests/test_shtodo.conf /etc/shtodo.conf
  run /tmp/shtodo
  result="$(echo 4)"
  [ "$output" = "blah" ]
}
