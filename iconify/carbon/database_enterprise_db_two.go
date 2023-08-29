package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatabaseEnterpriseDbTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M17.84 24V11.997h2.943v2h.115a3.312 3.312 0 0 1 3.357-2.277c2.46 0 3.84 1.702 3.84 4.691V24h-2.943v-7.29c0-1.701-.575-2.575-1.978-2.575c-1.218 0-2.391.644-2.391 1.932V24zM4.224 24V7.95h10.578v2.69H7.259v3.886h6.669v2.69H7.259v4.094h7.543V24z"/>`),
		g.Group(children),
	)
}