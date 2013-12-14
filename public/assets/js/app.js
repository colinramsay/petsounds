require.config({
    baseUrl: 'assets',
    paths: {
        jquery: 'lib/jquery-2.0.3.min'
    }
});

require(['jquery', 'js/searching', 'js/artists', 'js/releases'], function($, searching, artists, releases) {
    $(function() {
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