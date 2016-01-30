$(function () {
		  $('input#site-header-search')
		    .focus(function () { this.select(); })
		    .mouseup(function (e) { e.preventDefault(); })
		    .autocomplete({
		      position: {
		        my: "left top",
		        at: "left bottom",
		        offset: "0, 5",
		        collision: "none"
		      },
		      source: function (request, response) {
		        $.ajax({
		          url: "http://search.scu.edu/suggest?q=" + request.term + "&max=10&site=default_collection&client=default_frontend&format=rich",
		          dataType: "jsonp",
		          success: function (data) {
		            response($.map(data.results, function (item) {
		              return {
		                label: item.name,
		                value: item.name
		              };
		            }));
		          }
		        });
		      },
		      autoFill: true,
		      minChars: 2,
		      select: function (event, ui) {
		        $('#site-header-search').val(ui.item.value);
		        $('#site-header-search-form').trigger('submit');
		      }
		    });
		});	

