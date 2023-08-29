package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lpt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#000" fill-rule="nonzero"/><path fill="#FFF" d="M14.225 23.483h3.508v3.508h-3.508zm0-15.483h3.508v3.508h-3.508zm8.267 0H26v3.508h-3.508zM6 8h3.508v3.508H6zm12.358 7.742h3.508v3.508h-3.508zm-8.275 0h3.508v3.508h-3.508z"/></g>`),
		g.Group(children),
	)
}