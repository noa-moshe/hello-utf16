# hello-utf16: A vulnerable Go application

hello-utf16 is a hello world for utf-16 encoded text. send it your name in utf-16, and get
back a message in utf-8!

## Running / demo

Compile the application by running (this creates `hello-utf16`):
```sh
go build .
```

You can now run the program in a terminal by executing:
```sh
./hello-utf16
```

In another terminal you can watch its CPU usage by running `top` (or use the MacOS Activity Monitor)

Now, try running it. The file `world.txt` contains "world" in utf-16 encoding.

you can show the contents with (and see `...w.o.r.l.d` on the side):
```sh
hexdump -C world.txt
```

send it to hello-utf16 (this outputs `hello, world` in utf-8):
```sh
curl --data-binary "@world.txt" localhost:8080/hello
```
but uh-oh! there's a Denial of Service (DOS) in hello-utf16, introduced by a vulnerable dependency.

exploit the DOS using a specially crafted payload in the `dos.txt` file:
```sh
curl --data-binary "@dos.txt" localhost:8080/hello
```

This command will hang for a long time, so `control-c` cancel out of it. Meanwhile in `top`, `hello-utf16` is using 100% CPU. You could call the `curl` command again and keep using more cpu, but the attack is demonstrated!

`control-c` in the other terminal that's running `hello-utf16`, too.

## Detect and Fix

run Snyk and see the problem:

```sh
snyk test
```

We can see the vulnerability in `golang.org/x/text/encoding/unicode` and that it was fixed in version `v0.3.3`.

Upgrade to a fixed version using `go`:

```sh
go get golang.org/x/text/encoding/unicode@v0.3.3
```

(note the `v` in the version string. I often forget it, and the snyk test output doesn't show it)

Now we're fixed! Show the vulnerability is gone with Snyk:

```sh
snyk test
```

And demonstrate that it's no longer exploitable:

Rebuild the program:

```sh
go build .
```

Back in your original terminal, run it again:
```sh
./hello-utf16
```

Now re-execute the DOS attack:
```sh
curl --data-binary "@dos.txt" localhost:8080/hello
```

You'll get a response like `hello, �`. it completes, and doesn't cause a DOS. your other users are happy!

## Getting back to a vulnerable state

In case you want to finish up with `snyk monitor`, a fast way to return to a vulnerable state is to run:

```sh
go get golang.org/x/text/encoding/unicode@v0.3.2
```

or:
```sh
git reset --hard HEAD
```

Then `snyk monitor` will show the vulns in our web ui.

