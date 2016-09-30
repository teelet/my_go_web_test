/**
 * Created by shaohua5 on 16/8/24.
 */
$(function(){
    //得到焦点
    $("#password").focus(function(){
        $("#left_hand").animate({
            left: "150",
            top: " -38"
        },{step: function(){
            if(parseInt($("#left_hand").css("left"))>140){
                $("#left_hand").attr("class","left_hand");
            }
        }}, 2000);
        $("#right_hand").animate({
            right: "-64",
            top: "-38px"
        },{step: function(){
            if(parseInt($("#right_hand").css("right"))> -70){
                $("#right_hand").attr("class","right_hand");
            }
        }}, 2000);
    });
    //失去焦点
    $("#password").blur(function(){
        $("#left_hand").attr("class","initial_left_hand");
        $("#left_hand").attr("style","left:100px;top:-12px;");
        $("#right_hand").attr("class","initial_right_hand");
        $("#right_hand").attr("style","right:-112px;top:-12px");
    });

    $('.but_a').click(function () {
        $('#checkMsg').text('');
        //验证用户名,密码是否为空
        var username = $('#username').val().trim();
        if(username == ''){
            $('#checkMsg').text('用户名不能为空!');
            $('#username').focus();
            return false;
        }
        var passwrod = $('#password').val().trim();
        if(passwrod == ''){
            $('#checkMsg').text('密码不能为空!');
            $('#password').focus();
            return false;
        }
        //验证用户名密码真实性
        $.ajax({
            'url'     : '/ajax/checklogin',
            'type'    : 'post',
            'dataType': 'json',
            'data'    : {'username' : username, 'password' : $.sha1(passwrod)},
            'success' : function (data) {
                if(data.status == 1){ //非 0 即 错
                    $('#checkMsg').text(data.errorMsg);
                }else{
                    $('#password').val('');
                    $('#login_form').submit();
                }
            },
            'error'   : function () {
                return false;
            }
        });
    });
});
