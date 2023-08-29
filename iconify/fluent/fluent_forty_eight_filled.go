package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FluentFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M25.17 4.163a1.5 1.5 0 0 0-1.34 0l-12 6a1.5 1.5 0 0 0-.83 1.342v23.456c0 .516.265.996.702 1.27l12 7.54A1.5 1.5 0 0 0 26 42.5V30.432l11.17-5.585a1.5 1.5 0 0 0 0-2.684l-9.316-4.658l9.317-4.658a1.5 1.5 0 0 0 0-2.684l-12-6Z"/>`),
		g.Group(children),
	)
}