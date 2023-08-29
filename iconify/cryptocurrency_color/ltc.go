package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ltc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#BFBBBB"/><path fill="#FFF" d="M10.427 19.214L9 19.768l.688-2.759l1.444-.58L13.213 8h5.129l-1.519 6.196l1.41-.571l-.68 2.75l-1.427.571l-.848 3.483H23L22.127 24H9.252z"/></g>`),
		g.Group(children),
	)
}