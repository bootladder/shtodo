# shtodo

Instructions:

- `go install`
- make a /etc/shtodo.conf like tests/test_shtodo.conf
- Specify where your todo file is.  Make it a git repo
- add shtodo to your bashrc

**What does it do:**  

Print out your todo.txt when you open up a new shell.  
But not every time, only after a bit of time passes.  
Pushes and Pulls to github after more time passes  
Assumes you have SSH keys setup so the pushing and pulling can happen.

shtodo e  : edit your todo.txt  
shtodo cat : print your todo.txt  
shtodo pull  
shtodo push  

**For Reference:**
```
todopath: /tmp/mytodo/todo.txt
todointerval: 30
pushinterval: 300
pullinterval: 300
```
