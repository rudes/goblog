<?php

include_once 'sqlconfig.php';
$db=new mysqli(HOST,USER,PASSWORD,DATABASE);
session_start();
if($_SERVER["REQUEST_METHOD"]=="POST") {
$user=mysqli_real_escape_string($db,$_POST['username']);
$pass=mysqli_real_escape_string($db,$_POST['password']);
$sql="SELECT id FROM admin WHERE username='$user' and passcode='$pass'";
$result=$db->query($sql);
if ($result->num_rows == 1) {
$_SESSION['login_user']=$user;
header('Location: admin.php');
} else {
	echo "Login Name or Password Incorrect.";
}
}
?>
<head>
<title>Admin Login | Other Letters</title>
<meta name="viewport" content="width=device-width, initial-scale=1">
<script type="text/javascript" src="//code.jquery.com/jquery-2.1.0.js"></script>
<link rel="stylesheet" type="text/css" href="styles.css">
</head>
<body>
<?php include_once 'header.php'?>
<div class="container">
<div class="content">
<form action="<?php echo $_SERVER['PHP_SELF'] ?>" method="post">
<label>Username: </label> <input type="text" name="username"/><br>
<br>
<label>Password: </label> <input type="password" name="password"/><br>
<br>
<input type="submit" value="Submit"/><br>
</form>
</div></div>
</body>
