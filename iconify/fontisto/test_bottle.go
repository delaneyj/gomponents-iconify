package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TestBottle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.269 4.702H2.222A2.223 2.223 0 0 0 0 6.924v14.854C0 23.005.995 23.999 2.222 24h11.047a2.222 2.222 0 0 0 2.222-2.222V6.923a2.227 2.227 0 0 0-2.222-2.222zm.09 14.025H2.125V9.981H13.36zM2.314 3.925H13.18a.529.529 0 0 0 .526-.525V.526A.53.53 0 0 0 13.181 0H2.315a.531.531 0 0 0-.526.525V3.4c0 .29.235.525.525.525h.002z"/>`),
		g.Group(children),
	)
}