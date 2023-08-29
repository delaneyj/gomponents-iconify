package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Power(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.25 7.25a.75.75 0 0 0 1.5 0V.75a.75.75 0 0 0-1.5 0v6.5Zm4.04 5.157A5.5 5.5 0 1 1 5 3.39a.75.75 0 1 0-.818-1.257a7 7 0 1 0 7.635 0A.75.75 0 0 0 11 3.39a5.5 5.5 0 0 1 .291 9.017Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}