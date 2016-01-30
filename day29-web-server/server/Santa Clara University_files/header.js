/*! jRespond.js v 0.10 | Author: Jeremy Fields [jeremy.fields@viget.com], 2013 | License: MIT */

// Universal Module Definition
;(function (window, name, fn) {
  // Node module pattern
    if (typeof module === "object" && module && typeof module.exports === "object") {
        module.exports = fn;
    } else {
    // browser
        window[name] = fn;

        // AMD definition
        if (typeof define === "function" && define.amd) {
            define(name, [], function (module) {
                return fn;
            });
        }
    }
}(this, 'jRespond', function(win,doc,undefined) {

  'use strict';

  return function(breakpoints) {

    // array for registered functions
    var mediaListeners = [];

    // array that corresponds to mediaListeners and holds the current on/off state
    var mediaInit = [];

    // array of media query breakpoints; adjust as needed
    var mediaBreakpoints = breakpoints;

    // store the current breakpoint
    var curr = '';

    // the previous breakpoint
    var prev = '';

    // window resize event timer stuff
    var resizeTimer;
    var resizeW = 0;
    var resizeTmrFast = 100;
    var resizeTmrSlow = 500;
    var resizeTmrSpd = resizeTmrSlow;

    // cross browser window width
    var winWidth = function() {

      var w = 0;

      // IE
      if (typeof( window.innerWidth ) != 'number') {

        if (!(document.documentElement.clientWidth === 0)) {

          // strict mode
          w = document.documentElement.clientWidth;
        } else {

          // quirks mode
          w = document.body.clientWidth;
        }
      } else {

        // w3c
        w = window.innerWidth;
      }

      return w;
    };

    // determine input type
    var addFunction = function(elm) {
      if (elm.length === undefined) {
        addToStack(elm);
      } else {
        for (var i = 0; i < elm.length; i++) {
          addToStack(elm[i]);
        }
      }
    };

    // send media to the mediaListeners array
    var addToStack = function(elm) {
      var brkpt = elm['breakpoint'];
      var entr = elm['enter'] || undefined;

      // add function to stack
      mediaListeners.push(elm);

      // add corresponding entry to mediaInit
      mediaInit.push(false);

      if (testForCurr(brkpt)) {
        if (entr !== undefined) {
          entr.call(null, {entering : curr, exiting : prev});
        }
        mediaInit[(mediaListeners.length - 1)] = true;
      }
    };

    // loops through all registered functions and determines what should be fired
    var cycleThrough = function() {

      var enterArray = [];
      var exitArray = [];

      for (var i = 0; i < mediaListeners.length; i++) {
        var brkpt = mediaListeners[i]['breakpoint'];
        var entr = mediaListeners[i]['enter'] || undefined;
        var exit = mediaListeners[i]['exit'] || undefined;

        if (brkpt === '*') {
          if (entr !== undefined) {
            enterArray.push(entr);
          }
          if (exit !== undefined) {
            exitArray.push(exit);
          }
        } else if (testForCurr(brkpt)) {
          if (entr !== undefined && !mediaInit[i]) {
            enterArray.push(entr);
          }
          mediaInit[i] = true;
        } else {
          if (exit !== undefined && mediaInit[i]) {
            exitArray.push(exit);
          }
          mediaInit[i] = false;
        }
      }

      var eventObject = {
        entering : curr,
        exiting : prev
      };

      // loop through exit functions to call
      for (var j = 0; j < exitArray.length; j++) {
        exitArray[j].call(null, eventObject);
      }

      // then loop through enter functions to call
      for (var k = 0; k < enterArray.length; k++) {
        enterArray[k].call(null, eventObject);
      }
    };

    // checks for the correct breakpoint against the mediaBreakpoints list
    var returnBreakpoint = function(width) {

      var foundBrkpt = false;

      // look for existing breakpoint based on width
      for (var i = 0; i < mediaBreakpoints.length; i++) {

        // if registered breakpoint found, break out of loop
        if (width >= mediaBreakpoints[i]['enter'] && width <= mediaBreakpoints[i]['exit']) {
          foundBrkpt = true;

          break;
        }
      }

      // if breakpoint is found and it's not the current one
      if (foundBrkpt && curr !== mediaBreakpoints[i]['label']) {
        prev = curr;
        curr = mediaBreakpoints[i]['label'];

        // run the loop
        cycleThrough();

      // or if no breakpoint applies
      } else if (!foundBrkpt && curr !== '') {
        curr = '';

        // run the loop
        cycleThrough();
      }

    };

    // takes the breakpoint/s arguement from an object and tests it against the current state
    var testForCurr = function(elm) {

      // if there's an array of breakpoints
      if (typeof elm === 'object') {
        if (elm.join().indexOf(curr) >= 0) {
          return true;
        }

      // if the string is '*' then run at every breakpoint
      } else if (elm === '*') {
        return true;

      // or if it's a single breakpoint
      } else if (typeof elm === 'string') {
        if (curr === elm) {
          return true;
        }
      }
    };

    // self-calling function that checks the browser width and delegates if it detects a change
    var checkResize = function() {

      // get current width
      var w = winWidth();

      // if there is a change speed up the timer and fire the returnBreakpoint function
      if (w !== resizeW) {
        resizeTmrSpd = resizeTmrFast;

        returnBreakpoint(w);

      // otherwise keep on keepin' on
      } else {
        resizeTmrSpd = resizeTmrSlow;
      }

      resizeW = w;

      // calls itself on a setTimeout
      setTimeout(checkResize, resizeTmrSpd);
    };
    checkResize();

    // return
    return {
      addFunc: function(elm) { addFunction(elm); },
      getBreakpoint: function() { return curr; }
    };

  };

}(this,this.document)));


