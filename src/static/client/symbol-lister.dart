/*
 * Copyright 2012 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

#library('symbolLister');
#import('dart:html');

main() {
    final int symbolFrom = Math.parseInt(document.query('#symbolFrom').value);
    final int symbolTo = Math.parseInt(document.query('#symbolTo').value);

  final Element list = document.query('#list');
  int num = 0;
  for (int idx = symbolFrom; idx < symbolTo; idx++) {
     final symbol = new Element.html('<tr><td>' + num++ + '</td><td>' + new String.fromCharCodes([idx])+ '</td><td>' +
      idx
     + '</td></tr>');

    list.nodes.add(symbol);


    Element totalSymbols = document.query('#totalSymbols');
    totalSymbols.innerHTML = (symbolTo - symbolFrom).toString();

    Element decimalRange = document.query('#decimalRange');
    decimalRange.innerHTML = symbolFrom.toString() + ' - ' + symbolTo.toString();
  }
}
