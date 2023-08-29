package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M0 75.63V1200h159.302V650.952C606.706 528.936 893.764 1025.558 1200 718.693V75.63c-475.667 308.371-726.319-274.04-1200 0z"/>`),
		g.Group(children),
	)
}