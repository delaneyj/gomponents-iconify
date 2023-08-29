package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PodAutoscalerOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15 4a1.004 1.004 0 0 0-1-1h-4a1.004 1.004 0 0 0-1 1v1h6Zm-1.3 4l3.2 7H7.1l3.18-7ZM15 6H9L4 17h16Z"/><circle cx="12" cy="11" r="1.5" fill="currentColor"/><path fill="currentColor" d="M16 18v1H8v-1H6v1.003A1.998 1.998 0 0 0 8 21h8a1.998 1.998 0 0 0 2-1.997V18ZM5 11H3.5V8.5L0 12l3.5 3.5V13H5v-2zm19 1l-3.5-3.5V11H19v2h1.5v2.5L24 12z"/>`),
		g.Group(children),
	)
}