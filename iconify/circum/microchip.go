package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microchip(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 14.5h-2A1.5 1.5 0 0 1 9.5 13v-2A1.5 1.5 0 0 1 11 9.5h2a1.5 1.5 0 0 1 1.5 1.5v2a1.5 1.5 0 0 1-1.5 1.5Zm-2-4a.5.5 0 0 0-.5.5v2a.5.5 0 0 0 .5.5h2a.5.5 0 0 0 .5-.5v-2a.5.5 0 0 0-.5-.5Z"/><path fill="currentColor" d="M20.435 14.5h-1.93v-5h1.93a.5.5 0 0 0 0-1h-1.93V8a2.507 2.507 0 0 0-2.5-2.5h-.5V3.565a.508.508 0 0 0-.5-.5a.5.5 0 0 0-.5.5V5.5h-5V3.565a.508.508 0 0 0-.5-.5a.5.5 0 0 0-.5.5V5.5h-.5a2.5 2.5 0 0 0-2.5 2.5v.5h-1.94a.5.5 0 1 0 0 1h1.94v5h-1.94a.5.5 0 1 0 0 1h1.94v.5a2.5 2.5 0 0 0 2.5 2.5h.5v1.94a.5.5 0 0 0 1 0V18.5h5v1.94a.5.5 0 0 0 1 0V18.5h.5a2.507 2.507 0 0 0 2.5-2.5v-.5h1.93a.5.5 0 0 0 0-1Zm-2.93 1.5a1.5 1.5 0 0 1-1.5 1.5h-8a1.5 1.5 0 0 1-1.5-1.5V8a1.5 1.5 0 0 1 1.5-1.5h8a1.511 1.511 0 0 1 1.5 1.5Z"/>`),
		g.Group(children),
	)
}