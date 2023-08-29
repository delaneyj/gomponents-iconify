package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 .847V17a4 4 0 1 1-2-3.465V3.153L8 4.867V19a4 4 0 1 1-2-3.465V3.133L22 .847ZM6 19a2 2 0 1 0-4 0a2 2 0 0 0 4 0Zm14-2a2 2 0 1 0-4 0a2 2 0 0 0 4 0Z"/>`),
		g.Group(children),
	)
}