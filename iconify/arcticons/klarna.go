package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Klarna(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 7.914h7.769v32.172H4.5z"/><circle cx="39.584" cy="36.17" r="3.916" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.591 26.03A22.863 22.863 0 0 0 31.49 7.914h-8.162a18.051 18.051 0 0 1-11.06 16.507l12.797 15.665h9.007Z"/>`),
		g.Group(children),
	)
}