package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Analytics(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18 18V2H2v16h16zM16 5H4V4h12v1zM7 7v3h3c0 1.66-1.34 3-3 3s-3-1.34-3-3s1.34-3 3-3zm1 2V7c1.1 0 2 .9 2 2H8zm8-1h-4V7h4v1zm0 3h-4V9h4v2zm0 2h-4v-1h4v1zm0 3H4v-1h12v1z"/>`),
		g.Group(children),
	)
}