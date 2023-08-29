package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Soccer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="17" cy="28" r="2" fill="currentColor"/><path fill="currentColor" d="M8 20.586L13.586 15L15 16.414L9.414 22z"/><path fill="currentColor" d="M28 16.584L19.414 8H6v2h12.586l3 3L6 28.586L7.414 30L23 14.415L26.584 18L23 21.586L24.414 23L28 19.416a2.004 2.004 0 0 0 0-2.832zM24.5 9A3.5 3.5 0 1 1 28 5.5A3.504 3.504 0 0 1 24.5 9zm0-5A1.5 1.5 0 1 0 26 5.5A1.502 1.502 0 0 0 24.5 4z"/>`),
		g.Group(children),
	)
}