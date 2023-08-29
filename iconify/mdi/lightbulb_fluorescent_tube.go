package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightbulbFluorescentTube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m7.05 2.81l-.55.53l-1.39-1.42l-1.24 1.24l1.41 1.42l-.7.7l-1.42-1.41l-1.24 1.24L3.34 6.5l-.53.55l14.14 14.14l.55-.53l1.39 1.42l1.24-1.24l-1.41-1.41l.7-.71l1.42 1.41l1.24-1.23l-1.42-1.4l.53-.55L7.05 2.81Z"/>`),
		g.Group(children),
	)
}