package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PatreonLogoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M232 104a56 56 0 1 1-56-56a56 56 0 0 1 56 56ZM80 48H64a8 8 0 0 0-8 8v152a8 8 0 0 0 8 8h16a8 8 0 0 0 8-8V56a8 8 0 0 0-8-8Z" opacity=".2"/><path d="M176 40a64 64 0 1 0 64 64a64.07 64.07 0 0 0-64-64Zm0 112a48 48 0 1 1 48-48a48.05 48.05 0 0 1-48 48ZM80 40H64a16 16 0 0 0-16 16v152a16 16 0 0 0 16 16h16a16 16 0 0 0 16-16V56a16 16 0 0 0-16-16Zm0 168H64V56h16v152Z"/></g>`),
		g.Group(children),
	)
}