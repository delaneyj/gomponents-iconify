package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NzOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><g id="flagNz1x10"><g id="flagNz1x11"><path d="M0-.3v.5l1-.5z"/><path d="M.2.3L0-.1l1-.2z"/></g><use href="#flagNz1x11" transform="scale(-1 1)"/><use href="#flagNz1x11" transform="rotate(72 0 0)"/><use href="#flagNz1x11" transform="rotate(-72 0 0)"/><use href="#flagNz1x11" transform="scale(-1 1) rotate(72)"/></g></defs><path fill="#00247d" fill-rule="evenodd" d="M0 0h512v512H0z"/><g transform="translate(-148.7 90.5) scale(.60566)"><use width="100%" height="100%" fill="#fff" href="#flagNz1x10" transform="matrix(45.4 0 0 45.4 900 120)"/><use width="100%" height="100%" fill="#cc142b" href="#flagNz1x10" transform="matrix(30 0 0 30 900 120)"/></g><g transform="rotate(82 418.7 105.1) scale(.60566)"><use width="100%" height="100%" fill="#fff" href="#flagNz1x10" transform="rotate(-82 519 -457.7) scale(40.4)"/><use width="100%" height="100%" fill="#cc142b" href="#flagNz1x10" transform="rotate(-82 519 -457.7) scale(25)"/></g><g transform="rotate(82 418.7 105.1) scale(.60566)"><use width="100%" height="100%" fill="#fff" href="#flagNz1x10" transform="rotate(-82 668.6 -327.7) scale(45.4)"/><use width="100%" height="100%" fill="#cc142b" href="#flagNz1x10" transform="rotate(-82 668.6 -327.7) scale(30)"/></g><g transform="translate(-148.7 90.5) scale(.60566)"><use width="100%" height="100%" fill="#fff" href="#flagNz1x10" transform="matrix(50.4 0 0 50.4 900 480)"/><use width="100%" height="100%" fill="#cc142b" href="#flagNz1x10" transform="matrix(35 0 0 35 900 480)"/></g><path fill="#012169" d="M0 0h256v256H0z"/><path fill="#fff" d="M256 0v32l-95 96l95 93.5V256h-33.5L127 162l-93 94H0v-34l93-93.5L0 37V0h31l96 94l93-94z"/><path fill="#c8102e" d="m92 162l5.5 17L21 256H0v-1.5zm62-6l27 4l75 73.5V256zM256 0l-96 98l-2-22l75-76zM0 .5L96.5 95L67 91L0 24.5z"/><path fill="#fff" d="M88 0v256h80V0zM0 88v80h256V88z"/><path fill="#c8102e" d="M0 104v48h256v-48zM104 0v256h48V0z"/>`),
		g.Group(children),
	)
}