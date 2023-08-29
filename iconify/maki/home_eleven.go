package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M10.002 4.75a.25.25 0 0 1-.25.25H1.25A.25.25 0 0 1 1 4.75c0-.07.028-.14.08-.19l4.238-3.454l.016-.016a.248.248 0 0 1 .35.016L8 2.996V1.5a.5.5 0 0 1 1 0v2.31l.92.75c.052.05.081.119.08.19zM2 9.752a.248.248 0 0 0 .246.25h2.755v-2h1v2h2.752a.248.248 0 0 0 .248-.248V6.001H2v3.75z" fill="currentColor"/>`),
		g.Group(children),
	)
}