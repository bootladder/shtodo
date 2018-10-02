@test "Given a todo.txt exists, Given a /etc/shtodo.conf exists, Given shtodo.conf specifies todopath=/opt/other_path_to_todo.txt , Then, When running shtodo, It prints Todo" {
  echo blah > /opt/other_path_to_todo.txt
  echo "todopath: /opt/other_path_to_todo.txt" > /etc/shtodo.conf
  #cp /opt/tests/test_shtodo.conf /etc/shtodo.conf
  echo todo is: $(cat /opt/other_path_to_todo.txt)
  echo /etc/shtodo.conf is: $(cat /etc/shtodo.conf)
  run /opt/shtodo
  [ "$output" = "blah" ]
}
