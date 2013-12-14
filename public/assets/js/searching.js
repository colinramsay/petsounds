define(['jquery'], function($) {

	function onSearchSubmit(event) {
        event.preventDefault();

        var searching = $.post('/search', { artist: $(this).find('input').val() });

        searching.done(function(html) {
            $('#results').html(html);
        });

        return false;
    }

    return {
    	init: function() {
    		$('form[action$="search"]').submit(onSearchSubmit);
    	}
    }
});