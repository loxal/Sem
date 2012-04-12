// Generated Dart class from HTML template.
// DO NOT EDIT.

class EntityViewer {
  final Map<String, Object> _scopes;
  Element _fragment;

  int number;

  EntityViewer(this.number) : _scopes = new Map<String, Object>() {
    // Insure stylesheet for template exist in the document.
    add_foo_templatesStyles();

    _fragment = new DocumentFragment();
    var e0 = new Element.html('<fieldset></fieldset>');
    _fragment.elements.add(e0);
    var e1 = new Element.html('<legend>Symbol Display</legend>');
    e0.elements.add(e1);
    var e2 = new Element.html('<span style="font-size: 1em;" class="viewBox">&#x2724;</span>');
    e0.elements.add(e2);
    var e3 = new Element.html('<span style="font-size: 2em;" class="viewBox">&#x2658;</span>');
    e0.elements.add(e3);
    var e4 = new Element.html('<span style="font-size: 3em;" class="viewBox">&#x2620;</span>');
    e0.elements.add(e4);
    var e5 = new Element.html('<span style="font-size: 6em;" class="viewBox">&#x2624;</span>');
    e0.elements.add(e5);
    var e6 = new Element.html('<label for="symbolId">Symbol Id:</label>');
    e0.elements.add(e6);
    var e7 = new Element.html('<input value="x274a" type="text" id="symbolId">&nbsp;</input>');
    e0.elements.add(e7);
    var e8 = new Element.html('<p>E.g. "10046" in </p>');
    e0.elements.add(e8);
    var e9 = new Element.html('<em>decimal</em> notation or "');
    e8.elements.add(e9);
    var e10 = new Element.html('<strong>x</strong>27bd" in ');
    e8.elements.add(e10);
    var e11 = new Element.html('<em>hexadecimal</em> notation');
    e8.elements.add(e11);
    var e12 = new Element.html('<button id="symbol-display" class="icon-list">Display Symbol</button>');
    e0.elements.add(e12);
  }

  Element get root() => _fragment;

  // CSS class selectors for this template.
  static String get viewBox() => "viewBox";

  // Injection functions:
  // Each functions:

  // With functions:

  // CSS for this template.
  static final String stylesheet = '''
    
.viewBox {
  padding: .1em;
  margin: 0 .1em 0 .1em;
  float: right;
  border-radius: .2em;
  border: 1px solid hsl(33, 55%, 55%);
}

  ''';

  // Stylesheet class selectors:
}


// Inject all templates stylesheet once into the head.
bool foo_stylesheet_added = false;
void add_foo_templatesStyles() {
  if (!foo_stylesheet_added) {
    StringBuffer styles = new StringBuffer();

    // All templates stylesheet.
    styles.add(EntityViewer.stylesheet);

    foo_stylesheet_added = true;
    document.head.elements.add(new Element.html('<style>${styles.toString()}</style>'));
  }
}
