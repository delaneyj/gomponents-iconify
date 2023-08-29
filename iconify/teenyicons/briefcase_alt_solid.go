package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BriefcaseAltSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 2.5V3H1.5A1.5 1.5 0 0 0 0 4.5v9A1.5 1.5 0 0 0 1.5 15h12a1.5 1.5 0 0 0 1.5-1.5v-9A1.5 1.5 0 0 0 13.5 3H10v-.5a2.5 2.5 0 0 0-5 0ZM7.5 1A1.5 1.5 0 0 0 6 2.5V3h3v-.5A1.5 1.5 0 0 0 7.5 1ZM.66 7.367a10.083 10.083 0 0 0 13.68 0l-.68-.734a9.083 9.083 0 0 1-12.32 0l-.68.734Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}