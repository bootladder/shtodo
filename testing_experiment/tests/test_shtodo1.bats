@test "Given a todo.txt exists, Given a /etc/shtodo.conf exists, Given shtodo.conf specifies todopath=/tmp/other_path_to_todo.txt , Then, When running shtodo, It prints Todo" {
  echo blah > /tmp/other_path_to_todo.txt
  echo "todopath: /tmp/other_path_to_todo.txt" > /etc/shtodo.conf
  echo todo is: $(cat /tmp/other_path_to_todo.txt)
  echo /etc/shtodo.conf is: $(cat /etc/shtodo.conf)

  out=$(/opt/shtodo)
  echo output is $out
  echo $out | grep blah 
}
