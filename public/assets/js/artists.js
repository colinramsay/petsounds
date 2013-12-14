define(['jquery'], function($) {
	function onArtistClick(event) {
        event.preventDefault();
        var releasesQuery = $.get(this.href);

        releasesQuery.done(function(html) {
            $('#releases').html(html);
        });
    }

	return {
		init: function() {
			$('body').on('click', '#results a', onArtistClick);
		}
	}
});