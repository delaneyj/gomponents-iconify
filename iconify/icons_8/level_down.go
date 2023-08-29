package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LevelDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v2h15v17.063l-4.28-4.282l-1.44 1.44l6 6l.72.686l.72-.687l6-6l-1.44-1.44L22 24.063V5H5z"/>`),
		g.Group(children),
	)
}