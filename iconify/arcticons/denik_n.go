package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DenikN(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.826 42.498H5.5V5.5h10.13l14.72 14.835"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.174 5.502H42.5V42.5H32.37L17.65 27.665"/>`),
		g.Group(children),
	)
}