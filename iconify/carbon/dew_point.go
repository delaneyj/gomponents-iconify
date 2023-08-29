package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DewPoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 10a4 4 0 1 1 4-4a4.005 4.005 0 0 1-4 4zm0-6a2 2 0 1 0 2 2a2.002 2.002 0 0 0-2-2zm-9.5 26A5.496 5.496 0 0 1 9 24.52c0-3.443 4.344-21.014 4.53-21.76a1 1 0 0 1 1.94 0c.186.746 4.53 18.317 4.53 21.76A5.496 5.496 0 0 1 14.5 30zm0-22.756C13.04 13.38 11 22.409 11 24.52a3.5 3.5 0 0 0 7 0c0-2.111-2.04-11.14-3.5-17.276z"/>`),
		g.Group(children),
	)
}