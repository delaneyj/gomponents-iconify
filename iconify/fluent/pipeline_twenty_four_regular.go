package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PipelineTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2 6.25A2.25 2.25 0 0 1 6.486 6h11.028A2.25 2.25 0 0 1 22 6.25v10.5a2.25 2.25 0 0 1-4.486.25H6.486A2.25 2.25 0 0 1 2 16.75V6.25Zm3 0a.75.75 0 0 0-1.5 0v10.5a.75.75 0 0 0 1.5 0V6.25Zm1.5 9.25h11v-8h-11v8Zm14-9.25a.75.75 0 0 0-1.5 0v10.5a.75.75 0 0 0 1.5 0V6.25Z"/>`),
		g.Group(children),
	)
}