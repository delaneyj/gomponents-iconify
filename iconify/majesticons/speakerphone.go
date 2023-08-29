package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Speakerphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M21 4.723c0-1.535-1.659-2.498-2.992-1.736L12.734 6H7a5 5 0 0 0-1 9.9v3.6a2.5 2.5 0 0 0 5 0V16h1.734l5.274 3.013c1.333.762 2.992-.2 2.992-1.736V4.723z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}