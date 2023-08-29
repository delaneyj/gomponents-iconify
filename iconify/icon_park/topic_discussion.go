package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TopicDiscussion(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M44 8H4V38H19L24 43L29 38H44V8Z"/><path stroke="#fff" d="M21 15L20 32"/><path stroke="#fff" d="M28 15L27 32"/><path stroke="#fff" d="M33 20L16 20"/><path stroke="#fff" d="M32 27L15 27"/></g>`),
		g.Group(children),
	)
}