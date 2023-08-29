package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobeFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="M11 4a5 5 0 1 1-4.995 5.217L6 9l.005-.217A5 5 0 0 1 11 4z"/><path fill="currentColor" d="M14.133 1.502a1 1 0 0 1 1.365-.369A9.015 9.015 0 1 1 5.094 15.755a1 1 0 1 1 1.312-1.51a7.015 7.015 0 1 0 8.096-11.378a1 1 0 0 1-.369-1.365z"/><path fill="currentColor" d="M11 16a1 1 0 0 1 .993.883L12 17v4a1 1 0 0 1-1.993.117L10 21v-4a1 1 0 0 1 1-1z"/><path fill="currentColor" d="M15 20a1 1 0 0 1 .117 1.993L15 22H7a1 1 0 0 1-.117-1.993L7 20h8z"/></g>`),
		g.Group(children),
	)
}