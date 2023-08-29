package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageMedical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 6v20H6V6h20m0-2H6a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2Z"/><path fill="currentColor" d="M22 16v-2h-5v-2h3v-2h-3V8h-2v2h-3v2h3v2h-5v2h5v2h-5v2h5v2h-3v2h8v-2h-3v-2h5v-2h-5v-2h5z"/>`),
		g.Group(children),
	)
}