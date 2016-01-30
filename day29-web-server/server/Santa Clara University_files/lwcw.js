(function() {
  function isValidJQuery() {
    var result = false;
    // load jquery and plugin if jquery not loaded, otherwise just load plugin
    if (window.jQuery && window.jQuery.fn) { 
      var version_parts = parseFloat(window.jQuery.fn.jquery).toString().split('.');
      if (parseInt(version_parts[0], 10) === 2 || 
          (parseInt(version_parts[0], 10) === 1 && parseInt(version_parts[1], 10) >= 5)) {
        result = true;
      }
    }
    return result;
  }

  function getAggregateUrl(host, liveurl_dir, type, resources) {
    var urls = [];
    // build aggregate resource url
    for (var i = 0; i < resources.length; i++) {
      urls.push( resources[i].replace(/\//g, '%5C') );
    }
    return host + liveurl_dir + '/resource/' + type + '/' + urls.join('/');
  }

  function parseOptions(option_str) { 
    var result = {},
        pairs;
    if (option_str && option_str.indexOf('&')) {
      pairs = option_str.split('&');
      for (var i = 0; i < pairs.length; i++) {
        var pair = pairs[i].split("=");
        result[decodeURIComponent(pair[0])] = decodeURIComponent(pair[1]);
      }
    }
    return result;
  }

  function extend(a, b) {
    for (var key in b) {
      if(b.hasOwnProperty(key)) {
        a[key] = b[key];
      }
    }
    return a;
  }
  
  // dynamically load javascript and execute callback when finished
  function loadScript(url, callback) {
    var script = document.createElement("script");
    script.type = "text/javascript";

    if (script.readyState) { //IE
      script.onreadystatechange = function() {
        if (script.readyState === "loaded" || script.readyState === "complete") {
          script.onreadystatechange = null;
          callback();
        }
      };
    } else { //Others
      script.onload = function() {
        callback();
      };
    }
    script.src = url;
    document.getElementsByTagName('head')[0].appendChild(script);
  }

  function initPlugin($, host) {
    $(document).ready(function() {
      // for each widget, parse options in data-options attr (formatted like qstring), and call plugin
      $('.lwcw').each(function() {
        var $this      = $(this),
            option_str = $this.attr('data-options'),
            options    = { host: host },
            args;

        args = parseOptions(option_str);
        $.extend(options, args);

        // call iframe widget jquery plugin
        $this.lwWidget(options);
      });
    });
  }

  function init() {
    var script = document.getElementById('lw_lwcw'),
        option_str = script.getAttribute('data-options');

    var opts = {
      liveurl_dir: '/live'
    };
    if (option_str) {
      extend(opts, parseOptions(option_str));
    }

    // return right away if script tag that loads this is not found
    if (!script || !script.src) return;

    // get host from script tag's src
    var host = script.src.match(/^\s*https?:\/\/[^\/]+/i);
    if (host.length !== 1) return;
    host = host[0];

    var scripts = [
      '/livewhale/scripts/lib/date/formatter.js',
      '/livewhale/scripts/lib/date/timezone.js',
      '/livewhale/scripts/lib/date/user.js',
      '/livewhale/plugins/jquery/jquery.lw-widget.js'
    ];
    if (!isValidJQuery()) {
      scripts.unshift('/livewhale/thirdparty/jquery/jquery.js');
    }
    var url = getAggregateUrl(host, opts.liveurl_dir, 'js', scripts);
    loadScript(url, function() {
      initPlugin(livewhale.jQuery || jQuery, host);
    });
  }
  init();
}());
