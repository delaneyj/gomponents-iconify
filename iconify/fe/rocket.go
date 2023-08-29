package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rocket(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feRocket0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feRocket1" fill="currentColor"><path id="feRocket2" d="M14 22h-4c-.8 0-1.602-1.123-2.274-2.726L5 22l-1-1v-4l2.383-2.383C6.14 13.325 6 12.057 6 11c0-4 3-9 6-9s6 5 6 9c0 1.058-.14 2.325-.383 3.617L20 17v4l-1 1l-2.726-2.726C15.602 20.877 14.801 22 14 22Zm-2-2h-1.615a3.136 3.136 0 0 1-.179-.249c-.347-.532-.72-1.365-1.059-2.383C8.455 15.29 8 12.755 8 11c0-3.198 2.444-7 4-7s4 3.802 4 7c0 1.755-.455 4.291-1.147 6.368c-.34 1.018-.712 1.85-1.06 2.383a3.136 3.136 0 0 1-.178.249H12Zm0-8a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z"/></g></g>`),
		g.Group(children),
	)
}