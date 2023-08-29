package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Conventionalcommits(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0C5.382 0 0 5.382 0 12s5.382 12 12 12s12-5.382 12-12S18.618 0 12 0zm0 1.6c5.753 0 10.4 4.647 10.4 10.4S17.753 22.4 12 22.4S1.6 17.753 1.6 12S6.247 1.6 12 1.6z"/>`),
		g.Group(children),
	)
}