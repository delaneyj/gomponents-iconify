package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoteOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22.11 21.46l-1.27 1.27L19.1 21H5a2 2 0 0 1-2-2V4.9L1.11 3l1.28-1.27L20.7 20.04v.01l1.41 1.41M21 17.8L6.2 3H15l6 6v8.8M19.5 10L14 4.5V10h5.5Z"/>`),
		g.Group(children),
	)
}