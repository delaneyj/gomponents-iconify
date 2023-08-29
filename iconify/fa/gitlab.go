package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gitlab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m104 642l792 1015l-868-630q-18-13-25-34.5T3 950l101-308zm462 0h660L896 1657zM368 30l198 612H104L302 30q8-23 33-23t33 23zm1320 612l101 308q7 21 0 42.5t-25 34.5l-868 630l792-1015zm0 0h-462l198-612q8-23 33-23t33 23z"/>`),
		g.Group(children),
	)
}