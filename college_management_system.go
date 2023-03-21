package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var dbm *sql.DB

// connecting to database
func connectDB() {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/student_info?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	db.Ping()
	fmt.Println("My sql connected....")
	dbm = db
}

// creation of tables for our databse studentInfo
func createTable() {
	query := `Create table student(
		sid int auto_increment,
		fname text not null,
		lname text not null,
		dob int not null,
		enrolled_Year int not null,
		faculty text not null,
		email text not null,
		phNo bigint,
		created_at datetime,
		primary key(sid)
		);`
	_, err := dbm.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Student table created....")

}
func clearScreen() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func firstScreen() {
D:
	var inp string
	fmt.Printf("\n\n\n")
	fmt.Println("Press c to continue: ")
	fmt.Scanln(&inp)
	switch {
	case inp == "c":
		clearScreen()
		secondScreen()
	default:
		fmt.Println("Invalid Input")
		time.Sleep(3 * time.Second)
		goto D
	}

}

func secondScreen() {
A:
	clearScreen()
	fmt.Println("Press 1: Admin Sector Login: ")
	fmt.Println("Press 2: Student Sector Login: ")
	fmt.Println("Press 3: Exit: ")
	fmt.Printf("\n\n\n")

	var inp1 int
	fmt.Println("Enter your choices:")
	fmt.Scanln(&inp1)
	ch := inp1
	switch ch {
	case 1:
		clearScreen()
		adminLogin()
	case 2:
		clearScreen()
		studentLogin()
	case 3:
		clearScreen()
		os.Exit(5)
	default:
		fmt.Println("Enter the valid Input")
		goto A
	}
}
func adminLogin() {
B:
	clearScreen()
	fmt.Println("Admin Login Portal")
	fmt.Printf("\n\n\n")

	fmt.Print("Username:")
	var adUsername string
	fmt.Scanln(&adUsername)
	fmt.Print("Password:")
	var adPassword string
	fmt.Scanln(&adPassword)

	if adUsername == "admin" && adPassword == "hello" {
		clearScreen()
		adminInterFace()
	} else {
		fmt.Println("Invalid Username or Password")
		time.Sleep(3 * time.Second)
		goto B
	}

}

func adminInterFace() {
invalidInput:
	fmt.Println("Press 1: Add Student Record")
	fmt.Println("Press 2: See Student Details")
	fmt.Println("Press 3: Update Details of Student ")
	fmt.Println("Press 4: Remove Student Record")
	fmt.Println("press 5: Send Notice for Student")
	fmt.Println("Press 6: Logout")
	var inp2 int
	fmt.Println("Enter your choice:")
	fmt.Scanln(&inp2)
	switch inp2 {
	case 1:
		clearScreen()
		addRecord()
	case 2:
		clearScreen()
		seeStudentDetail()
	case 3:
		clearScreen()
		updateStudentDetail()
	case 4:
		clearScreen()
		removeRecord()
	case 5:
		clearScreen()
		adminLogout()
	default:
		fmt.Println("--Invalid Input--\n")
		time.Sleep(3 * time.Second)
		clearScreen()
		fmt.Println("Enter the valid Input\n")
		goto invalidInput
	}
}

// adding student detail to database
func addRecord() {
	fmt.Println("Welcome to Admin Add Record Interface")
	//var sid int
	var dob, enrolled_Year int
	var phNo int64
	var fname, lname, faculty, email string
	fmt.Println("\n\nNote: Enter the first letter of First name and Last name in Capital Letter and every other value in small letter along with Composite attributes Without Space :-\n\n")
	//fmt.Println("Enter sid")
	//fmt.Scanln(&sid)
	fmt.Println("Enter First Name")
	fmt.Scanln(&fname)
	fmt.Println("Enter Last Name")
	fmt.Scanln(&lname)
	fmt.Println("Enter Date of Birth(ddmmyy)")
	fmt.Scanln(&dob)
	fmt.Println("Enter Enrolled Year")
	fmt.Scanln(&enrolled_Year)
	fmt.Println("Enter Faculty Enrolled")
	fmt.Scanln(&faculty)
	fmt.Println("Enter Email Address")
	fmt.Scanln(&email)
	fmt.Println("Enter Phone Number")
	fmt.Scanln(&phNo)
	created_at := time.Now()
	result, err := dbm.Exec(`insert into student (fname,lname,dob,enrolled_Year,faculty,email,phNo,created_at)
	    values(?,?,?,?,?,?,?,?)`, fname, lname, dob, enrolled_Year, faculty, email, phNo, created_at)
	if err != nil {
		log.Fatal(err)
	} else {
		value, _ := result.LastInsertId()
		fmt.Println("Record Inserted for Student ID=", value)
	}
	fmt.Println("\n\n--Wait for Some Time--")
	time.Sleep(3 * time.Second)
	clearScreen()
	adminInterFace()
}

