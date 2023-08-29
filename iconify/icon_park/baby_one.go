package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BabyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-width="4"><circle cx="24" cy="8" r="5" fill="#2F88FF" stroke-linejoin="round"/><path d="M5 28C5 28 22 7.75 43 28"/><path fill="#2F88FF" stroke-linejoin="round" d="M19 28V24.2105C19 24.2105 19 19 24 19C29 19 29 24.2105 29 24.2105V28V32C29 32 29 37 24 37C19 37 19 32 19 32V28Z"/><path stroke-linejoin="round" d="M29 32L37 37L31 44"/><path stroke-linejoin="round" d="M19 32L11 37L17 44"/></g>`),
		g.Group(children),
	)
}