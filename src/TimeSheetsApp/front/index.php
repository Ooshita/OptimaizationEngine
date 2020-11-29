<?php
    if(isset($_GET['person'], $_GET['days'])) {
      $fp = fopen("/parameter/share/parameter.txt", "w");
      $write_str = $_GET['person'].",".$_GET['days'];
      fwrite($fp, $write_str);
      fclose($fp);
    }
?>
<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8" />
  <title>sample</title>
  <link href="./css/style.css" rel="stylesheet">
  <script language="javascript" type="text/javascript">
    function OnButtonClick() {
        fetch("http://localhost:5000", {
            method: "GET",
            mode: "cors",
            })
            .then(response => response.text())
            .then(text => {
                document.getElementById('output').innerHTML = text
        }).catch(err => {
            console.error(err);
        });
    }

  </script>
  </head>
<body>
  <form name="formname" id="id_form" action="">
    <label id="001">従業員数（人） : </label>
    <input type="text" name="person" size="5" maxlenglabel="2" value="">
    </br>
    <label id="002">勤務日数（日）: </label>
    <input type="text" name="days" size="5" maxlenglabel="2" value="">
    </br>
    </br>
    <input type="submit" value="パラメータの保存" />
    </br>
  </form>
  <input type="button" value="タブーサーチスタート!!" onclick="OnButtonClick();"/>
  <br />
  <br />
  <div id="output"></div>
</body>
</html>