package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LevelUpAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m17 4l-.719.688l-8.5 8.5l1.438 1.439L16 7.844V26H6v2h12V7.844l6.781 6.781l1.438-1.438l-8.5-8.5L17 4z"/>`),
		g.Group(children),
	)
}