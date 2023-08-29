package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KhanAcademy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M3.292 29.86S-1.64 18.073 7.203 9.281C15.219 1.292 27.334.052 29.432 0c0 0 3.256 10.677-4.864 22.776c-8.109 12.099-19.365 8.817-19.365 8.817s11.115-14.167 9.776-14.083c-.932.057-6.765 6.473-11.687 12.353z"/>`),
		g.Group(children),
	)
}