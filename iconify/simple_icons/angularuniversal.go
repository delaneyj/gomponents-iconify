package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Angularuniversal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.6 11.28v1.44a.48.48 0 0 1-.48.48H8.88a.48.48 0 0 1-.48-.48v-1.44a.48.48 0 0 1 .48-.48h6.24a.48.48 0 0 1 .48.48zM12 15.6a1.2 1.2 0 1 0 0 2.4a1.2 1.2 0 0 0 0-2.4zm3.12-8.4H8.88a.48.48 0 0 0-.48.48v1.44c0 .265.215.48.48.48h6.24a.48.48 0 0 0 .48-.48V7.68a.48.48 0 0 0-.48-.48zm8.04-3.204l-1.716 14.736L11.976 24L2.52 18.732L.84 3.996L11.976 0L23.16 3.996zM16.8 6.24a1.44 1.44 0 0 0-1.44-1.44H8.64A1.44 1.44 0 0 0 7.2 6.24v11.52c0 .795.645 1.44 1.44 1.44h6.72a1.44 1.44 0 0 0 1.44-1.44V6.24z"/>`),
		g.Group(children),
	)
}