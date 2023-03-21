# College-Management-System
This is simple college record system created with golang and mysql

Read further to know how this project works:
->firstly establish the connecetion with database 
->call createtable() function during first execution only and this fuction shouldn't be called as table named in our program is already created in database. If this function is called after execution, there will be error as our program tries to reate a table with same name which is already in the database.
->the admin and the student have separate login with different functionalities

FOR ADMIN 
->for this program the username is provided as "admin" and password as "hello" in source code itself
->admin can perform different activities such as adding, viewing, updating,removing student detail
->while adding record first letter of First name and Last name should be in Capital Letter and every other value in small letter along with Composite attributes Without whitespace
->date of birth should be provided in "ddmmyy" format without any whitespace and other symbol (i.e. 23112002)
->admin can view all student record at once or specific student record (one or many student with same value) only by providing first name or last name or enrolled year or faculty enrolled
->can update student detail by providing the sid of student and updation can be done for specific value only or for all values in the table
->can remove student record from table by entering his/her sid

FOR STUDENT
->firstly student have to enter the sid which value should be in table. For entering the sid value for 3 times user is banned for 30 seconds from logging in
->if entered sid is found in the database user should enter the "email" as username and "dob" as password that match the enetered sid value. If not then user will be banned foe 60 second for entering wrong username and password that doesnt match entered sid value
Username: email
Password: dob (ddmmyy format)
