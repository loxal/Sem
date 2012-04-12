/*
 * Copyright 2012 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

#library('symbolLister');
#import('dart:html');
#import('dart:dom', prefix:'dom');
#source('hello.dart');
#source('fruit.dart');
#source('EntityViewer.dart');

class EntityLister {

InputElement symbolFrom;
InputElement symbolTo;
LabelElement symbolToLabel;
DivElement app;

preinit() {
    final HeadingElement h1 = new Element.tag('h1');
    h1.innerHTML = 'Symbol Lister';
    final TitleElement title = new Element.tag('title');
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
    symbolTo.value = '10000';

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

    app.nodes.add(fieldset);

    app.nodes.add(h1);
}

TableSectionElement tbody;
initContainer() {
    final TableElement container = new Element.tag('table');
    final TableCaptionElement tableCaption = new Element.tag('caption');
    tableCaption.innerHTML = 'Container';
    container.elements.add(tableCaption);
    TableSectionElement thead = new Element.tag('thead');
    TableSectionElement tfoot = new Element.tag('tfoot');
    tbody = new Element.tag('tbody');
    TableRowElement header = new Element.tag('tr');
    TableRowElement footer = new Element.tag('tr');
    totalSymbols = new Element.tag('td');
    decimalRange = new Element.tag('td');
    footer.elements.add(totalSymbols);
    footer.elements.add(decimalRange);
    TableElement thNum = new Element.tag('th');
    thNum.innerHTML = '#';
    TableElement thSymbol = new Element.tag('th');
    thSymbol.innerHTML = 'Symbol';
    TableElement thNotation = new Element.tag('th');
    thNotation.innerHTML = 'Decimal Notation';
    header.elements.add(thNum);
    header.elements.add(thSymbol);
    header.elements.add(thNotation);
    thead.elements.add(header);
    tfoot.elements.add(footer);
    container.elements.add(thead);
    container.elements.add(tbody);
    container.elements.add(tfoot);

    app.elements.add(container);

}

TableCellElement totalSymbols;
TableCellElement decimalRange;
refreshSymbolList() {
     final int symbolFrom = Math.parseInt(symbolFrom.value);
     final int symbolTo = Math.parseInt(symbolTo.value);

   tbody.nodes.clear();
   int num = 1;
   for (int idx = symbolFrom; idx < symbolTo; idx++) {
      final symbol = new Element.html('<tr><td>' + num++ + '</td><td>' + new String.fromCharCodes([idx])+ '</td><td>' +
       idx
      + '</td></tr>');

     tbody.elements.add(symbol);

     totalSymbols.innerHTML = (symbolTo - symbolFrom).toString();

     decimalRange.innerHTML = symbolFrom.toString() + ' - ' + symbolTo.toString();
   }
}

init() {
  document.head.nodes.add(getStylesheet());

    final ButtonElement display = document.body.query('#symbol-display');
    display.on.click.add((e) => displaySymbol());

    refreshSymbolList();
}

getStylesheet() {
    final LinkElement styleSheet = new Element.tag("link");
    styleSheet.rel = "stylesheet";
    styleSheet.type="text/css";
    styleSheet.href="/static/theme/icon/css/font-awesome.css";
    return styleSheet;
}

displaySymbol() {
    final List<Element> a = document.queryAll('.viewBox');
    final Element symbolId = document.query('#symbolId');
    for(Element e in a) {
      e.innerHTML = '&#' + symbolId.value + ';';
    }
}

EntityLister() {
    print('constructed');
}
}

main() {
    final EntityLister my = new EntityLister();
    my.app =  new Element.tag('div');
    document.body.elements.add(my.app);
    my.preinit();
    my.initContainer();

  final List<String> fruits = ['APPLES', 'ORANGES', 'bananas'];
  final Hello hello = new Hello("Bob", fruits);
  hello.p.on.click.add((e) => print('clicked on paragraph!'));
  document.body.elements.add(hello.root);

  EntityViewer entityViewer = new EntityViewer(93);
  document.body.elements.add(entityViewer.root);

    my.init();
    my.refreshSymbolList();
}
