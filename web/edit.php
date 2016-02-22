<?php
include('lock.php');
include_once 'sqlconfig.php';
?>
<!DOCTYPE html>
<html>
<head>
<title>Edit Panel | Other Letters</title>
<meta name="viewport" content="width=device-width, initial-scale=1">
<script type="text/javascript" src="//code.jquery.com/jquery-2.1.0.js"></script>
<link href="//netdna.bootstrapcdn.com/bootswatch/3.0.1/cosmo/bootstrap.min.css" rel="stylesheet"> 
<script src="//netdna.bootstrapcdn.com/bootstrap/3.0.1/js/bootstrap.min.js"></script> 
<link href="//netdna.bootstrapcdn.com/font-awesome/4.0.3/css/font-awesome.min.css" rel="stylesheet">
<link href="summernote/summernote.css" rel="stylesheet">
<script src="summernote/summernote.min.js"></script>
<script>
$(document).ready(function() {
				$('#summernote').summernote({
								height: 250,
				});
});
function postContent() {
				$('textarea[name="content"]').html($('#summernote').code());
}
</script>
<link rel="stylesheet" type="text/css" href="styles.css">
<?php include_once 'header.php';
function updatePost() {
	$id=mb_substr(md5($_POST['title']),0,10);
	$content=$_POST['content'];
	$db=new mysqli(HOST,USER,PASSWORD,DATABASE);
	if ($db->connect_error) {
		die("Connection Failed: " . $db->connect_error);
	}
	$sqlupdate="UPDATE blog_posts SET CONTENT='".$content."' WHERE ID='".$id."'";
	if ($db->query($sqlupdate) === TRUE) {
		echo "<div class='container'>Successfully Updated: ".$id."</div>";
	} else {
		die("Update Error: ".$db->error);
	}
	$db->close();
}
if($_POST['submit']) { updatePost();}
?>
</head>
<body>
<?php
if ($_GET["action"] == 'edit') {
	$db=new mysqli(HOST,USER,PASSWORD,DATABASE);
	if ($db->connect_error) {
		die("Connection Failed: " . $db->connect_error);
	}
	$sqlquery="SELECT * FROM blog_posts WHERE ID='".$_GET['id']."'";
	$result=$db->query($sqlquery);
	$row=$result->fetch_assoc();
	?>
	<div class='container'>
	<form action='<?php echo $_SERVER['PHP_SELF'] ?>' method='post' onsubmit="return postContent()">
	<div class='title'>
	<label>Title: </label>
	<input type='text' name='title' style='background:rgba(0,0,0,0);outline:none;' value='<?php echo $row['TITLE'] ?>'/><br>
	</div><br>
	<textarea  id="summernote" class='content' name='content'><?php echo $row['CONTENT'] ?></textarea><br>
	<div class='post_footer'>
	<input type='submit' value='Submit' name='submit'/>
	</div></form></div>
<?php
	$db->close();
} elseif ($_GET["action"] == 'remove') {
	$id=$_GET['id'];
	$db=new mysqli(HOST,USER,PASSWORD,DATABASE);
	if ($db->connect_error) {
		die("Connection Failed: " . $db->connect_error);
	}
	$sqlupdate="DELETE FROM blog_posts WHERE ID='".$id."'";
	if ($db->query($sqlupdate) === TRUE) {
		echo "<div class='container'>Successfully Deleted: ".$id."</div>";
		header("refresh:3;url=index.php");	
	} else {
		die("Update Error: ".$db->error);
	}
	$db->close();
}
?>
</body>
</html>
