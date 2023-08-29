package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TransformThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 2.5V6h8.586L20 2.586L21.414 4L18 7.414V16h3.5v2H18v3.5h-2V18H6V8H2.5V6H6V2.5h2ZM8 8v6.586L14.586 8H8Zm8 1.414L9.414 16H16V9.414Z"/>`),
		g.Group(children),
	)
}