$(function () {
  $(".detail-btn").on('click', function (event) {
    event.preventDefault();
    let fenge = "/"
    if (location.pathname == "/") {
      fenge = ""
    }
    location.href = location.pathname + fenge + $(this).attr("attr-href");
  });

  $("#upload-input").on("change", function (event) {
    let formData = new FormData();
    for(let i=0;i<$(this)[0].files.length;i++){
      formData.append("files", $(this)[0].files[i]);
    }
    formData.append("uploadPath", location.pathname);
    $.ajax({
      url: '/upload', /*接口域名地址*/
      type: 'post',
      data: formData,
      contentType: false,
      processData: false,
      success: function (res) {
        console.log(res);
        if (res["code"] == 200) {
          alert('成功');
        } else {
          alert('失败');
        }
      }
    })
  })
})
