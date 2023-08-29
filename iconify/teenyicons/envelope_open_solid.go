package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EnvelopeOpenSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M6.756.35a1.5 1.5 0 0 1 1.488 0l6 3.428a1.5 1.5 0 0 1 .408.341L7.5 7.933L.348 4.12c.113-.135.25-.251.408-.341l6-3.429Z"/><path fill="currentColor" d="M0 5.067V13.5A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5V5.067l-7.5 4l-7.5-4Z"/>`),
		g.Group(children),
	)
}