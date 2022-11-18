# Stack Trace Grep

A simple and handy tool to grep Java Stack Traces (Exceptions) from log-files. Practically it's similar to grep, but not so powerful like grep ;)
The only purpose of this tool is to extract complete Stack Traces from files.

## Usage

Just execute the binary with two parameters. The text to search and the file to search in.

`sgrep search-text filename`

## Build

To build stacktrace-grep from source you'll have to install golang. Afterwards you can either use 

`make` 

or call "go build" directly

`go build -o bin/sgrep src/sgrep.go`

from the root-folder. You'll find the created binary in the bin-directory.

## Example

With the following call, you can grep all NullPointerExceptions with the complete Stack Trace from the sample file data.log 

`bin/sgrep NullPointerException sample/data.log`

Snippet from sample/data.log
```
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Faucibus in ornare quam viverra orci. Porttitor rhoncus dolor purus non enim praesent. Massa sapien faucibus et molestie. Metus vulputate eu scelerisque felis imperdiet proin fermentum leo. Id consectetur purus ut faucibus pulvinar elementum integer enim neque. Vestibulum morbi blandit cursus risus at ultrices mi tempus. Risus pretium quam vulputate dignissim suspendisse in est. Adipiscing commodo elit at imperdiet dui accumsan sit amet. In tellus integer feugiat scelerisque. Urna nunc id cursus metus aliquam eleifend mi in nulla. Montes nascetur ridiculus mus mauris vitae ultricies leo integer. Erat pellentesque adipiscing commodo elit at. Cras pulvinar mattis nunc sed blandit libero volutpat sed. Quam elementum pulvinar etiam non quam lacus suspendisse faucibus.
Exception in thread "main" java.lang.NullPointerException
	at Sample.ende(Sample.java:18)
	at Sample.bar(Sample.java:14)
	at Sample.foo(Sample.java:10)
	at Sample.main(Sample.java:3
Magna sit amet purus gravida quis blandit turpis cursus. A cras semper auctor neque. Sit amet consectetur adipiscing elit duis. Ornare suspendisse sed nisi lacus sed. Eu augue ut lectus arcu bibendum at varius vel pharetra. Arcu non odio euismod lacinia at quis risus sed. Fermentum dui faucibus in ornare quam. Quis auctor elit sed vulputate mi sit amet mauris. Interdum varius sit amet mattis vulputate enim. Leo duis ut diam quam nulla porttitor massa id. Velit ut tortor pretium viverra suspendisse potenti nullam. Iaculis urna id volutpat lacus laoreet. Fames ac turpis egestas integer eget aliquet nibh praesent.
Feugiat vivamus at augue eget arcu dictum varius. Sed viverra ipsum nunc aliquet bibendum enim. Est ante in nibh mauris cursus mattis molestie a. Ullamcorper malesuada proin libero nunc consequat interdum varius. Libero justo laoreet sit amet cursus. Amet porttitor eget dolor morbi non arcu risus. Fermentum odio eu feugiat pretium nibh ipsum. Ante metus dictum at tempor commodo. Enim sed faucibus turpis in eu. Vitae proin sagittis nisl rhoncus mattis rhoncus urna. Arcu ac tortor dignissim convallis. Blandit massa enim nec dui nunc mattis enim ut. Tempor commodo ullamcorper a lacus vestibulum sed. Praesent elementum facilisis leo vel fringilla est ullamcorper.
Sagittis purus sit amet volutpat consequat mauris. In dictum non consectetur a. Ultrices in iaculis nunc sed augue. Tristique sollicitudin nibh sit amet commodo nulla. Elit pellentesque habitant morbi tristique senectus et. Pellentesque pulvinar pellentesque habitant morbi tristique senectus et netus et. Vestibulum lectus mauris ultrices eros in cursus turpis. Nibh cras pulvinar mattis nunc. Velit sed ullamcorper morbi tincidunt ornare. Tellus molestie nunc non blandit massa enim nec dui. Mattis nunc sed blandit libero volutpat sed. Blandit volutpat maecenas volutpat blandit aliquam etiam erat. Eget duis at tellus at urna condimentum mattis pellentesque id. Consectetur adipiscing elit pellentesque habitant morbi tristique senectus. Aliquet lectus proin nibh nisl condimentum id venenatis. At urna condimentum mattis pellentesque. Mauris vitae ultricies leo integer malesuada.
Exception in thread "main" java.lang.IllegalStateException
	at Sample.ende(Sample.java:18)
	at Sample.bar(Sample.java:14)
	at Sample.foo(Sample.java:10)
	at Sample.main(Sample.java:3
Tortor condimentum lacinia quis vel eros donec. Leo vel orci porta non pulvinar neque laoreet suspendisse interdum. Lobortis scelerisque fermentum dui faucibus in. Orci phasellus egestas tellus rutrum. Vitae et leo duis ut. Id nibh tortor id aliquet lectus. Eleifend quam adipiscing vitae proin sagittis nisl rhoncus mattis rhoncus. Risus commodo viverra maecenas accumsan lacus vel facilisis volutpat est. Tortor dignissim convallis aenean et tortor. Eget felis eget nunc lobortis mattis aliquam faucibus purus. Quisque egestas diam in arcu cursus euismod quis viverra. Lacus vestibulum sed arcu non odio euismod lacinia.
Aliquam vestibulum morbi blandit cursus risus at ultrices. Porta non pulvinar neque laoreet suspendisse interdum consectetur. Eu facilisis sed odio morbi quis commodo. Egestas erat imperdiet sed euismod nisi porta lorem mollis. Ornare suspendisse sed nisi lacus sed viverra tellus. Ut venenatis tellus in metus vulputate eu scelerisque felis imperdiet. Maecenas volutpat blandit aliquam etiam erat velit scelerisque in dictum. Id aliquet lectus proin nibh nisl condimentum id venenatis. Felis donec et odio pellentesque diam volutpat. Arcu non odio euismod lacinia at quis risus sed. Urna id volutpat lacus laoreet non curabitur. Id aliquet lectus proin nibh. Viverra ipsum nunc aliquet bibendum enim. Vestibulum mattis ullamcorper velit sed ullamcorper. Ultrices sagittis orci a scelerisque purus semper. Et tortor consequat id porta nibh venenatis cras sed. Feugiat pretium nibh ipsum consequat nisl.
Id donec ultrices tincidunt arcu. Viverra maecenas accumsan lacus vel facilisis volutpat est velit. Lectus urna duis convallis convallis tellus. Nunc aliquet bibendum enim facilisis. At imperdiet dui accumsan sit amet nulla facilisi. Risus commodo viverra maecenas accumsan lacus. Porttitor rhoncus dolor purus non enim praesent. Ac tincidunt vitae semper quis lectus. Cursus eget nunc scelerisque viverra mauris in aliquam sem fringilla. Nibh tellus molestie nunc non. Neque vitae tempus quam pellentesque.
Tincidunt id aliquet risus feugiat. Vitae proin sagittis nisl rhoncus mattis. Cum sociis natoque penatibus et magnis dis parturient. Ut sem viverra aliquet eget sit. Ac tortor dignissim convallis aenean et tortor at risus. In nibh mauris cursus mattis molestie a. Id semper risus in hendrerit gravida rutrum quisque. Massa vitae tortor condimentum lacinia quis vel eros donec ac. Porta non pulvinar neque laoreet suspendisse. Nisi est sit amet facilisis. Lacus sed turpis tincidunt id aliquet risus feugiat. Montes nascetur ridiculus mus mauris. Eu tincidunt tortor aliquam nulla facilisi cras. Fringilla phasellus faucibus scelerisque eleifend.
Exception in thread "main" java.lang.NullPointerException
	at Sample.ende(Sample.java:18)
	at Sample.bar(Sample.java:14)
	at Sample.foo(Sample.java:10)
	at Sample.main(Sample.java:3
Caused by: java.lang.IllegalStateException
	at Sample.ende(Sample.java:18)
	at Sample.bar(Sample.java:14)
	at Sample.foo(Sample.java:10)
	at Sample.main(Sample.java:3
Mi sit amet mauris commodo quis imperdiet. Tincidunt nunc pulvinar sapien et. Felis eget nunc lobortis mattis aliquam faucibus purus in massa. Consequat mauris nunc congue nisi vitae suscipit tellus. Ultricies mi quis hendrerit dolor magna eget est. Massa enim nec dui nunc mattis enim ut tellus elementum. Lorem ipsum dolor sit amet consectetur adipiscing elit. Ultrices neque ornare aenean euismod elementum nisi quis eleifend. Sed felis eget velit aliquet sagittis id consectetur purus. At auctor urna nunc id cursus metus. Lorem dolor sed viverra ipsum. At augue eget arcu dictum varius duis. Sed ullamcorper morbi tincidunt ornare massa eget egestas purus viverra. Congue mauris rhoncus aenean vel elit. Senectus et netus et malesuada. Sed risus pretium quam vulputate dignissim.
Exception in thread "main" java.lang.IllegalStateException
	at Sample.ende(Sample.java:18)
	at Sample.bar(Sample.java:14)
	at Sample.foo(Sample.java:10)
	at Sample.main(Sample.java:3
At ultrices mi tempus imperdiet. Sit amet venenatis urna cursus eget. Cursus mattis molestie a iaculis at erat. Odio morbi quis commodo odio aenean sed. Vitae justo eget magna fermentum iaculis eu non. Vivamus at augue eget arcu dictum. Et tortor consequat id porta. Mi ipsum faucibus vitae aliquet nec ullamcorper sit amet. Donec massa sapien faucibus et molestie ac feugiat sed lectus. Et netus et malesuada fames ac. Libero id faucibus nisl tincidunt eget nullam non nisi est. Felis bibendum ut tristique et egestas quis. Feugiat nisl pretium fusce id velit ut tortor pretium. Ut sem viverra aliquet eget sit amet tellus cras. Pulvinar elementum integer enim neque. A cras semper auctor neque vitae.
Purus gravida quis blandit turpis cursus in hac. Amet nisl purus in mollis. Consequat semper viverra nam libero. Arcu ac tortor dignissim convallis. Sodales ut eu sem integer vitae justo eget magna. Vitae aliquet nec ullamcorper sit amet risus. Magna fringilla urna porttitor rhoncus. Sed velit dignissim sodales ut eu sem integer vitae. Iaculis urna id volutpat lacus laoreet non. Sem integer vitae justo eget. Condimentum id venenatis a condimentum. Nunc aliquet bibendum enim facilisis. Tellus mauris a diam maecenas sed enim.
Exception in thread "main" java.lang.NullPointerException
	at Sample.ende(Sample.java:18)
	at Sample.bar(Sample.java:14)
	at Sample.foo(Sample.java:10)
	at Sample.main(Sample.java:3
Caused by: java.lang.NullPointerException
	at Sample.ende(Sample.java:18)
	at Sample.bar(Sample.java:14)
	at Sample.foo(Sample.java:10)
	at Sample.main(Sample.java:3
Porttitor eget dolor morbi non arcu risus quis varius quam. Tortor posuere ac ut consequat semper viverra nam. Turpis cursus in hac habitasse. Augue ut lectus arcu bibendum. In cursus turpis massa tincidunt dui ut. Et netus et malesuada fames ac. Turpis nunc eget lorem dolor sed viverra ipsum. Enim neque volutpat ac tincidunt vitae semper quis lectus. Ac odio tempor orci dapibus ultrices in iaculis nunc. Diam quis enim lobortis scelerisque fermentum dui faucibus. Diam in arcu cursus euismod. Et magnis dis parturient montes nascetur.
Velit ut tortor pretium viverra suspendisse potenti nullam. Hendrerit gravida rutrum quisque non tellus. Viverra justo nec ultrices dui sapien eget. Suscipit tellus mauris a diam maecenas sed enim ut sem. Nulla posuere sollicitudin aliquam ultrices sagittis. Turpis egestas sed tempus urna et pharetra pharetra. Tristique risus nec feugiat in. Libero nunc consequat interdum varius sit amet mattis. Sodales ut eu sem integer. Nunc vel risus commodo viverra maecenas accumsan lacus vel facilisis. Amet consectetur adipiscing elit ut aliquam purus sit amet luctus. Enim sed faucibus turpis in eu mi bibendum neque egestas. Neque volutpat ac tincidunt vitae semper. Blandit libero volutpat sed cras ornare arcu dui. Donec ac odio tempor orci dapibus ultrices in. Luctus accumsan tortor posuere ac ut consequat semper viverra nam. Sed tempus urna et pharetra pharetra massa massa.
Placerat vestibulum lectus mauris ultrices. Nullam eget felis eget nunc lobortis mattis aliquam. Ullamcorper malesuada proin libero nunc consequat interdum. Est placerat in egestas erat imperdiet sed euismod. Arcu cursus euismod quis viverra. Neque vitae tempus quam pellentesque nec nam aliquam sem et. Purus in massa tempor nec. Rhoncus est pellentesque elit ullamcorper dignissim cras. Congue quisque egestas diam in arcu. Egestas erat imperdiet sed euismod nisi porta lorem mollis aliquam. Sed augue lacus viverra vitae congue. Quam vulputate dignissim suspendisse in. Dui sapien eget mi proin sed libero enim sed.
Sagittis nisl rhoncus mattis rhoncus urna neque viverra justo. Tempor id eu nisl nunc mi. Dignissim enim sit amet venenatis. Nibh cras pulvinar mattis nunc. Facilisis leo vel fringilla est ullamcorper eget nulla. Nibh nisl condimentum id venenatis a condimentum vitae sapien pellentesque. Hac habitasse platea dictumst quisque sagittis purus sit amet volutpat. Lorem ipsum dolor sit amet consectetur adipiscing. Lorem ipsum dolor sit amet consectetur adipiscing elit. Tempus iaculis urna id volutpat.
Quis auctor elit sed vulputate mi. Posuere lorem ipsum dolor sit amet consectetur adipiscing. Ac tortor dignissim convallis aenean et tortor at risus. Egestas fringilla phasellus faucibus scelerisque eleifend donec pretium vulputate. Cras adipiscing enim eu turpis. Ultricies leo integer malesuada nunc vel risus commodo viverra maecenas. Mauris a diam maecenas sed enim ut. Nibh sit amet commodo nulla facilisi nullam vehicula ipsum. Tincidunt ornare massa eget egestas purus viverra. Cras semper auctor neque vitae. Mauris commodo quis imperdiet massa tincidunt nunc pulvinar sapien et. Nunc sed velit dignissim sodales ut eu. Vitae turpis massa sed elementum tempus. Gravida arcu ac tortor dignissim convallis aenean et tortor at. Duis tristique sollicitudin nibh sit. Cursus risus at ultrices mi. Nulla pharetra diam sit amet nisl suscipit. Lectus proin nibh nisl condimentum id venenatis. Sed velit dignissim sodales ut eu sem integer vitae justo.
Ac turpis egestas integer eget. Pretium nibh ipsum consequat nisl vel pretium lectus quam id. Tempus urna et pharetra pharetra. Consectetur a erat nam at lectus urna duis convallis. Diam donec adipiscing tristique risus. Etiam non quam lacus suspendisse faucibus interdum posuere lorem. Quam nulla porttitor massa id neque. Sit amet nisl purus in mollis nunc. Cras pulvinar mattis nunc sed blandit. Volutpat blandit aliquam etiam erat velit scelerisque in dictum. Ut lectus arcu bibendum at varius vel pharetra vel. Vel elit scelerisque mauris pellentesque pulvinar pellentesque habitant morbi. Proin sagittis nisl rhoncus mattis rhoncus. Commodo sed egestas egestas fringilla phasellus faucibus scelerisque.
Vestibulum lorem sed risus ultricies tristique nulla aliquet enim. Nunc mattis enim ut tellus elementum sagittis. Id diam maecenas ultricies mi eget mauris pharetra et. Suscipit adipiscing bibendum est ultricies integer quis auctor. Nunc eget lorem dolor sed viverra ipsum nunc. Id diam vel quam elementum pulvinar. Quam quisque id diam vel quam elementum pulvinar. Neque laoreet suspendisse interdum consectetur libero id faucibus nisl. Elementum integer enim neque volutpat. Aliquet eget sit amet tellus cras adipiscing enim eu. Ut sem nulla pharetra diam sit amet. Sed enim ut sem viverra. Lectus urna duis convallis convallis tellus. Phasellus faucibus scelerisque eleifend donec. Eget nullam non nisi est sit amet facilisis magna. At urna condimentum mattis pellentesque id nibh tortor id aliquet. Nunc consequat interdum varius sit amet. Neque vitae tempus quam pellentesque nec nam aliquam sem. Eros in cursus turpis massa tincidunt dui.
In nisl nisi scelerisque eu ultrices vitae auctor. Tortor condimentum lacinia quis vel eros donec ac odio tempor. Mattis vulputate enim nulla aliquet porttitor lacus luctus. Lectus quam id leo in vitae turpis. Ipsum faucibus vitae aliquet nec ullamcorper sit amet risus. Eu feugiat pretium nibh ipsum consequat. Pharetra magna ac placerat vestibulum lectus. Hendrerit gravida rutrum quisque non tellus orci. Iaculis at erat pellentesque adipiscing commodo elit at imperdiet. Non quam lacus suspendisse faucibus interdum posuere lorem. Ornare aenean euismod elementum nisi. Auctor urna nunc id cursus metus aliquam eleifend. Fusce ut placerat orci nulla pellentesque dignissim enim sit amet. Habitant morbi tristique senectus et netus. Nam aliquam sem et tortor consequat id. Dui accumsan sit amet nulla. Ultricies mi quis hendrerit dolor magna eget est lorem. Arcu odio ut sem nulla pharetra diam. Faucibus turpis in eu mi bibendum neque.
Mus mauris vitae ultricies leo integer. Dolor sed viverra ipsum nunc aliquet bibendum enim facilisis. Tincidunt ornare massa eget egestas. Non odio euismod lacinia at quis risus sed vulputate. Mauris ultrices eros in cursus turpis massa tincidunt. Malesuada fames ac turpis egestas maecenas pharetra convallis. Nam at lectus urna duis. Ac orci phasellus egestas tellus rutrum tellus pellentesque eu tincidunt. Enim sit amet venenatis urna cursus eget nunc scelerisque. Lacus luctus accumsan tortor posuere ac ut consequat. Neque egestas congue quisque egestas diam in arcu. Eros in cursus turpis massa tincidunt dui ut ornare.
Non nisi est sit amet facilisis. Diam quis enim lobortis scelerisque fermentum dui faucibus in. In iaculis nunc sed augue lacus viverra vitae congue. Urna nunc id cursus metus. Duis ut diam quam nulla porttitor. Risus in hendrerit gravida rutrum quisque non tellus orci ac. Varius sit amet mattis vulputate enim nulla aliquet porttitor lacus. Quis risus sed vulputate odio ut enim. Elementum tempus egestas sed sed risus pretium quam. Vulputate ut pharetra sit amet aliquam id diam. Amet est placerat in egestas erat imperdiet sed euismod. Arcu ac tortor dignissim convallis aenean et tortor at.
```

will result in

```
Exception in thread "main" java.lang.NullPointerException
	at Sample.ende(Sample.java:18)
	at Sample.bar(Sample.java:14)
	at Sample.foo(Sample.java:10)
	at Sample.main(Sample.java:3
Exception in thread "main" java.lang.NullPointerException
	at Sample.ende(Sample.java:18)
	at Sample.bar(Sample.java:14)
	at Sample.foo(Sample.java:10)
	at Sample.main(Sample.java:3
Caused by: java.lang.IllegalStateException
	at Sample.ende(Sample.java:18)
	at Sample.bar(Sample.java:14)
	at Sample.foo(Sample.java:10)
	at Sample.main(Sample.java:3
Exception in thread "main" java.lang.NullPointerException
	at Sample.ende(Sample.java:18)
	at Sample.bar(Sample.java:14)
	at Sample.foo(Sample.java:10)
	at Sample.main(Sample.java:3
Caused by: java.lang.NullPointerException
	at Sample.ende(Sample.java:18)
	at Sample.bar(Sample.java:14)
	at Sample.foo(Sample.java:10)
	at Sample.main(Sample.java:3
```
