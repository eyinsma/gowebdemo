<html>
<head>
<title>Upload File</title>
</head>

<form enctype="multipart/form-data" action="http://localhost:8090/upload" method="post">
    <input type="file" name="browsefile"/>
    <input type="hidden" name = "token" value="{{.}}"/>
    <input type = "submit" value="upload"/>

</form>

</html>