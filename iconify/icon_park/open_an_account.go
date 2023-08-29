package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OpenAnAccount(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M42 24V9C42 7.34315 40.6569 6 39 6H9C7.34315 6 6 7.34315 6 9V39C6 40.6569 7.34315 42 9 42H24"/><path d="M39.0508 33L39.0508 43"/><path d="M44 37.9497L34 37.9497"/><circle cx="24" cy="18" r="5" fill="#2F88FF"/><path d="M33 31C33 26.5817 28.9706 23 24 23C19.0294 23 15 26.5817 15 31"/></g>`),
		g.Group(children),
	)
}