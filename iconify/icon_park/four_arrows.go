package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourArrows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 24L19 19M24 8V24V8ZM24 24L29 19L24 24Z"/><path d="M24 24L19 29M24 40V24V40ZM24 24L29 29L24 24Z"/><path fill="#2F88FF" d="M20 12C22.2091 12 24 10.2091 24 8C24 5.79086 22.2091 4 20 4C17.7909 4 16 5.79086 16 8C16 10.2091 17.7909 12 20 12Z"/><path fill="#2F88FF" d="M8 32C10.2091 32 12 30.2091 12 28C12 25.7909 10.2091 24 8 24C5.79086 24 4 25.7909 4 28C4 30.2091 5.79086 32 8 32Z"/><path d="M8 24H40"/><path fill="#2F88FF" d="M40 24C42.2091 24 44 22.2091 44 20C44 17.7909 42.2091 16 40 16C37.7909 16 36 17.7909 36 20C36 22.2091 37.7909 24 40 24Z"/><path fill="#2F88FF" d="M28 44C30.2091 44 32 42.2091 32 40C32 37.7909 30.2091 36 28 36C25.7909 36 24 37.7909 24 40C24 42.2091 25.7909 44 28 44Z"/></g>`),
		g.Group(children),
	)
}