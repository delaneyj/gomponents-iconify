package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmilingFaceFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12Zm5.293-1.293a1 1 0 0 0 1.414 0L9 10.414l.293.293a1 1 0 0 0 1.414-1.414l-1-1a1 1 0 0 0-1.414 0l-1 1a1 1 0 0 0 0 1.414Zm8 0a1 1 0 0 0 1.414-1.414l-1-1a1 1 0 0 0-1.414 0l-1 1a1 1 0 0 0 1.414 1.414l.293-.293l.293.293ZM9.4 13.5a1 1 0 0 0-1.731 1.002A4.998 4.998 0 0 0 12 17a4.998 4.998 0 0 0 4.33-2.5a1 1 0 1 0-1.73-1c-.52.899-1.49 1.5-2.6 1.5a2.998 2.998 0 0 1-2.6-1.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}