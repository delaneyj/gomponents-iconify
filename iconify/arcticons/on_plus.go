package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OnPlus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.879 33.465v5.622m-2.813-2.811h5.622M16.517 15.453v6.15m3.583-3.19a6.949 6.949 0 0 1 3.248 5.898h0c0 3.826-3.059 6.926-6.83 6.926h0c-3.772 0-6.83-3.1-6.83-6.926a6.949 6.949 0 0 1 3.27-5.912M25.833 31.26V19.996c0-2.073 3.353-2.749 4.19-.47l3.784 10.302c.837 2.278 4.19 1.603 4.19-.471V18.096"/>`),
		g.Group(children),
	)
}