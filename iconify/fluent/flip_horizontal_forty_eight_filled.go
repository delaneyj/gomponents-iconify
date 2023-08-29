package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipHorizontalFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M43.761 39.312A1.5 1.5 0 0 1 42.5 40h-15a1.5 1.5 0 0 1-1.5-1.5v-33a1.5 1.5 0 0 1 2.866-.62l15 33a1.5 1.5 0 0 1-.105 1.432ZM29 12.425V37h11.17L29 12.425ZM5 40a1 1 0 0 1-.905-1.425l16-34A1 1 0 0 1 22 5v34a1 1 0 0 1-1 1H5Z"/>`),
		g.Group(children),
	)
}