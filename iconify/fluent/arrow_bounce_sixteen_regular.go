package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowBounceSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M8 4.5a.5.5 0 0 0-.5-.5h-6a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 1 0V5.707l6.397 6.397a.5.5 0 0 0 .707 0l5.75-5.75a.5.5 0 0 0-.708-.707L8.75 11.043L2.707 5H7.5a.5.5 0 0 0 .5-.5Z"/>`),
		g.Group(children),
	)
}