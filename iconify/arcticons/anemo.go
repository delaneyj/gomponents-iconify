package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anemo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.28 36.055a4.5 4.5 0 1 0 4.217-6.056H7.5m10.78-18.054a4.5 4.5 0 1 1 4.218 6.056H4.5m30.277-.057A4.5 4.5 0 1 1 38.996 24H5.999"/>`),
		g.Group(children),
	)
}