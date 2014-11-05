jQuery(function() {
	
	jQuery('.bl_newsletter button').click(function() {

		$that = jQuery(this);
		$that.addClass('disabled');
		
		var email = $that.prev().val();
		if(email == ''){
			jQuery('body').prepend('<div class="bl_alert"><h4 style="text-align:center"><i class="icon-cancel-circle"></i>&nbsp;'+locale.no_email_provided+'</h4></div>');
			jQuery('.bl_alert').slideDown().delay(3000).slideUp();	
			$that.removeClass('disabled');
			return false;
		}

		jQuery.post('/wp-content/themes/bliss/inc/mailchimp/subscribe.php', {email: email}, function(output){
			output = jQuery.trim(output);
			var obj = jQuery.parseJSON(output);
			if(obj.error){
				jQuery('body').prepend('<div class="bl_alert"><h4 style="text-align:center"><i class="icon-cancel-circle"></i>&nbsp;'+obj.error+'</h4></div>');
				jQuery('.bl_alert').slideDown().delay(3000).slideUp();						
			}else if(obj.status == 'ok'){
				jQuery('body').prepend('<div class="bl_alert"><h4 style="text-align:center"><i class="icon-ok-circle"></i>&nbsp;'+locale.thank_you_for_subscribing+'</h4></div>');
				jQuery('.bl_alert').slideDown().delay(3000).slideUp();
			}
		});		
		$that.removeClass('disabled');
		$that.prev().val('');
		return false;
	});
});