// viewing student detail from admin interface
func seeStudentDetail() {
	//structure to hold values from database
	type student struct {
		sid           int
		fname         string
		lname         string
		dob           int
		enrolled_Year int
		faculty       string
		email         string
		phNo          int64
		create_at     time.Time
	}
	var s student
	fmt.Println("Welcome to Admin Student Detail Interface")
again:
	fmt.Println("\n")
	fmt.Println("Enter 1: See All Student Detail")
	fmt.Println("Enter 2: See Specific Student Detail Only")
	fmt.Println("Enter 3: Back")
	var a int
	fmt.Println("Enter Your Choice")
	fmt.Scanln(&a)
	clearScreen()
	switch a {
	case 1:
		rows, err := dbm.Query(`select * from student`)
		if err != nil {
			log.Fatal(err)
		} else {
			/*
				fmt.Println("ID\t\tFName\t\tLName\t\tDOB\t\tEYear\t\tfaculty\t\tCreated At\n\n")
				for rows.Next() {
					rows.Scan(&s.sid, &s.fname, &s.lname, &s.dob, &s.enrolled_Year, &s.faculty, &s.create_at)
					fmt.Println(s.sid, "\t", s.fname, "\t", s.lname, "\t", s.dob, "\t", s.enrolled_Year, "\t", s.faculty, "\t", s.create_at)
			*/
			for rows.Next() {
				rows.Scan(&s.sid, &s.fname, &s.lname, &s.dob, &s.enrolled_Year, &s.faculty, &s.email, &s.phNo, &s.create_at)
				fmt.Println("Student ID     :", s.sid)
				fmt.Println("Name           :", s.fname, s.lname)
				fmt.Println("DOB            :", s.dob)
				fmt.Println("Enrolled Year  :", s.enrolled_Year)
				fmt.Println("Faculty        :", s.faculty)
				fmt.Println("Email          :", s.email)
				fmt.Println("Phone No.      :", s.phNo)
				fmt.Println("Created Date   :", s.create_at)
				fmt.Println("\n\n")
			}
			time.Sleep(5 * time.Second)
			goto again
		}
	case 2:
		clearScreen()
	specificAgain:
		fmt.Println("--Specific Student Detail Only--\n")
		fmt.Println("1: Detail By First Name")
		fmt.Println("2: Detail By Last Name")
		fmt.Println("3: Detail By Date of Birth")
		fmt.Println("4: Detail By Enrolled Year")
		fmt.Println("5: Detail By Faculty Enrolled")
		var a int
		fmt.Println("Enter your choice:-")
		fmt.Scanln(&a)
		switch a {
		//detail by fname
		case 1:
			var name1 string
			fmt.Println("\nEnter the First Name you want to search for:-")
			fmt.Scanln(&name1)
			fmt.Printf("\n")
			rows, err := dbm.Query(`select * from student where fname=?`, name1)
			if err != nil {
				log.Fatal(err)
			} else {
				for rows.Next() {
					rows.Scan(&s.sid, &s.fname, &s.lname, &s.dob, &s.enrolled_Year, &s.faculty, &s.email, &s.phNo, &s.create_at)
					fmt.Println("Student ID     :", s.sid)
					fmt.Println("Name           :", s.fname, s.lname)
					fmt.Println("DOB            :", s.dob)
					fmt.Println("Enrolled Year  :", s.enrolled_Year)
					fmt.Println("Faculty        :", s.faculty)
					fmt.Println("Email          :", s.email)
					fmt.Println("Phone No.      :", s.phNo)
					fmt.Println("Created Date   :", s.create_at)
					fmt.Println("\n\n")
				}
				time.Sleep(5 * time.Second)
				goto again
			}
		//detail by lname
		case 2:
			var name2 string
			fmt.Println("\nEnter the Last Name you want to search for:-")
			fmt.Scanln(&name2)
			fmt.Printf("\n")
			rows, err := dbm.Query(`select * from student where lname=?`, name2)
			if err != nil {
				log.Fatal(err)
			} else {
				for rows.Next() {
					rows.Scan(&s.sid, &s.fname, &s.lname, &s.dob, &s.enrolled_Year, &s.faculty, &s.email, &s.phNo, &s.create_at)
					fmt.Println("Student ID     :", s.sid)
					fmt.Println("Name           :", s.fname, s.lname)
					fmt.Println("DOB            :", s.dob)
					fmt.Println("Enrolled Year  :", s.enrolled_Year)
					fmt.Println("Faculty        :", s.faculty)
					fmt.Println("Email          :", s.email)
					fmt.Println("Phone No.      :", s.phNo)
					fmt.Println("Created Date   :", s.create_at)
					fmt.Println("\n\n")
				}
				time.Sleep(5 * time.Second)
				goto again
			}
		//detail by date of birth
		case 3:
			var dob1 int
			fmt.Println("\nEnter the DOB(ddmmyy) you want to search for:-")
			fmt.Scanln(&dob1)
			fmt.Printf("\n")
			rows, err := dbm.Query(`select * from student where dob=?`, dob1)
			if err != nil {
				log.Fatal(err)
			} else {
				for rows.Next() {
					rows.Scan(&s.sid, &s.fname, &s.lname, &s.dob, &s.enrolled_Year, &s.faculty, &s.email, &s.phNo, &s.create_at)
					fmt.Println("Student ID     :", s.sid)
					fmt.Println("Name           :", s.fname, s.lname)
					fmt.Println("DOB            :", s.dob)
					fmt.Println("Enrolled Year  :", s.enrolled_Year)
					fmt.Println("Faculty        :", s.faculty)
					fmt.Println("Email          :", s.email)
					fmt.Println("Phone No.      :", s.phNo)
					fmt.Println("Created Date   :", s.create_at)
					fmt.Println("\n\n")
				}
				time.Sleep(5 * time.Second)
				goto again
			}
		//detail by enrolled year
		case 4:
			var enrolYear int
			fmt.Println("\nEnter the Year you want to search for:-")
			fmt.Printf("\n")
			fmt.Scanln(&enrolYear)
			rows, err := dbm.Query(`select * from student where enrolled_Year=?`, enrolYear)
			if err != nil {
				log.Fatal(err)
			} else {
				for rows.Next() {
					rows.Scan(&s.sid, &s.fname, &s.lname, &s.dob, &s.enrolled_Year, &s.faculty, &s.email, &s.phNo, &s.create_at)
					fmt.Println("Student ID     :", s.sid)
					fmt.Println("Name           :", s.fname, s.lname)
					fmt.Println("DOB            :", s.dob)
					fmt.Println("Enrolled Year  :", s.enrolled_Year)
					fmt.Println("Faculty        :", s.faculty)
					fmt.Println("Email          :", s.email)
					fmt.Println("Phone No.      :", s.phNo)
					fmt.Println("Created Date   :", s.create_at)
					fmt.Println("\n\n")
				}
				time.Sleep(5 * time.Second)
				goto again
			}
		//detail by faculty
		case 5:
			var faculty1 string
			fmt.Println("\nEnter the Faculty you want to search for:-")
			fmt.Scanln(&faculty1)
			fmt.Printf("\n")
			rows, err := dbm.Query(`select * from student where faculty=?`, faculty1)
			if err != nil {
				log.Fatal(err)
			} else {
				for rows.Next() {
					rows.Scan(&s.sid, &s.fname, &s.lname, &s.dob, &s.enrolled_Year, &s.faculty, &s.email, &s.phNo, &s.create_at)
					fmt.Println("Student ID     :", s.sid)
					fmt.Println("Name           :", s.fname, s.lname)
					fmt.Println("DOB            :", s.dob)
					fmt.Println("Enrolled Year  :", s.enrolled_Year)
					fmt.Println("Faculty        :", s.faculty)
					fmt.Println("Email          :", s.email)
					fmt.Println("Phone No.      :", s.phNo)
					fmt.Println("Created Date   :", s.create_at)
					fmt.Println("\n\n")
				}
				time.Sleep(5 * time.Second)
				goto again
			}
		default:
			fmt.Println("Invalid Choice\n")
			goto specificAgain

		} //closing second switch
	case 3:
	invalidInput:
		fmt.Printf("\nAre you sure you want to go back? (y/n):")
		var back string
		fmt.Scanln(&back)
		if back == "y" {
			clearScreen()
			adminInterFace()
		} else if back == "n" {
			clearScreen()
			goto again
		} else {
			fmt.Println("Invalid Input")
			goto invalidInput
		}
		adminInterFace()
	default:
		fmt.Println("Invalid Choice")
		goto again

	} //closing of first switch
}

