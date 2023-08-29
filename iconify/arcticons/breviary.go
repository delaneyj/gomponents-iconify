package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Breviary(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.34 42.68a21.48 21.48 0 1 1 6.29 2.37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.63 45.05L26 28.19l10.54 4l2.37-6.29l-10.55-4l3.9-10.36l-6.3-2.37L22 19.53l-10.54-4l-2.33 6.33l10.54 4l-6.33 16.82"/>`),
		g.Group(children),
	)
}