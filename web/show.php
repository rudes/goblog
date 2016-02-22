<?php
	if(empty($_GET['id'])) {
		header("Location: index.php");
	}
?>
<!DOCTYPE html>
<html>
<head>
<title>Other Letters</title>
<meta name="viewport" content="width=device-width, initial-scale=1">
<script type="text/javascript" src="//code.jquery.com/jquery-2.1.0.js"></script>
<link rel="stylesheet" type="text/css" href="styles.css">
</head>
<body>
<div class="header"><div class="web_title">
<span id="web_title_text">Letters From the Other Side</span><br>
<span id="web_subtitle">The Mailbox Has Been Discovered</span>
</div></div><br><br><br>
<?php
	include_once 'sqlconfig.php';
	$conn=new mysqli(HOST,USER,PASSWORD,DATABASE);
	if ($conn->connect_error) {
		die("Connection Failed: " . $conn->connect_error);
	}
	$sql="SELECT ID,TITLE,CONTENT,DATE,TIME FROM blog_posts WHERE ID='".$_GET['id']."'";
	$result=$conn->query($sql);
	if ($result->num_rows > 0) {
		echo "<div class='container'>";
		while($row=$result->fetch_assoc()) {
			echo "<div class='title' id='".$row["ID"]."'>";
			echo $row["TITLE"];
			echo "</div><div class='content'>";
			echo $row["CONTENT"];
			echo "</div><div class='post_footer'>";
			echo "Date Posted: ".$row["DATE"]."<a class='footer' href='index.php#".$_GET['id']."'>Back</a>";
			echo "</div>";
		}
		echo "</div>";
	} else {
		echo "0 Posts Found.";
	}
	$result->free();
	$conn->close();
?>
</body>
</html>
