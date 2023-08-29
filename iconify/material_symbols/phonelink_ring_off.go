package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhonelinkRingOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 23.3L.7 3.5l1.4-1.4l19.8 19.8l-1.4 1.4ZM5 6.425l2 2V18h9.6l2.4 2.4v.6q0 .825-.588 1.413T17 23H7q-.825 0-1.413-.588T5 21V6.425ZM8.825 6L5.15 2.3q.25-.575.738-.937T7 1h10q.825 0 1.413.588T19 3v4h-2V6H8.825Zm9.525 8.85L16.9 13.4q.3-.275.463-.637t.162-.763q0-.4-.162-.763T16.9 10.6l1.45-1.45q.575.575.875 1.313t.3 1.537q0 .8-.3 1.538t-.875 1.312Zm2.45 2.45l-1.4-1.4q.775-.775 1.2-1.775T21.025 12q0-1.125-.425-2.125T19.4 8.1l1.4-1.4q1.075 1.05 1.65 2.425T23.025 12q0 1.5-.575 2.875T20.8 17.3Z"/>`),
		g.Group(children),
	)
}