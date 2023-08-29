package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<rect width="10" height="10" x="3" y="3" fill="currentColor" rx="2"/>`),
		g.Group(children),
	)
}