// updating student detail by sid
func updateStudentDetail() {
	fmt.Println("Welcome to Admin Update Student Detail interface\n")
	var sid1, update1 int
seeDetailAgain:
	fmt.Println("Enter the sid of Student to Update his/her detail")
	fmt.Scanln(&sid1)
	fmt.Println("\n1: First Name")
	fmt.Println("2: Last Name")
	fmt.Println("3: Date Of Birth")
	fmt.Println("4: Enrolled Year")
	fmt.Println("5: Faculty")
	fmt.Println("6: Email Address")
	fmt.Println("7: Phone No")
	fmt.Println("8: All Detail")
	fmt.Println("9: None (Return Back)")
	fmt.Println("\nWhat you want to Update ?\n")
	fmt.Scanln(&update1)
	switch update1 {
	case 1:
		var name4 string
		fmt.Println("Enter the First Name to Update in  Table: ")
		fmt.Scanln(&name4)
		_, err := dbm.Exec(`update student set fname=? where sid=?`, name4, sid1)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("\n--Record Successfully Updated--\n")
			time.Sleep(3 * time.Second)
			clearScreen()
			adminInterFace()
		}

	case 2:
		var surName string
		fmt.Println("Enter the Surname to Update in Table: ")
		fmt.Scanln(&surName)
		_, err := dbm.Exec(`update student set lname=? where sid=?`, surName, sid1)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("\n--Record Successfully Updated--\n")
			time.Sleep(3 * time.Second)
			clearScreen()
			adminInterFace()
		}
	case 3:
		var dob2 int
		fmt.Println("Enter the Date of Birth(ddmmyy) to Update in Table: ")
		fmt.Scanln(&dob2)
		_, err := dbm.Exec(`update student set dob=? where sid=?`, dob2, sid1)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("\n--Record Successfully Updated--\n")
			time.Sleep(3 * time.Second)
			clearScreen()
			adminInterFace()
		}
	case 4:
		var enrolled_Year1 int
		fmt.Println("Enter the Year of Enrollment to Update in Table: ")
		fmt.Scanln(&enrolled_Year1)
		_, err := dbm.Exec(`update student set enrolled_Year=? where sid=?`, enrolled_Year1, sid1)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("\n--Record Successfully Updated--\n")
			time.Sleep(3 * time.Second)
			clearScreen()
			adminInterFace()
		}
	case 5:
		var faculty2 int
		fmt.Println("Enter the Faculty name to Update in Table: ")
		fmt.Scanln(&faculty2)
		_, err := dbm.Exec(`update student set faculty=? where sid=?`, faculty2, sid1)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("\n--Record Successfully Updated--\n")
			time.Sleep(3 * time.Second)
			clearScreen()
			adminInterFace()
		}
	case 6:
		var email1 int
		fmt.Println("Enter the Email Address to Update in Table: ")
		fmt.Scanln(&email1)
		_, err := dbm.Exec(`update student set email=? where sid=?`, email1, sid1)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("--Record Successfully Updated--\n")
			time.Sleep(3 * time.Second)
			clearScreen()
			adminInterFace()
		}
	case 7:
		var phNo1 int64
		fmt.Println("Enter the Phone Number to Update in Table: ")
		fmt.Scanln(&phNo1)
		_, err := dbm.Exec(`update student set phNo=? where sid=?`, phNo1, sid1)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("--Record Successfully Updated--\n")
			time.Sleep(3 * time.Second)
			clearScreen()
			adminInterFace()
		}
	case 8:
		var dob, enrolled_Year int
		var phNo int64
		var fname, lname, faculty, email string
		fmt.Println("\n\nNote: Enter the first letter of First name and Fast name in Capital Letter and every other value in small letter along with Composite attributes Without Space :-\n\n")
		fmt.Println("Enter First Name")
		fmt.Scanln(&fname)
		fmt.Println("Enter Last Name")
		fmt.Scanln(&lname)
		fmt.Println("Enter Date of Birth(ddmmyy)")
		fmt.Scanln(&dob)
		fmt.Println("Enter Enrolled Year")
		fmt.Scanln(&enrolled_Year)
		fmt.Println("Enter Faculty Enrolled")
		fmt.Scanln(&faculty)
		fmt.Println("Enter Email Address")
		fmt.Scanln(&email)
		fmt.Println("Enter Phone Number")
		fmt.Scanln(&phNo)
		_, err := dbm.Exec(`update student set fname=?,lname=?,dob=?,enrolled_Year=?,faculty=?,email=?,phNo=? where sid=?`,
			fname, lname, dob, enrolled_Year, faculty, email, phNo, sid1)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Println("\n--Record Successfully Updated--\n")
			time.Sleep(3 * time.Second)
			clearScreen()
			adminInterFace()
		}
	case 9:
	invalidInput:
		fmt.Printf("\nAre you sure you want to go Back? (y/n):")
		var adLg string
		fmt.Scanln(&adLg)
		if adLg == "y" {
			clearScreen()
			adminInterFace()
		} else if adLg == "n" {
			clearScreen()
			updateStudentDetail()
		} else {
			fmt.Println("--Invalid Input--")
			goto invalidInput
		}
	default:
		fmt.Println("\n--Invalid Input--")
		fmt.Println("Enter the choice between 1-9")
		goto seeDetailAgain

	}
}