$(document).ready(function(){
  var $header           = $('header.site-header'),
      $headerSlideMenus = $header.find('li.slide-menu'),
      $gatewayNav       = $header.find('.gateway-navigation'),
      $mainNav          = $header.find('.main-navigation'),
      $body             = $('body'),
      $window           = $(window),
      mainNavOffset     = $mainNav.offset();

  var respond = jRespond([
      {
          label: 'mobile',
          enter: 0,
          exit: 767
      },{
          label: 'tablet',
          enter: 768,
          exit: 991
      },{
          label: 'laptop',
          enter: 992,
          exit: 1199
      },{
          label: 'desktop',
          enter: 1200,
          exit: 10000
      }
  ]);
  
  // Offcanvas menus for mobile
  var initMobileMenus = function(){
        var $body = $('body');
        $body.wrapInner('<div id="js-mobile-content-wrapper"></div>');
        $('[data-toggle="offcanvas"]').on('click.mobile',function(){
          var $$=$(this),
              $menu=$($$.data('target')),
              bodyClass = $menu.hasClass('offcanvas-left') ? 'active-left' : $menu.hasClass('offcanvas-right') ? 'active-right' : false;
          if(!bodyClass) { return; }
          $body.toggleClass(bodyClass);
        });
      },
      uninitMobileMenus = function(){
        $('[data-toggle="offcanvas"]').off('.mobile').off('.mobileclose');
        $('body').removeClass('active-left').removeClass('active-right');
        $('#js-mobile-content-wrapper>:first-child').unwrap();
      };
      
  // Add slidedown/slideup animations to header dropdowns using Bootstrap's events:
  $('#scu-main-navigation+nav .scu-nav').hover(function(){
    $body.toggleClass('nav-dropdown-open');
  });

  // Slide menus for tablet/desktop drawers in gateway navigation
  var toggleSlideMenus = function(e){
        e.preventDefault();
        var $$           = $(this).parents('li.slide-menu'),
            $panel       = $$.find('div.slide-panel'),
            panelH       = $panel.outerHeight(),
            $animated    = $mainNav.hasClass('affix')? $gatewayNav.add($mainNav) : $gatewayNav;
        if($$.hasClass('open')){ // We're closing it here
          $panel.find('.row').fadeOut(function(){$panel.hide().find('.row').show();});
                  //$('body').stop().animate({"paddingTop": 0},function(){$$.removeClass('open');});
          $animated.stop().animate({"top":"-="+panelH},{
            complete: function() {
                $$.removeClass('open');
                initMainNavAffix();
                if(!$mainNav.hasClass('affix')){ $mainNav.css('top',''); }
              },
            step: function(now,fx){
              //If ($this is $mainNav) and (now - $(window).scrollTop() <= $mainNav's default position ) then
              var $this = $(this);
              if($this.hasClass('affix') && $this.hasClass('main-navigation') && (now + $(window).scrollTop() <= mainNavOffset.top)) {
                $this.removeClass('affix').addClass('affix-top');
              }
            }
          });
        }else{ // Opening or switching here
          var $otherPanel = $headerSlideMenus.not($$).filter('.open').removeClass('open').find('div.slide-panel');
          if($otherPanel.length){
            panelH -= $otherPanel.outerHeight();
          }
          $animated.stop().animate({"top":"+="+panelH},{
            complete: function(){
                initMainNavAffix();
                if(!$mainNav.hasClass('affix')){ $mainNav.css('top',$gatewayNav.outerHeight()+$gatewayNav.offset().top-$(document).scrollTop()-1+'px') }
                //\\$mainNav.css('top',$gatewayNav.outerHeight()+$gatewayNav.offset().top- $(document).scrollTop()-1+'px');
              },
            step: function(now,fx){
                var $this=$(this);
                // If $animated does not contain $mainNav, and $this = $gatewayNav, and $gatewayNav's bottome edge is past the main nav's top offset, convert main nav to affix and animate its top:
                if(!$animated.filter('.main-navigation').length && $this.hasClass('gateway-navigation') && (now + $gatewayNav.outerHeight() - 1 >= mainNavOffset.top - $(window).scrollTop())) {
                  $mainNav.removeClass('affix-top').addClass('affix').css('top',(now + $gatewayNav.outerHeight() - 1));
                }else if($this.hasClass('affix') && $this.hasClass('main-navigation') && (now + $(window).scrollTop() <= mainNavOffset.top)){ // but if $this is $mainNav, and it gets to mainNavOffset.top (its original position), drop it off where it started.
                  $this.removeClass('affix').addClass('affix-top');
                }
              }
          });
          $otherPanel.fadeOut();
          $$.addClass('open');
          $panel.fadeIn();
        }
      },
      initToggleSlideMenus = function(){
        $headerSlideMenus.on('click.tabletUp', 'a[role="button"]', toggleSlideMenus);
      },
      uninitToggleSlideMenus = function(){
        $headerSlideMenus.off('.tabletUp');
      };
  
  var initMainNavAffix = function(){
    var gOffset = $gatewayNav.offset().top - $(document).scrollTop(),
        reaffixMain = false,
        newTop = (mainNavOffset.top > $gatewayNav.outerHeight() + gOffset - 1) ?
          mainNavOffset.top - $gatewayNav.outerHeight() - gOffset - 1 :
          $gatewayNav.outerHeight() + gOffset - 1;
        
    // Main nav affix:
    if($mainNav.data('bs.affix')){
      $mainNav.data('bs.affix').options.offset.top = newTop;
    }else{
      $mainNav.affix({
        offset: {
          top: newTop
        }
      });
    }
  };

  var adjustHeaderPositions = function(){
    initMainNavAffix();
    if($headerSlideMenus.filter('.open').length){
      $headerSlideMenus.filter('.open').each(function(){
        var $$ = $(this),
            panelHeight = $$.find('div.slide-panel').outerHeight();
        $$.parents('.gateway-navigation').css('top',panelHeight+'px');
        if($mainNav.hasClass('affix')) {
          $mainNav.css('top',panelHeight-1+$gatewayNav.outerHeight()+'px');
        }
      });
    }
  };

  // Hook up init functions with breakpoints:
  mediaCheck({
    media: '(max-width: 767px)', // Mobile only
    entry: function(){
      uninitToggleSlideMenus();
      initMobileMenus();
    },
    exit: function(){
      uninitMobileMenus();
      initToggleSlideMenus();
      setTimeout(function(){
        mainnavOffset = $mainNav.offset();
        initMainNavAffix();
      }, 250);
    }
  });
  mediaCheck({
    media: '(min-width: 1065px) and (max-width: 1199px)',
    entry: adjustHeaderPositions,
    exit: adjustHeaderPositions
  });

  function initHomepageCanWeSwipe() {
    var $box_wrap = $('.can-we-boxes'),
        $boxes    = $box_wrap.find('.can-we-box'),
        margin    = 30,
        box_width, position;

    function init() {
      if ($window.width() >= 768) { return; }
      position = 1;
      $boxes.width($window.width() * 0.5);
      box_width = $boxes.outerWidth();
      $box_wrap.css('left', $body.width() - box_width - margin - 25); 
    }
    function swipeLeft() {
      if (position >= $boxes.length) { return; }
      $box_wrap.animate({ left: $box_wrap.position().left - box_width - margin }, 300);
      position++;
    }
    function swipeRight() {
      if (position === 1) { return; }
      $box_wrap.animate({ left: $box_wrap.position().left + box_width + margin }, 300);
      position--;
    }

    respond.addFunc({
      breakpoint: 'mobile',
      enter: function() {
        init();
        $boxes
          .on('swipeleft', swipeLeft)
          .on('swiperight', swipeRight); 
        $window.on('resize', init);
      },
      exit: function() {
        $window.off('resize', init);
        $boxes.off('swipeleft').off('swiperight');
        $boxes.css('width', 'auto');
        $box_wrap.css('left', 'auto');
      }
    });
  }
  initHomepageCanWeSwipe();
});
