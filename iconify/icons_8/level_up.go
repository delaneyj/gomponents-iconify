package icons_8

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LevelUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m21 4.094l-.72.687l-6 6l1.44 1.44L20 7.936V25H5v2h17V7.937l4.28 4.282l1.44-1.44l-6-6l-.72-.686z"/>`),
		g.Group(children),
	)
}