// deleting student detail by sid
func removeRecord() {
	var stuSid int
	var adminChoice string
	fmt.Println("Welcome to admin remove record interface\n")
	fmt.Println("Enter the sid of the Student to Delete his/her record:   ")
	fmt.Scanln(&stuSid)
askAgain:
	fmt.Println("Are you sure want to Delete the record of student of sid ", stuSid, "? (y/n) :")
	fmt.Scanln(&adminChoice)
	if adminChoice == "y" {
		_, err := dbm.Exec(`delete from student where sid=?`, stuSid)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Record successfully Deleted for sid ", stuSid)
		time.Sleep(3 * time.Second)
		adminInterFace()
	} else if adminChoice == "n" {
		adminInterFace()
	} else {
		fmt.Println("--Invalid Input--")
		goto askAgain

	}
}
func adminLogout() {
	fmt.Println("Welcome to admin logout interface")
invalidInput:
	fmt.Printf("\nAre you sure you want to logout? (y/n):")
	var adLg string
	fmt.Scanln(&adLg)
	if adLg == "y" {
		clearScreen()
		secondScreen()
	} else if adLg == "n" {
		clearScreen()
		adminInterFace()
	} else {
		fmt.Println("Invalid Input")
		goto invalidInput
	}
}

func studentLogin() {
	clearScreen()
	counte := 0
A:
	if counte >= 3 {
		fmt.Printf("\n\n")
		fmt.Println("+--------------------------------------------------------------------------------+")
		fmt.Println("| You have been banned for 5 second from login because of multiple incorrect sid |")
		fmt.Println("+--------------------------------------------------------------------------------+")
		time.Sleep(5 * time.Second)
	C:
		fmt.Printf("Press 'e' to go back to menu! ")
		var ab string
		fmt.Scanln(&ab)
		if ab == "e" {
			clearScreen()
			secondScreen()
		} else {
			fmt.Println("Invalid Input!")
			goto C
		}

	}
	fmt.Println("Welcome to Student Login interface")
	fmt.Printf("\n\n")

	type student struct {
		sid           int
		fname         string
		lname         string
		dob           int
		enrolled_Year int
		faculty       string
		email         string
		phNo          int64
		create_at     time.Time
	}
	var s student
	var a int
	count := 0
	fmt.Printf("Enter the Student Id: ")
	fmt.Scanln(&a)
	rows, err := dbm.Query(`select * from student`)
	if err != nil {
		log.Fatal(err)
	} else {
		for rows.Next() {
			rows.Scan(&s.sid, &s.fname, &s.lname, &s.dob, &s.enrolled_Year, &s.faculty, &s.email, &s.phNo, &s.create_at)
			if a == s.sid {
			B:
				if count >= 3 {
					fmt.Printf("\n\n")
					fmt.Println("+------------------------------------------------------------------------------------------+")
					fmt.Println("| You have been banned for 5 second from login because of multiple incorrect login attempt |")
					fmt.Println("+------------------------------------------------------------------------------------------+")
					time.Sleep(5 * time.Second)
					clearScreen()
					secondScreen()
				}
				clearScreen()
				fmt.Println("Student Login Portal! ")
				fmt.Println("---------------------")
				fmt.Printf("\n\n")
				var user string
				fmt.Printf("Username: ")
				fmt.Scanln(&user)

				var pass int
				fmt.Printf("Password: ")
				fmt.Scanln(&pass)
				x := s.sid

				if user == s.email && pass == s.dob {
					fmt.Println("The Login credential is correct!!!!!")
					time.Sleep(1 * time.Second)
					clearScreen()
					studentInterface(x)
				} else {
					count += 1
					fmt.Println("Invalid Username or password")
					time.Sleep(1 * time.Second)
					goto B
				}
			}
		}
	}
	fmt.Println("Invalid SID.")
	counte += 1
	time.Sleep(1 * time.Second)
	clearScreen()
	goto A
}

