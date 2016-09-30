/**
 * Created by shaohua5 on 16/8/24.
 */
$(function () {
    //ul列表 收起 放下
    $(".toggle_li a").on("click", function () {
        var ele = $(this).siblings("ul");
        var is = ele.is(":hidden");
        if(is == true){
            ele.show(100);
        }else{
            ele.hide(100);
        }
    });

    $("ul li a").on("mousemove", function () {
        $("ul li a").removeClass("add_back_color");
        $(this).addClass("add_back_color");
    });

    $("ul li a").on("click", function () {
        var position = $(this).siblings("input").val();
        if(position != "none"){
            $(".user_position b").text(position);
        }
    });

    //iframe 高度自适应
    $("#iframepage").on("load", function () {
        var newHeight = $(this).contents().find("body").height() + 50;
        $(this).height(newHeight);
    });

});