CKEDITOR.dialog.add( 'highlight_js', function( editor ) {
    return {
        title: editor.lang.highlight_js.highlight.dialogTitle,
        minWidth: 800,
        minHeight: 500,
        contents: [
            {
                id: 'tab-basic',
                label: 'Basic Settings',
                elements: [
                    {
                        type: 'select',
                        id: 'code_lan',
                        label: editor.lang.highlight_js.highlight.selectLabel,
                        items: [
                                [ 'Auto', 'auto' ],
                                [ 'no-highlight', 'no-highlight' ],
                                ['--------------popular----------------', 'popular'],
                                [ 'ActionScript', 'actionscript' ],
                                [ 'Bash', 'bash' ],
                                [ 'Brainfuck', 'brainfuck' ],
                                [ 'C#', 'cs' ],
                                [ 'C++', 'cpp' ],
                                [ 'CSS', 'css' ],
                                [ 'CoffeeScript', 'coffeescript' ],
                                [ 'D', 'd' ],
                                [ 'Delphi', 'delphi' ],
                                [ 'Django', 'django' ],
                                [ 'DOS .bat', 'dos'],
                                [ 'Erlang', 'erlang' ],
                                [ 'Go', 'go' ],
                                [ 'HTML', 'xml' ],
                                [ 'HTTP', 'http' ],
                                [ 'Haskell', 'haskell' ],
                                [ 'ini', 'ini' ],
                                [ 'JAVA', 'java' ],
                                [ 'JSON', 'json' ],
                                [ 'JavaScript', 'javascript' ],
                                [ 'Lisp', 'lisp' ],
                                [ 'Lua', 'lua' ],
                                [ 'Markdown', 'markdown' ],
                                [ 'Objective C', 'objectivec' ],
                                [ 'PHP', 'php' ],
                                [ 'Perl', 'perl' ],
                                [ 'Python', 'python' ],
                                [ 'Python profile', 'profile' ],
                                [ 'R', 'r' ],
                                [ 'Ruby', 'ruby' ],
                                [ 'SQL', 'sql' ],
                                [ 'Scala', 'scala' ],
                                [ 'Smalltalk', 'smalltalk' ],
                                [ 'TeX', 'tex' ],
                                [ 'XML', 'xml' ],
                                ['--------------other------------------', 'other'],
                                [ '1C ', '1c' ],
                                [ 'AVR Assembler', 'avrasm"' ],
                                [ 'Apache', 'apache' ],
                                [ 'AppleScript', 'applescript' ],
                                [ 'Axapta', 'axapta' ],
                                [ 'CMake', 'cmake' ],
                                [ 'Clojure', 'clojure' ],
                                [ 'Diff', 'diff' ],
                                [ 'MEL', 'mel' ],
                                [ 'Matlab', 'matlab' ],
                                [ 'Nginx', 'ngix' ],
                                [ 'OpenGL Shading Language', 'glsl' ],
                                [ 'Parser3', 'parser3' ],
                                [ 'RenderMan RIB', 'rib' ],
                                [ 'RenderMan RSL', 'rsl' ],
                                [ 'Rust', 'rust' ],
                                [ 'VBScript', 'vbscript' ],
                                [ 'VHDL', 'vhdl' ],
                                [ 'Vala', 'vala' ]
                              ],
                        'default': 'auto',
                        onChange: function(api) {
                            if(this.getValue() === 'popular' || this.getValue() === 'other'){
                                this.setValue(previous_value);
                            } else{
                                previous_value = this.getValue();
                            }
                        },
                        setup: function(element){
                            this.setValue(element.getAttribute( "class"));
                        },
                        commit: function(element){
                            var class_code = this.getValue();
                            if(class_code!=='auto'){
                                element.setAttribute("class", this.getValue());
                            }else{
                                element.removeAttribute( 'class' );
                            }
                        }
                    },
                    {
                        type: 'textarea',
                        id: 'pre_code',
                        rows: 20,
                        style: "width: 100%",
                        setup: function(element) {
                            this.setValue( element.getText() );
                        },
                        commit: function(element) {
                            element.setText( this.getValue() );
                        }
                    }
                ]
            }

        ],
        onShow: function(){
            var selection = editor.getSelection();
            var element = selection.getStartElement();

            if (element){
                element = element.getAscendant( 'code', true );
            }

            if (!element || element.getName() !== 'code'){
                previous_value = 'auto';
                element = editor.document.createElement('code');
                this.insertMode = true;
            }
            else{
                this.insertMode = false;
            }

            this.element = element;

            if (!this.insertMode){
                this.setupContent(this.element);
            }
        },
        onOk: function() {
            var dialog = this;
            var code = this.element;

            this.commitContent(code);

            if (this.insertMode){
                var pre = editor.document.createElement('pre');
                pre.append(code);

                editor.insertElement(pre);
            }
        }
    };
});
