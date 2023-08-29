package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsAdjustOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.354 8.75H4a.75.75 0 0 1 0-1.5h9.354a2.751 2.751 0 0 1 5.293 0H20a.75.75 0 0 1 0 1.5h-1.354a2.751 2.751 0 0 1-5.293 0ZM14.75 8a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0Zm-4.104 8.75H20a.75.75 0 0 0 0-1.5h-9.354a2.751 2.751 0 0 0-5.292 0H4a.75.75 0 0 0 0 1.5h1.354a2.751 2.751 0 0 0 5.292 0ZM6.75 16a1.25 1.25 0 1 1 2.5 0a1.25 1.25 0 0 1-2.5 0Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}