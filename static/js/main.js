$(function () {
    $("#cb-img-view").on("closed.mdui.dialog", function () {
        $("#cb-img-view").html("");
    });
    $("#cb-video-view").on("closed.mdui.dialog", function () {
        $("#cb-video-view").html("");
    });
    $("#cb-content-view").on("closed.mdui.dialog", function () {
        $("#cb-content-view pre code").text("");
    });

    $(".cb-detail-btn").on('click', function (event) {
        event.preventDefault();

        let isImg = $(this).attr("attr-isImg");
        let isVideo = $(this).attr("attr-isVideo");

        let path = "/getFileDetail?path=" + $(this).attr("attr-href");

        var loadinginst = new mdui.Dialog("#cb-global-loading-view", {modal: true,});
        loadinginst.open();

        if (isImg === "true") {
            let img = new Image();
            img.onerror = function () {
                loadinginst.close();
            };
            img.onload = function () {
                loadinginst.close();
                $("#cb-img-view").html(img);
                var inst = new mdui.Dialog("#cb-img-view");
                inst.open();
            };
            img.src = path;
            return
        }
        if (isVideo === "true") {
            loadinginst.close();
            $("#cb-video-view").append(`<video src="${path}" controls="controls"></video>`);
            var inst = new mdui.Dialog("#cb-video-view");
            inst.open();
            return;
        }


        var size = parseInt($(this).attr("attr-size"));
        if (size > 5 * 1024 * 1024) {
            toast("文件不能超过5M");
            loadinginst.close();
            return;
        }

        $.get(path, function (data) {
            loadinginst.close();
            if (data.code === -1000) {
                toast(data.message);
                return
            }
            $("#cb-content-view pre code").text(data.toString());
            $("#cb-content-view pre code").each(function (i, block) {
                hljs.highlightBlock(block);
            });
            $("#cb-content-view pre code").each(function () {
                $(this).html("<ul><li>" + $(this).html().replace(/\n/g, "\n</li><li>") + "\n</li></ul>");
            });
            var inst = new mdui.Dialog("#cb-content-view");
            inst.open();
        }, "text")
    });

    $(".back-btn").on("click", function () {
        history.back();
    });

    // 上传
    $("#cb-upload-input").on("change", function (event) {
        var inst = new mdui.Dialog("#cb-loading-view", {modal: true,});
        inst.open();

        var files = $('#cb-upload-input').prop('files');
        var data = new FormData();
        data.append("basePath", location.pathname);
        for (var i = 0; i < files.length; i++) {
            data.append("files", files[i]);
        }

        $.ajax({
            url: "/upload",
            type: "POST",
            data: data,
            cache: false,
            processData: false,
            contentType: false,
            xhr: function () {
                myXhr = $.ajaxSettings.xhr();
                if (myXhr.upload) {
                    myXhr.upload.addEventListener('progress', progressHandlingFunction, false);
                }
                return myXhr;
            },
            success: function (res) {
                if (res.code === 200) {
                    toast("上传成功！");
                    setTimeout(function () {
                        location.reload();
                    }, 1000);
                } else {
                    toast("上传失败！" + res.message);
                }
            },
            error: function (request, textStatus, errorThrown) {
                toast(textStatus + ":" + errorThrown.toString())
            },
            complete: function () {
                inst.close();
            }
        });
    });

    // 创建文件夹
    $(".cb-create-item").on("click", function () {
        let inst = new mdui.Dialog("#cb-create-view");
        inst.open();
    });

    /// 删除
    let filePath = $(this).attr("attr-filePath");
    $(".cb-delete-btn").on("click", function () {
        filePath = $(this).attr("attr-filePath");
        let inst = new mdui.Dialog("#cb-remove-tips");
        inst.open();
    });

    // 确认删除
    $(".cb-delete-confirm").on("click", function () {
        let query = "path=" + encodeURIComponent(filePath);
        $.get("/remove?" + query, function (data) {
            if (data.code === -1000) {
                toast(data.message);
            } else {
                location.reload();
            }
        })
    })
});

function toast(message) {
    mdui.snackbar(message, {
        position: "right-top",
    });
}

function progressHandlingFunction(e) {
    if (e.lengthComputable) {
        var percent = Math.floor(e.loaded / e.total * 100);
        $("#cb-loading-view .mdui-progress-determinate").css("width", percent + "%");
    }
}