CKEDITOR.plugins.add('highlight_js',
{
	requires : [ 'dialog' ],
	lang : [ 'en', 'ru' ],
	init : function(editor)
	{
		var pluginName = 'highlight_js';
		var command = editor.addCommand(pluginName, new CKEDITOR.dialogCommand(pluginName) );
		command.modes = { wysiwyg:1, source:1 };
		//command.canUndo = false;
		editor.ui.addButton('Highlight',
		{
				label : editor.lang.highlight_js.highlight.title,
				command : pluginName,
				icon: this.path + 'icons/highlight_js.gif'
		});

		CKEDITOR.dialog.add(pluginName, this.path + 'dialogs/highlight_js.js' );
	}
});

