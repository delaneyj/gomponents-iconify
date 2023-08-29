package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Batteryfull(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M992 512h-96v128q0 53-37.5 90.5T768 768H128q-53 0-90.5-37.5T0 640V128q0-53 37.5-90.5T128 0h640q53 0 90.5 37.5T896 128v128h96q13 0 22.5 9.5t9.5 22.5v192q0 13-9.5 22.5T992 512z"/>`),
		g.Group(children),
	)
}