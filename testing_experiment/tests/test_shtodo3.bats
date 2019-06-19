@test "Given a todo.txt does not exist, Given a /etc/shtodo.conf exists, Given shtodo.conf specifies todopath=/tmp/badpath.txt , Then, When running shtodo, It Fails" {
  echo "todopath: /tmp/badpath.txt" > /etc/shtodo.conf
  echo /etc/shtodo.conf is: $(cat /etc/shtodo.conf)
  run /opt/shtodo
  [ "$status" != "0" ]
}

