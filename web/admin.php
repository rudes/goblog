<?php
include('lock.php');
?>
<!DOCTYPE html>
<html>
<head>
<title>Control Panel | Other Letters</title>
<meta name="viewport" content="width=device-width, initial-scale=1">
<script type="text/javascript" src="//code.jquery.com/jquery-2.1.0.js"></script>
<link href="//netdna.bootstrapcdn.com/bootswatch/3.0.1/cosmo/bootstrap.min.css" rel="stylesheet"> 
<script src="//netdna.bootstrapcdn.com/bootstrap/3.0.1/js/bootstrap.min.js"></script> 
<link href="//netdna.bootstrapcdn.com/font-awesome/4.0.3/css/font-awesome.min.css" rel="stylesheet">
<link href="summernote/summernote.css" rel="stylesheet">
<script src="summernote/summernote.min.js"></script>
<link rel="stylesheet" type="text/css" href="styles.css">
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
</head>
<body>
<?php
include_once 'header.php';
function addPost() { 
	include_once 'sqlconfig.php';
	$id=mb_substr(md5($_POST['title']),0,10);
	$title=$_POST['title'];
	//$content=$_POST['content'];
	$content=$_POST['content'];
	$date=date("Y/m/d");
	$time=date("H:i:s");
	$db=new mysqli(HOST,USER,PASSWORD,DATABASE);
	$sqlquery="SELECT ID FROM blog_posts WHERE CONTENT='".$content."'";
	$result=$db->query($sqlquery);
	if ($result->num_rows > 0) {
		echo "<div class='container'>Post has already been submitted.</div>";
	} else {
	$sqlinsert="INSERT INTO blog_posts (ID,TITLE,CONTENT,DATE,TIME) VALUES ('".$id."','".$title."','".$content."','".$date."','".$time."')";
	if ($db->query($sqlinsert) === TRUE) {
		echo "<div class='container'>Successfully Posted: ".$title." under ID: ".$id."</div>";
	} else {
		die("Update Error: ".$db->error);
	}
	}
	$db->close();	
}
if ($_POST['submit']) { addPost(); }
?>
<div class="container">
<form action="<?php echo $_SERVER['PHP_SELF'] ?>" method="post" onsubmit="return postContent()">
<div class="title">
<label>Title:	</label>
<input type="text" name="title" style="background:rgba(0,0,0,0);outline:none;"/><br>
</div><br>
<!--<textarea  class="content" name="content"></textarea><br>-->
<textarea id="summernote" class="content" name="content"><?php if(!empty($_SESSION['content'])) echo $_SESSION['content']; ?></textarea><br>
<div class="post_footer">
<input type="submit" value="Submit" name="submit"/>
</div>
</form>
</div>
</body>
</head>
