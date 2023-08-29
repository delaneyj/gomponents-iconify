package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Texas(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M156.3 41.88V222.6l-130.57 3c35.98 40.7 60.88 78.7 123.07 126c36.4-48.2 78.8-54.7 144.7 100.5l66.8 18c.7-49.4-15.2-97.8 126-155.3l-11.3-93l-3-60.7l-211.9-39.8V41.88z"/>`),
		g.Group(children),
	)
}