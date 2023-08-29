package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Braille(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M8.5 15A2.5 2.5 0 0 0 6 12.5h-.5v8H6A2.5 2.5 0 0 1 8.5 23m0-6V5.5H9A2.5 2.5 0 0 1 11.5 8v2.5m6 12.5v-8a2.5 2.5 0 0 0-2.5-2.5h-.55m.05 2.5v-2a2.5 2.5 0 0 0-2.5-2.5h-.5m0 0V15M1.5 1v2m6-2v2m6-2v2m3-2v2m6-2v2m-21 2v2m3-2v2m12-2v2m3-2v2m3-2v2"/>`),
		g.Group(children),
	)
}