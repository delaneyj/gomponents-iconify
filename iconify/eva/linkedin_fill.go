package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkedinFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaLinkedinFill0"><g id="evaLinkedinFill1"><g id="evaLinkedinFill2" fill="currentColor"><path d="M15.15 8.4a5.83 5.83 0 0 0-5.85 5.82v5.88a.9.9 0 0 0 .9.9h2.1a.9.9 0 0 0 .9-.9v-5.88a1.94 1.94 0 0 1 2.15-1.93a2 2 0 0 1 1.75 2v5.81a.9.9 0 0 0 .9.9h2.1a.9.9 0 0 0 .9-.9v-5.88a5.83 5.83 0 0 0-5.85-5.82Z"/><rect width="4.5" height="11.7" x="3" y="9.3" rx=".9" ry=".9"/><circle cx="5.25" cy="5.25" r="2.25"/></g></g></g>`),
		g.Group(children),
	)
}