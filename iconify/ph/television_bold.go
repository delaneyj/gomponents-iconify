package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelevisionBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M216 60h-59l27.52-27.52a12 12 0 0 0-17-17L128 55L88.49 15.51a12 12 0 0 0-17 17L99 60H40a20 20 0 0 0-20 20v120a20 20 0 0 0 20 20h176a20 20 0 0 0 20-20V80a20 20 0 0 0-20-20ZM44 84h84v112H44Zm168 112h-60V84h60Zm-44-80a16 16 0 1 1 16 16a16 16 0 0 1-16-16Zm32 48a16 16 0 1 1-16-16a16 16 0 0 1 16 16Z"/>`),
		g.Group(children),
	)
}