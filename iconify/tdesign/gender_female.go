package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GenderFemale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 4.006A5 5 0 0 1 12.252 14h-.504A5 5 0 0 1 12 4.006Zm1 11.93a7.002 7.002 0 0 0-1-13.93a7 7 0 0 0-1 13.93V17H8v2h3v3h2v-3h3v-2h-3v-1.065Z"/>`),
		g.Group(children),
	)
}