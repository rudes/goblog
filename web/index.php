<!DOCTYPE html>
<html>
<title>Other Letters</title>
<meta name="viewport" content="width=device-width, initial-scale=1">
<script type="text/javascript" src="//code.jquery.com/jquery-2.1.0.js"></script>
<link rel="stylesheet" type="text/css" href="styles.css">
</head>
<body>
<?php
	session_start();
	$login_session=$_SESSION["login_user"];
	include_once 'header.php';
	include_once 'sqlconfig.php';
	$conn=new mysqli(HOST,USER,PASSWORD,DATABASE);
	if ($conn->connect_error) {
		die("Connection Failed: " . $conn->connect_error);
	}
	$sql="SELECT ID,TITLE,CONTENT,DATE,TIME FROM blog_posts ORDER BY DATE DESC, TIME DESC";
	$result=$conn->query($sql);
	if ($result->num_rows > 0) {
		while($row=$result->fetch_assoc()) {
		echo "<div class='container'>";
			echo "<div class='title' id='".$row["ID"]."'>";
			echo "<a href='show.php?id=".$row["ID"]."' class='title'>".$row["TITLE"]."</a>";
			echo "</div><div class='content'>";
			echo $row["CONTENT"];
			echo "</div><div class='post_footer'>";
			echo "Date Posted: ".$row["DATE"]."<a class='footer' href='index.php#home'>Back to Top</a>";
			if(isset($login_session)) { 
				echo "<br><a class='edit' href='edit.php?action=edit&id=".$row["ID"]."'>Edit</a>";
				echo "<a class='footer' href='edit.php?action=remove&id=".$row["ID"]."'>Delete </a>"; 
			}
			echo "</div>";
		echo "</div>";
		echo "<br>";
		}
	} else {
		echo "0 Posts Found.";
	}
	$result->free();
	$conn->close();
?>
</body>
</html>
