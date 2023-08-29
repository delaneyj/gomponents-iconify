package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ComponentSwitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 8.5a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7Zm4.243 0A5.477 5.477 0 0 1 13 12c0 1.33-.472 2.55-1.257 3.5H16.5a3.5 3.5 0 1 0 0-7h-4.757ZM2 12a5.5 5.5 0 0 1 5.5-5.5h9a5.5 5.5 0 1 1 0 11h-9A5.5 5.5 0 0 1 2 12Z"/>`),
		g.Group(children),
	)
}