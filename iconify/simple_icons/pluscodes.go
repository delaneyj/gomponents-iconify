package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pluscodes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0a2.4 2.4 0 0 0-2.4 2.4A2.4 2.4 0 0 0 12 4.8a2.4 2.4 0 0 0 2.4-2.4A2.4 2.4 0 0 0 12 0zM9.543 9.543v4.914h4.914V9.543zM2.4 9.6A2.4 2.4 0 0 0 0 12a2.4 2.4 0 0 0 2.4 2.4A2.4 2.4 0 0 0 4.8 12a2.4 2.4 0 0 0-2.4-2.4zm19.2 0a2.4 2.4 0 0 0-2.4 2.4a2.4 2.4 0 0 0 2.4 2.4A2.4 2.4 0 0 0 24 12a2.4 2.4 0 0 0-2.4-2.4zM12 19.2a2.4 2.4 0 0 0-2.4 2.4A2.4 2.4 0 0 0 12 24a2.4 2.4 0 0 0 2.4-2.4a2.4 2.4 0 0 0-2.4-2.4z"/>`),
		g.Group(children),
	)
}