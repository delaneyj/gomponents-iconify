package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightUpCornerArrowOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M209.5 23.3L314.2 128L0 442.2L69.8 512L384 197.8l104.7 104.7L512 0z"/>`),
		g.Group(children),
	)
}