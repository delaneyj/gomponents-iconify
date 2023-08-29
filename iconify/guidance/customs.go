package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Customs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m18.5 18.5l5 5M8 7.5v8m6-8v8m-5.5-8a2.5 2.5 0 1 1 5 0m-2.5 14C5.201 21.5.5 16.799.5 11S5.201.5 11 .5S21.5 5.201 21.5 11S16.799 21.5 11 21.5Zm6-6H5v-.25l.128-.77a18.124 18.124 0 0 0 0-5.96L5 7.75V7.5h12v.25l-.128.77a18.126 18.126 0 0 0 0 5.96l.128.77v.25Z"/>`),
		g.Group(children),
	)
}