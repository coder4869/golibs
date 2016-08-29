```
<1> mkdir **: permission denied
Solution 1: 
	step 1: go to target directory in terminal
	step 2: run "$ chmod -R 777 ." which means grant permission to current directory
Solution 2: 
	change second parameter of os.MkdirAll() (called in gllog.NewLog()) as "07**",
	like 0777 or 0766 or 0764 or 0744 et al

More points:
	unix permission bits can be a number of 4 characters, like 0777.
	each character makes up by 4(r-read), 2(w-write) and 1(x-excute). for example 7=4+2+1.
	0777 equals -rwxrwxrwx, meaning as following permissions to current file(or dir):
	first 7 means current user has read(4)+write(2)+excute(1) permissions,
	second 7 means user's group has read(4)+write(2)+excute(1) permissions,
	last 7 means other user has read(4)+write(2)+excute(1) permissions.
```

