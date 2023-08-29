package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PentagonFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="m10.205 2.6l-6.96 5.238A3 3 0 0 0 2.2 11.176l2.896 8.765A3 3 0 0 0 7.946 22h8.12a3 3 0 0 0 2.841-2.037l2.973-8.764a3 3 0 0 0-1.05-3.37l-7.033-5.237l-.091-.061l-.018-.01l-.106-.07a3 3 0 0 0-3.377.148z"/></g>`),
		g.Group(children),
	)
}