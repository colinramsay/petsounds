require.config({
    baseUrl: 'assets',
    paths: {
        jquery: 'lib/jquery-2.0.3.min'
    }
});

require(['jquery', 'js/searching', 'js/artists', 'js/releases', 'js/settings'], function($, searching, artists, releases, settings) {
    $(function() {
        settings.init();
    	searching.init();
        artists.init();
        releases.init();
    });

    $(document).bind("ajaxSend", function(){
		$("#loading").show();
	}).bind("ajaxComplete", function(){
		$("#loading").hide();
	});
});