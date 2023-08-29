package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thermo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.715 36.785a5.715 5.715 0 0 1-11.43 0m.001-25.571a5.715 5.715 0 1 1 11.43 0m7.069 7.118a5.715 5.715 0 0 1 0 11.43m-25.571-.002a5.715 5.715 0 0 1 0-11.428M24 29.76a5.715 5.715 0 0 1-5.714-5.715m18.501 5.715H24m-5.714-18.547v12.834"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.214 29.76h7.072v7.025m11.429-.001V29.76m0-18.547v7.119h7.071m-25.572 0h7.072"/>`),
		g.Group(children),
	)
}