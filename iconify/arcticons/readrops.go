package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Readrops(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="9" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="20.45" cy="27.55" r="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.06 20H43.5m-39 0h11.44M30.7 30h12.8m-39 0h12.79M4.5 39.6H24m9 0h10.5m0-31.2H24m-9 0H4.5"/>`),
		g.Group(children),
	)
}