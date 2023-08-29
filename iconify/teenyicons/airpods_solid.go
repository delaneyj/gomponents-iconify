package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirpodsSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 3.5A3.148 3.148 0 0 0 3.461.376l-.32.04a3.167 3.167 0 0 0-1.508.609L0 2.25v2.5l1.633 1.225c.441.33.96.54 1.508.609l.32.04c.182.023.362.03.539.022V15h3V3.5ZM4 4H2V3h2v1Zm4-.5A3.148 3.148 0 0 1 11.539.376l.32.04a3.167 3.167 0 0 1 1.508.609L15 2.25v2.5l-1.633 1.225c-.441.33-.96.54-1.508.609l-.32.04a3.18 3.18 0 0 1-.539.022V15H8V3.5Zm3 .5h2V3h-2v1Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}