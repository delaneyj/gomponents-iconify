package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotebookAndPen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M4 6V42H30V6H4Z"/><path stroke="#fff" d="M12 42V6"/><path fill="#2F88FF" stroke="#000" d="M44 6H36V38L40 42L44 38V6Z"/><path stroke="#fff" d="M36 12H44"/><path stroke="#000" d="M30 6H4"/><path stroke="#000" d="M30 42H4"/><path stroke="#000" d="M36 6V22"/><path stroke="#000" d="M44 6V22"/></g>`),
		g.Group(children),
	)
}