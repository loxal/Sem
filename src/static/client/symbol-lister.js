/*
 * Copyright 2012 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

function SymbolTable(symbolFrom, symbolTo) {
    this.symbolFrom = symbolFrom;
    this.symbolTo = symbolTo;

    var HEX_NUM_BASE = 16;


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
