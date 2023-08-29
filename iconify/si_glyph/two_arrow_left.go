package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwoArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.146 8.027L8.89 3.032H4.95L1.007 8.027l4.042 4.964h3.918l.01-.011l-3.831-4.953zm7.783 0l3.951-4.995h-3.939L8.976 8.027l4.064 4.964h3.918l.01-.011l-4.039-4.953z"/>`),
		g.Group(children),
	)
}