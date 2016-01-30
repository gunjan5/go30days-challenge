$(function() {
  var $gallery = $('.gallery.module'),
      $image   = $gallery.find('.full-image .gallery-image img'),
      $caption = $gallery.find('.full-image .gallery-caption');


  $gallery.on('click', '.slides .gallery-image', function(e) {
    e.preventDefault();

    var $this = $(this),
        img_src = $this.find('.thumbnail img').attr('src'),
        caption = $this.find('.gallery-caption').html() || '';

    console.log($this.find('.gallery-image img').length);

    $image.fadeOut(100, function() {
      $image.attr('src', img_src);
    }).fadeIn(100);

    $caption.html(caption);

    return true;
  });
  
  // Show thumbnail scrollbar once images load
  $(window).load(function(){
      $gallery.find('.slides').mCustomScrollbar().show();
  });
});
