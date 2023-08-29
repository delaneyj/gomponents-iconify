package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlashlightSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.5 1a.5.5 0 0 1 .5.5v1a.5.5 0 1 1-1 0v-1a.5.5 0 0 1 .5-.5Zm2.646.146a.5.5 0 1 1 .707.708l-1 1a.5.5 0 1 1-.707-.708l1-1Zm-7.707.708a1.5 1.5 0 0 1 2.122 0l5.585 5.585a1.5 1.5 0 0 1 0 2.122l-.939.939L5.5 2.793l.94-.94ZM5 3.707v3.586l-3.646 3.646a1.5 1.5 0 0 0 0 2.122l1.586 1.585a1.5 1.5 0 0 0 2.122 0L8.707 11h3.586L5 3.707Zm1.855 6.147l-1 1a.5.5 0 0 1-.708-.708l1-1a.5.5 0 1 1 .708.708ZM14.5 5h-1a.5.5 0 0 1 0-1h1a.5.5 0 1 1 0 1Z"/>`),
		g.Group(children),
	)
}