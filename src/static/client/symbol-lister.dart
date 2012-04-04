/*
 * Copyright 2012 Alexander Orlov <alexander.orlov@loxal.net>. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

#library('test');

#import('dart:html');

main() {
  for (int idx = 9985; idx < 10175; idx++) {
//    print(new String.fromCharCodes([idx]));
  }

    var msg = 'blub';
  var loudify = (msg, testt1) => '!!! ${msg.toUpperCase()} ${testt1} !!!';
  print (loudify("mdddob", 'test!!!!!??'));

   var callbacks = [];
    for (var i = 0; i < 2; i++) {
      callbacks.add(() => print(i));
    }
    callbacks.forEach((c) => c());

    fest(c) => print (c);

    print((e) => print('33'));

    print (fest(23));

//    var t = document.query('#symbol-from');
    Element tt = document.query('#symbol-from');
    tt.value = 'blub';
//    t.value = 'fest';
}
