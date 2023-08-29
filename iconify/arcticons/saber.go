package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Saber(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h21.79L42.5 29.29V7.5a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 29.29H29.29V42.5m-1.21-27.343c-7.104-2.86-10.606.644-10.92 3.689c-.568 5.506 10.19.902 11.123 5.88s-7.37 9.984-12.6 5.345"/>`),
		g.Group(children),
	)
}