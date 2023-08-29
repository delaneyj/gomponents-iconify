package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DataAreaTwentyFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M3 2.489A.5.5 0 0 0 2 2.5v15a.5.5 0 0 0 .5.5h15a.5.5 0 1 0 0-1H3V2.489ZM16 16V5.5a.5.5 0 0 0-.812-.39l-4.735 3.787l-3.205-1.831a.5.5 0 0 0-.45-.023L4 8.286V16h12Z"/>`),
		g.Group(children),
	)
}