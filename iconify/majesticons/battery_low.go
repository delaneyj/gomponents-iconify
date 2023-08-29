package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryLow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 8a3 3 0 0 1 3-3h12a3 3 0 0 1 3 3v8a3 3 0 0 1-3 3H4a3 3 0 0 1-3-3V8zm19 2a1 1 0 0 1 1-1a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2a1 1 0 0 1-1-1v-4zM7 10a1 1 0 0 0-2 0v4a1 1 0 1 0 2 0v-4z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}