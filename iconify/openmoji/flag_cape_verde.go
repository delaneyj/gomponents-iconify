package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagCapeVerde(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#1e50a0" d="M5 17h62v38H5z"/><path fill="#d22f27" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 41h62v6H5z"/><g fill="#f1b31c"><circle cx="25.1" cy="51" r="1.75"/><circle cx="25.1" cy="37" r="1.75"/><circle cx="29.215" cy="49.663" r="1.75"/><circle cx="20.985" cy="38.337" r="1.75"/><circle cx="29.215" cy="38.337" r="1.75"/><circle cx="20.985" cy="49.663" r="1.75"/><circle cx="31.757" cy="41.837" r="1.75"/><circle cx="18.443" cy="46.163" r="1.75"/><circle cx="31.757" cy="46.163" r="1.75"/><circle cx="18.443" cy="41.837" r="1.75"/></g><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}