package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ManualGear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M40 12V24H8"/><path d="M24 12V24V36"/><path d="M8 12V24V36"/><path fill="#2F88FF" d="M44 8C44 10.2091 42.2091 12 40 12C37.7909 12 36 10.2091 36 8C36 5.79086 37.7909 4 40 4C42.2091 4 44 5.79086 44 8Z"/><path fill="#2F88FF" d="M28 8C28 10.2091 26.2091 12 24 12C21.7909 12 20 10.2091 20 8C20 5.79086 21.7909 4 24 4C26.2091 4 28 5.79086 28 8Z"/><path fill="#2F88FF" d="M12 8C12 10.2091 10.2091 12 8 12C5.79086 12 4 10.2091 4 8C4 5.79086 5.79086 4 8 4C10.2091 4 12 5.79086 12 8Z"/><path fill="#2F88FF" d="M28 40C28 42.2091 26.2091 44 24 44C21.7909 44 20 42.2091 20 40C20 37.7909 21.7909 36 24 36C26.2091 36 28 37.7909 28 40Z"/><path fill="#2F88FF" d="M12 40C12 42.2091 10.2091 44 8 44C5.79086 44 4 42.2091 4 40C4 37.7909 5.79086 36 8 36C10.2091 36 12 37.7909 12 40Z"/><path fill="#2F88FF" d="M40 44C42.2091 44 44 42.2091 44 40C44 37.7909 42.2091 36 40 36C37.7909 36 36 37.7909 36 40C36 42.2091 37.7909 44 40 44Z"/></g>`),
		g.Group(children),
	)
}