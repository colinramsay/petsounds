define(['jquery'], function($) {

	function onShowSettingsClick(event) {
		event.preventDefault();

		var settingsReq = $.get($(this).attr('href'));

		settingsReq.done(function(response) {
			$('#settings').remove();

			$('body').append(response);
		});
	}

	function onSettingsSave(event) {
		event.preventDefault();

		var saveRequest = $.post($(this).attr('action'), $(this).serialize());

		saveRequest.done(function(response) {
			alert('Saved settings!');
			$('#settings').remove();
		});

		saveRequest.fail(function(response) {
			alert('Settings could not be saved!');
		});
	}

	function onSettingsCancel(event) {
		event.preventDefault();

		$('#settings').remove();
	}

	return {
		init: function() {
			$('header a').click(onShowSettingsClick);

			$('body').on('submit', '#settings form', onSettingsSave);
			$('body').on('click', '#settings button[type=button]', onSettingsCancel);
		}
	}
});