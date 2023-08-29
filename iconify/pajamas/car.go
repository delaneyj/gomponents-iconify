package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Car(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 2.5H5a.5.5 0 0 0-.5.5v2h7V3a.5.5 0 0 0-.5-.5ZM1.75 4.25H3V3a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v1.25h1.25a.75.75 0 0 1 0 1.5h-.918A3.995 3.995 0 0 1 15 9v5a1 1 0 1 1-2 0v-1H3v1a1 1 0 1 1-2 0V9c0-1.339.658-2.524 1.668-3.25H1.75a.75.75 0 0 1 0-1.5ZM12 11.5h1.5V9A2.5 2.5 0 0 0 11 6.5H5A2.5 2.5 0 0 0 2.5 9v2.5H4V11a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}