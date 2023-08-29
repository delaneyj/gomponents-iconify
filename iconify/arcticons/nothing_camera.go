package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NothingCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="6.85" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="38.88" cy="16.06" r="1.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 33V15a3.46 3.46 0 0 0-3.42-3.5H8a3.46 3.46 0 0 0-3.5 3.42V33a3.46 3.46 0 0 0 3.42 3.5H40a3.46 3.46 0 0 0 3.5-3.42V33Z"/>`),
		g.Group(children),
	)
}