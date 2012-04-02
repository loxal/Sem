/*
 * Copyright 2012 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

function SymbolTable(symbolFrom, symbolTo) {
    this.symbolFrom = symbolFrom;
    this.symbolTo = symbolTo;

    var HEX_NUM_BASE = 16;

    document.writeln('<table>');

    document.writeln('<thead>');
    document.write('<tr>');
    document.write('<th>#</th>');
    document.write('<th>Symbol</th>');
    document.write('<th>Decimal Notation</th>');
    document.write('<th>Hexadecimal Notation</th>');
    document.writeln('</tr>');
    document.writeln('</thead>');

    document.writeln('<tfoot>');
    document.write('<tr>');
    document.write('<td colspan="2"><strong>Total Symbols:</strong> ' + eval(symbolTo - symbolFrom + 1) + '</td>');
    document.write('<td><strong>Decimal Range:</strong> ' + symbolFrom + ' - ' + symbolTo + '</td>');
    document.write('<td><strong>Hexadecimal Range:</strong> ' + symbolFrom.toString(HEX_NUM_BASE) + ' - ' + symbolTo.toString(HEX_NUM_BASE) + '</td>');
    document.writeln('</tr>');
    document.writeln('</tfoot>');

    document.writeln('<tbody>');
    symbolNum = 1;
    for (symbolId = symbolFrom; symbolId <= symbolTo; symbolId++) {
        document.write('<tr>');
        document.write('<td>' + symbolNum++ + '</td>');
        document.write('<td>&#' + symbolId + ';</td>');
        document.write('<td>&amp;#' + symbolId + ';</td>');
        document.write('<td>&amp;#x' + symbolId.toString(HEX_NUM_BASE) + ';</td>');
        document.writeln('</tr>');
    }
    document.writeln('</tbody>');

    document.writeln('</table>');
}

//default values have symbolTo be set
symbolFrom = parseInt(document.getElementById('symbol-from').value);
symbolTo = parseInt(document.getElementById('symbol-to').value);

//create&display table
SymbolTable(symbolFrom, symbolTo);

// assign the new entity to all preview boxes
function displaySymbol() {
    symbolEntities = new Array();
    symbolEntities = document.getElementById('symbol-entity-box').getElementsByTagName('span');

    for (boxIdx = 0; boxIdx < symbolEntities.length; boxIdx++) {
        symbolEntities[boxIdx].innerHTML = '&#' + document
                .getElementById('symbolId').value + ';';
    }
}
