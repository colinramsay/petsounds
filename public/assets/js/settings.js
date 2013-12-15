define(['jquery'], function($) {

	function onShowSettingsClick(event) {
		event.preventDefault();

		var settingsReq = $.get($(this).attr('href'));

		settingsReq.done(function(response) {
			$('#settings').remove();

			$('body').append(response);
		});
	}

	return {
		init: function() {
			$('header a').click(onShowSettingsClick);
		}
	}
});