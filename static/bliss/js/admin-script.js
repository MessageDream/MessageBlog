function bl_open_uploader(element, class_name){

	$that = jQuery(element);
    wp.media.editor.send.attachment = function(props, attachment){

    	$that.prev().val(attachment.url);

    	jQuery('.'+class_name + ' > img').remove();
    	jQuery('.'+class_name).prepend('<img src="'+attachment.url+'" style="width:100%">');
    }

    wp.media.editor.open(this);

    return false;
}

jQuery(function() {
	jQuery('.cf-nav a').click(function(){
		jQuery('.bl_status_show').hide();
	});
	jQuery('.cf-nav a[href="#post-format-status"]').click(function(){
		console.log(tinyMCE.activeEditor.selection.getContent());
		if(tinyMCE.activeEditor.getContent() == ''){
			tinyMCE.activeEditor.selection.setContent('Status Post');
		}
		jQuery('.bl_status_show').show();
	});
	if(!jQuery('.cf-nav a[href="#post-format-status"]').hasClass('current')){
		jQuery('.bl_status_show').hide();
	}
});