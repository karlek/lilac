WIP
---
This project is a *work in progress*. The implementation is *incomplete* and
subject to change. The documentation can be inaccurate.

lilac
=====

Lilac is a lightweight flac decoder and alsa player written in golang.

Currently only supports 16 bits encoded flac files. (It's really lightweight!)

The name lilac is a fusion of `lite` and `flac`. Lilac is the english word for the most rockstar flower in the world (_syren_ for all you swedish fans out there).

__Notice:__

Test before playing a flac file that it's signed 16 bits before hand. Otherwise the decoder will eat your speakers.

You can test with mpv.
```shell
$ mpv a.flac
...
AO: [pulse] 44100Hz stereo 2ch s16
...
```

Public domain
-------------
I hereby release this code into the [public domain](https://creativecommons.org/publicdomain/zero/1.0/).
