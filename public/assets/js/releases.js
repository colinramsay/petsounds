define(['jquery'], function($) {
	function onReleasesClick(event) {
        event.preventDefault();
        var releasesQuery = $.get(this.href);

        releasesQuery.done(function(response) {
            alert(response);
        });
    }

    return {
    	init: function() {
    		$('body').on('click', '#releases a', onReleasesClick);
    	}
    }
});