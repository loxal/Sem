/*
 * Copyright 2012 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

#library('symbolLister');
#import('dart:html');
#source('hello.dart');
#source('fruit.dart');

main() {
  List fruits = ['apples', 'ORANGES', 'bananas'];
  Hello hello = new Hello("Bob", fruits);
  hello.p.on.click.add((e) => print('clicked on paragraph!'));
  document.body.elements.add(hello.root);

//    init();
}

refreshSymbolList() {
     final int symbolFrom = Math.parseInt(document.query('#symbolFrom').value);
     final int symbolTo = Math.parseInt(document.query('#symbolTo').value);

   final Element list = document.query('#list');
   list.nodes.clear();
   int num = 0;
   for (int idx = symbolFrom; idx < symbolTo; idx++) {
      final symbol = new Element.html('<tr><td>' + num++ + '</td><td>' + new String.fromCharCodes([idx])+ '</td><td>' +
       idx
      + '</td></tr>');

     list.nodes.add(symbol);

     final Element totalSymbols = document.query('#totalSymbols');
     totalSymbols.innerHTML = (symbolTo - symbolFrom).toString();

     final Element decimalRange = document.query('#decimalRange');
     decimalRange.innerHTML = symbolFrom.toString() + ' - ' + symbolTo.toString();
   }
}

init() {
    final ButtonElement click = document.body.query('#symbol-display');
    click.on.click.add((e) => displaySymbol());

    final ButtonElement refresh = document.body.query('#refresh');
    refresh.on.click.add((e) => refreshSymbolList());

    refreshSymbolList();
}

displaySymbol() {
    final List<Element> a = document.queryAll('.symbol');
    final Element symbolId = document.query('#symbolId');
    for(Element e in a) {
      e.innerHTML = '&#' + symbolId.value + ';';
    }
}
