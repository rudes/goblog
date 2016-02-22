<?php
include_once 'sqlconfig.php';
$db=new mysqli(HOST,USER,PASSWORD,DATABASE);
session_start();
$user_check=$_SESSION['login_user'];
$ses_sql=$db->query("select username from admin where username='$user_check'");
$row=$ses_sql->fetch_assoc();
$login_session=$row['username'];
if(!isset($login_session))
{
	header("Location: login.php");
}
?>
