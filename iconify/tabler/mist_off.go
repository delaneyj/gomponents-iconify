package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MistOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5h9M3 10h7m8 0h1M5 15h5m4 0h1m4 0h2M3 20h9m4 0h3M3 3l18 18"/>`),
		g.Group(children),
	)
}