func studentInterface(a int) {
A:
	type student struct {
		sid           int
		fname         string
		lname         string
		dob           int
		enrolled_Year int
		faculty       string
		email         string
		phNo          int64
		create_at     time.Time
	}
	var g student
	row, err := dbm.Query(`select * from student where sid = ?`, a)
	if err != nil {
		log.Fatal(err)
	} else {
		for row.Next() {
			row.Scan(&g.sid, &g.fname, &g.lname, &g.dob, &g.enrolled_Year, &g.faculty, &g.email, &g.phNo, &g.create_at)
			fmt.Printf("\n\n")
			fmt.Println("Welcome, Dear", g.fname)
			fmt.Println("---------------------------------")
			fmt.Println("\t\t\t\t\t\t User Id:", a)
			fmt.Println("\t\t\t\t\t\t ---------------------")
			fmt.Println("Your Profile Details:")
			fmt.Println("\t Name            : ", g.fname, g.lname)
			fmt.Println("\t DOB             : ", g.dob)
			fmt.Println("\t Enrolled Year   : ", g.enrolled_Year)
			fmt.Println("\t Faculty         : ", g.faculty)
			fmt.Println("\t Email           : ", g.email)
			fmt.Println("\t Phone No.       : ", g.phNo)
			fmt.Println("Created Date       : ", g.create_at)
			fmt.Printf("\n\n")
			fmt.Printf("\n\n")
			var inp string
			fmt.Printf("Press 'e' to logout! ")
			fmt.Scanln(&inp)
			if inp == "e" {
				fmt.Printf("Are you sure you want to Logout? (y/n): ")
				var inp1 string
			B:
				fmt.Scanln(&inp1)
				if inp1 == "y" {
					clearScreen()
					secondScreen()
				} else if inp1 == "n" {
					clearScreen()
					goto A
				} else {
					fmt.Println("Invalid Input!!!")
					goto B
				}
			} else {
				fmt.Println("Invalid Input!!!")
			}
		}
	}
}

func main() {
	fmt.Println("###############################")
	fmt.Println("#                             #")
	fmt.Println("#  XYZ COllEGE RECORD SYSTEM  #")
	fmt.Println("#                             #")
	fmt.Println("###############################")
	connectDB()
	//createTable() //no need to run this function multiple times as creation is completed at first execution
	firstScreen()
}
