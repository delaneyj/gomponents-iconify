package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shortcake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#FFF" d="m58 17.057l-17.751 2.728l-1.327 1.208L7 26a2 2 0 0 0-2 2v12h1a5 5 0 0 0 5-4.999a5 5 0 1 1 10 0a5 5 0 0 0 10 0a5 5 0 1 1 10 0a5 5 0 0 0 10 0a5 5 0 1 1 10 0A5 5 0 0 0 66 40h1V28c0-4.711-3.26-8.663-7.647-9.722c-.434-.104-.88-1.124-1.334-1.17"/><circle cx="49.5" cy="18" r="5" fill="#ea5a47"/><path fill="#fcea2b" d="M67 39v20a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V39"/><path fill="#fcea2b" d="M67 40h-1a5 5 0 0 1-5-5a5 5 0 0 0-10 0a5 5 0 0 1-10 0a5 5 0 0 0-10 0a5 5 0 0 1-10 0a5 5 0 0 0-10 0a5 5 0 0 1-5 5H5"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M67 28v31a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V28a2 2 0 0 1 2-2l33.996-5.439m17.023-2.51C63.063 18.561 67 22.822 67 28"/><path d="M5 40h1a5 5 0 0 0 5-4.999a5 5 0 1 1 10 0a5 5 0 0 0 10 0a5 5 0 1 1 10 0a5 5 0 0 0 10 0a5 5 0 1 1 10 0A5 5 0 0 0 66 40h1M5 48h62M5 52h62"/><circle cx="49.5" cy="18" r="5"/><path d="M49.5 16c0-5.523 4.477-10 10-10"/></g>`),
		g.Group(children),
	)
}