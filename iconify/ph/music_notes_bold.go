package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicNotesBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M215.38 22.54a12 12 0 0 0-10.29-2.18l-128 32A12 12 0 0 0 68 64v103.35A40 40 0 1 0 92 204v-82.63l104-26v40A40 40 0 1 0 220 172V32a12 12 0 0 0-4.62-9.46ZM52 220a16 16 0 1 1 16-16a16 16 0 0 1-16 16ZM92 96.63V73.37l104-26v23.26ZM180 188a16 16 0 1 1 16-16a16 16 0 0 1-16 16Z"/>`),
		g.Group(children),
	)
}