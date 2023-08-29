package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNoteSimpleLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m209.72 58.25l-80-24A6 6 0 0 0 122 40v113.05A46 46 0 1 0 134 184V48.06l72.27 21.69a6 6 0 1 0 3.45-11.5ZM88 218a34 34 0 1 1 34-34a34 34 0 0 1-34 34Z"/>`),
		g.Group(children),
	)
}