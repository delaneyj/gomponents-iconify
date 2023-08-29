package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nomad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M220.28 512L0 381.024V127.405L219.089 0L438 126.214v257.191L220.28 512m108.855-410.858l-74.1 43.374v87.547l-74.244-44.087l-71.39 41.79v182.983l74.1-45.442v-98.495l79.144 49.962L329.135 278V101.142z"/>`),
		g.Group(children),
	)
}