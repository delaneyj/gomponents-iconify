package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Van(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m254.07 114.79l-45.54-53.06A16 16 0 0 0 196.26 56H32a16 16 0 0 0-16 16v112a16 16 0 0 0 16 16h17a32 32 0 0 0 62 0h50a32 32 0 0 0 62 0h17a16 16 0 0 0 16-16v-64a8 8 0 0 0-1.93-5.21ZM230.59 112H176V72h20.26ZM104 112V72h56v40ZM88 72v40H32V72Zm-8 136a16 16 0 1 1 16-16a16 16 0 0 1-16 16Zm112 0a16 16 0 1 1 16-16a16 16 0 0 1-16 16Zm31-24a32 32 0 0 0-62 0h-50a32 32 0 0 0-62 0H32v-56h208v56Z"/>`),
		g.Group(children),
	)
}