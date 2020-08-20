$(function () {
    $(".cb-detail-btn").on('click', function (event) {
        event.preventDefault();
        // location.href = location.origin + $(this).attr("attr-href");
        // if($(this).attr("attr-isDir")=="true"){
        //   console.log(location.origin , $(this).attr("attr-href"))
        //   location.href = location.origin + $(this).attr("attr-href");
        //   return
        // }
        // if($(this).attr("attr-isImg")=="true") {
        //   $(".img-content").attr("src",location.origin + $(this).attr("attr-href"))
        //   $(".img-view").fadeIn();
        //
        //   return
        // }
        var size = parseInt($(this).attr("attr-size"));
        if (size > 5 * 1024 * 1024) {
            toast("文件不能超过5M");
            return;
        }
        $.get("/getFileDetail?path=" + $(this).attr("attr-href"), function (data) {
            if (data.code === -1000){
                toast(data.message);
                return
            }
            var inst = new mdui.Dialog("#cb-content-view");
            $("#cb-content-view").text(data);
            inst.open();
        })
    });

    $(".back-btn").on("click", function () {
        history.back();
    });

    $(".img-view").on("click", function () {
        $(this).fadeOut()
    });

    $(".close-btn").on("click", function () {
        $(this).parent().fadeOut();
    });

    $("#cb-upload-input").on("change", function (event) {
        $(".cb-upload-form").submit()
    });

    $(".cb-create-item").on("click", function () {
        var inst = new mdui.Dialog("#cb-create-view");
        inst.open();
    });

    $(".cb-delete-btn").on("click", function () {
        var form = $("<form action='/remove' method='post'></form>")
        form.append("<input type='hidden' name='basePath' value='" + $(this).attr("attr-basePath") + "'>");
        form.append("<input type='hidden' name='fileName' value='" + $(this).attr("attr-fileName") + "'>");
        $(document.body).append(form);
        form.submit();
    })

});

function toast(message) {
    mdui.snackbar(message, {
        position: "right-top",
    });
}
