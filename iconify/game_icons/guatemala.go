package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Guatemala(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m190.8 25.24l175.1 1.08l-.1 215.08l27.3-.9l78 19.9l-97.3 76.3l-20.5 68.7l-84.7 81.4c-106.1-3.9-179.25-36.3-227.72-90l31.21-136.1l44.61-40.2l117.8-2.1c-.5-50.9-50.7-91.8-109.5-130.92l47 3.81z"/>`),
		g.Group(children),
	)
}