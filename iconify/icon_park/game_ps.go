package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GamePs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M44 28H28V44H44V28Z"/><path fill="#2F88FF" d="M13 4L22 20H4L13 4Z"/><path fill="#2F88FF" d="M36 20C40.4183 20 44 16.4183 44 12C44 7.58172 40.4183 4 36 4C31.5817 4 28 7.58172 28 12C28 16.4183 31.5817 20 36 20Z"/><path stroke-linecap="round" d="M4 28L20 44"/><path stroke-linecap="round" d="M20 28L4 44"/></g>`),
		g.Group(children),
	)
}