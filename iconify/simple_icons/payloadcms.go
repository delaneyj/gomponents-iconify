package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Payloadcms(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.068 0L22.08 6.625v12.573L13.787 24V11.427L2.769 4.808L11.068 0ZM1.92 18.302l8.31-4.812v9.812l-8.31-5Z"/>`),
		g.Group(children),
	)
}