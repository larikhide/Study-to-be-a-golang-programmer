<!DOCTYPE html>
<html>
    <head>
        <meta charset="utf-8">
        <title>Личный сайт студента GeekBrains</title>
        <link rel="stylesheet" href="style.css">
    </head>
    <body>
<div class="content">
    <?php
        include "menu.php";
    ?>
    <h1>Личный сайт студента GeekBrains</h1>

    <div class="center">
    <img src="img/photo2.png" alt="Здесь должно быть фото">
        <div class="box_text">

            <p><b>Добрый день!</b><br> Меня зовут <i>Илон Таск</i>. Я был разочарован началом нового сезона мультсериала Рик и Морти. Поэтому хочу освоить новую профессию, и сейчас Вы видите мою первую попытку по написанию сайта. </p>

            <p>Сайт создан при выполнении заданий с IT-портала <a href="https://geekbrains.ru/"> GeekBrains.

            <p>Здесь вы можете сыграть в несколько игр, которые я написал, по ходу выполнения подготовительного курса: <br>

            <?php
            include "menu.php";
            ?>

            </p>
        </div>
    </div>
</div>


    <div class="footer">
        Copyright &copy; <?php echo date("Y");?> Elon Tusk
    </div>


    </body>
</html>
