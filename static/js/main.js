$(function () {
  $(".detail-btn").on('click', function (event) {
    event.preventDefault();
    // location.href = location.origin + $(this).attr("attr-href");
    if($(this).attr("attr-isDir")=="true"){
      console.log(location.origin , $(this).attr("attr-href"))
      location.href = location.origin + $(this).attr("attr-href");
      return
    }
    if($(this).attr("attr-isImg")=="true") {
      $(".img-content").attr("src",location.origin + $(this).attr("attr-href"))
      $(".img-view").fadeIn();

      return
    }
    $.get(location.origin + $(this).attr("attr-href"),function (data) {
      $(".file-view").fadeIn();
      $(".content-view").text(data)
    })
  });

  $(".img-view").on("click",function () {
    $(this).fadeOut()
  })

  $(".close-btn").on("click",function () {
    $(this).parent().fadeOut();
  })

  $("#upload-input").on("change", function (event) {
    $(".upload-form").submit()
  })

  $(".create-item").on("click",function () {
    $(".create-view").fadeIn()
  })

  $(".delete-btn").on("click",function () {
    var form = $("<form action='/remove' method='post'></form>")
    form.append("<input type='hidden' name='basePath' value='" + $(this).attr("attr-basePath") +"'>");
    form.append("<input type='hidden' name='fileName' value='" + $(this).attr("attr-fileName") +"'>");
    $(document.body).append(form);
    form.submit();
  })

})
