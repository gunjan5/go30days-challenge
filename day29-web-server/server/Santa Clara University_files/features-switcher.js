/**
 * Switcher (departments/majors/minors)
 */

	$(function() {
	  var $switcher = $('.switcher');
	
	  $switcher.on('click',function() {
		  $switcher.toggleClass('on');
		  $switcher.toggleClass('two-column');
	  });
	});
