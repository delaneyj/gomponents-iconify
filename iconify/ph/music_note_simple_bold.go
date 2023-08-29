package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNoteSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m211.45 52.51l-80-24A12 12 0 0 0 116 40v100.22A52 52 0 1 0 140 184V56.13l64.55 19.36a12 12 0 1 0 6.9-23ZM88 212a28 28 0 1 1 28-28a28 28 0 0 1-28 28Z"/>`),
		g.Group(children),
	)
}