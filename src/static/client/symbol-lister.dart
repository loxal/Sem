/*
 * Copyright 2012 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

#library('symbolLister');
#import('dart:html');
#source('hello.dart');
#source('fruit.dart');

InputElement symbolFrom;
InputElement symbolTo;
LabelElement symbolToLabel;

preinit() {
HeadingElement h1 = new Element.tag('h1');
h1.innerHTML = 'Symbol Lister';
TitleElement title = new Element.tag('title');
title.innerHTML = 'My Title';
document.head.nodes.add(title);

FieldSetElement fieldset = new Element.tag('fieldset');
LegendElement legend = new Element.tag('legend');
legend.innerHTML = 'Range';

LabelElement symbolFromLabel = new Element.tag('label');
symbolFromLabel.innerHTML = 'From:';

symbolFrom = new Element.tag('input');
symbolFrom.type = 'text';
symbolFrom.value = '9985';

symbolTo = new Element.tag('input');
symbolTo.type = 'text';
symbolTo.value = '10175';

symbolToLabel= new Element.tag('label');
symbolToLabel.innerHTML = 'To:';

final ButtonElement refresh = new Element.tag('button');
refresh.innerHTML = 'Refresh';
refresh.classes.add('icon-refresh');
refresh.on.click.add((e) => refreshSymbolList());

fieldset.nodes.add(symbolFromLabel);
fieldset.nodes.add(symbolFrom);
fieldset.nodes.add(symbolToLabel);
fieldset.nodes.add(symbolTo);
fieldset.nodes.add(refresh);

fieldset.nodes.add(legend);

final DivElement main = document.query('#main');
main.nodes.add(fieldset);

final ParagraphElement desc= new Element.tag('p');
desc.innerHTML = '''
        This tool provides an overview over HTML entities that include dingbats and other useful
        symbols. The
        main advantage is that you do not need to carry about any external graphics. You can simply copy&amp;paste the
        corresponding symbol or its HTML entity code to your website or document. And do not forget: These are genuine
        characters, not images!
''';
main.nodes.add(desc);
main.nodes.add(h1);



}

main() {
    preinit();

  final List<String> fruits = ['APPLES', 'ORANGES', 'bananas'];
  final Hello hello = new Hello("Bob", fruits);
  hello.p.on.click.add((e) => print('clicked on paragraph!'));
  document.body.elements.add(hello.root);

    init();
    refreshSymbolList();
}

refreshSymbolList() {
     final int symbolFrom = Math.parseInt(symbolFrom.value);
     final int symbolTo = Math.parseInt(symbolTo.value);

   final Element list = document.query('#list');
   list.nodes.clear();
   int num = 0;
   for (int idx = symbolFrom; idx < symbolTo; idx++) {
      final symbol = new Element.html('<tr><td>' + num++ + '</td><td>' + new String.fromCharCodes([idx])+ '</td><td>' +
       idx
      + '</td></tr>');

     list.elements.add(symbol);

     final Element totalSymbols = document.query('#totalSymbols');
     totalSymbols.innerHTML = (symbolTo - symbolFrom).toString();

     final Element decimalRange = document.query('#decimalRange');
     decimalRange.innerHTML = symbolFrom.toString() + ' - ' + symbolTo.toString();
   }
}

init() {
  document.head.nodes.add(getStylesheet());

    final ButtonElement display = document.body.query('#symbol-display');
    display.on.click.add((e) => displaySymbol());

//    final ButtonElement refresh = document.body.query('#refresh');
//    refresh.on.click.add((e) => refreshSymbolList());

    refreshSymbolList();
}

getStylesheet() {
LinkElement styleSheet = new Element.tag("link");
styleSheet.rel = "stylesheet";
styleSheet.type="text/css";
styleSheet.href="/static/theme/icon/css/font-awesome.css";
return styleSheet;
}

displaySymbol() {
    final List<Element> a = document.queryAll('.symbol');
    final Element symbolId = document.query('#symbolId');
    for(Element e in a) {
      e.innerHTML = '&#' + symbolId.value + ';';
    }
}
