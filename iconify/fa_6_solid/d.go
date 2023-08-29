package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func D(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 96c0-35.3 28.7-64 64-64h96c123.7 0 224 100.3 224 224S283.7 480 160 480H64c-35.3 0-64-28.7-64-64V96zm160 0H64v320h96c88.4 0 160-71.6 160-160S248.4 96 160 96z"/>`),
		g.Group(children),
	)
}