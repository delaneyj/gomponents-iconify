package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagAngola(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path d="M5 17h62v38H5z"/><path fill="#d22f27" d="M5 17h62v19H5z"/><path fill="#f1b31c" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" d="m36.996 29.035l1.672 5.144L34.291 31h5.41l-4.376 3.179l1.671-5.144z"/><path fill="none" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m44.804 41.074l.107-.163a9.996 9.996 0 0 0-6.6-15.239m-7.602 17.971a9.99 9.99 0 0 0 8.765 1.407q.465-.145.91-.333m-9.969-10.062L47 46"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}