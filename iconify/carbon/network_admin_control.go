package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkAdminControl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 20.415L25.586 19L23 21.587L20.414 19L19 20.415L21.586 23L19 25.586L20.414 27L23 24.414L25.586 27L27 25.586L24.414 23L27 20.415zM24 4a4.005 4.005 0 0 0-4 4a3.951 3.951 0 0 0 .567 2.019L10.019 20.567A3.952 3.952 0 0 0 8 20a4 4 0 1 0 4 4a3.951 3.951 0 0 0-.567-2.019l10.548-10.548A3.952 3.952 0 0 0 24 12a4 4 0 0 0 0-8zM8 26a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2zm16-16a2 2 0 1 1 2-2a2.002 2.002 0 0 1-2 2zM8 11.18L5.41 8.59L4 10l4 4l7-7l-1.41-1.41L8 11.18z"/>`),
		g.Group(children),
	)
}