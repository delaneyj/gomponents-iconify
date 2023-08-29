package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoShooting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-width="4"><path stroke="#000" stroke-linejoin="round" d="M15 12L18 6H30L33 12H15Z"/><path fill="#2F88FF" stroke="#000" stroke-linejoin="round" d="M41 12H7C5.34315 12 4 13.2536 4 14.8V39.2C4 40.7464 5.34315 42 7 42H41C42.6569 42 44 40.7464 44 39.2V14.8C44 13.2536 42.6569 12 41 12Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="M28 20L20 34"/><circle cx="24" cy="27" r="8" stroke="#fff" stroke-linecap="round"/></g>`),
		g.Group(children),
	)
}