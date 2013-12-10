/**
 * Function toggles code block visibility with some animation
 * 
 * @param blockNumber
 *            int Block number to operate on
 */
function wpsh_toggleBlock(blockNumber) {
	// Toggling visibility with an effect
	jQuery("#wpshdi_" + blockNumber).slideToggle();

	// Changing title style
	var titleBlock = jQuery("#wpshdt_" + blockNumber);
	if (titleBlock.hasClass("wp-synhighlighter-collapsed")) {
		titleBlock.attr("class", "wp-synhighlighter-expanded");
	} else {
		titleBlock.attr("class", "wp-synhighlighter-collapsed");
	}
	return false;
}

/**
 * Function prints code in a given block by opening a new window and printing
 * from there
 * 
 * @param blockNumber
 * @return
 */
function wpsh_print(blockNumber) {
	var newwin = window.open('', 'printwin',
			'left=100,top=100,width=400,height=400');
	newwin.document.write('<html>\n<head>\n');
	newwin.document
			.write('<title>Printed code version produced by WordPress WP-SynHighlight plugin</title>\n');
	newwin.document.write('<script>\n');
	newwin.document.write('function chkstate(){\n');
	newwin.document.write('if(document.readyState=="complete"){\n');
	newwin.document.write('window.close()\n');
	newwin.document.write('}\n');
	newwin.document.write('else{\n');
	newwin.document.write('setTimeout("chkstate()",2000)\n');
	newwin.document.write('}\n');
	newwin.document.write('}\n');
	newwin.document.write('function print_win(){\n');
	newwin.document.write('window.print();\n');
	newwin.document.write('chkstate();\n');
	newwin.document.write('}\n');
	newwin.document.write('<\/script>\n');
	newwin.document.write('</head>\n');
	newwin.document.write('<body onload="print_win()">\n');
	newwin.document.write(jQuery("#wpshdi_" + blockNumber).html());
	newwin.document.write('</body>\n');
	newwin.document.write('</html>\n');
	newwin.document.close();
}

/**
 * Function shows code for a given block in a new window
 * 
 * @param blockNumber
 * @return
 */
function wpsh_code(blockNumber) {
	var newwin = window.open('', 'printwin',
			'left=100,top=100,width=600,height=400,scrollbars=yes');
	newwin.document.write('<html>\n<head>\n')
	newwin.document
			.write('<title>Code only version produced by WordPress WP-SynHighlight plugin</title>\n')
	newwin.document.write('<body>')
	newwin.document.write(jQuery("#wpshdi_" + blockNumber).html());
	newwin.document.write('</body>\n')
	newwin.document.write('</html>\n')
	newwin.document.close()
}