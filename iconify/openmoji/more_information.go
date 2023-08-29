package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreInformation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#D0CFCE" d="M56 32.8H39.2V16h-6.4v16.8H16v6.4h16.8V56h6.4V39.2H56z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M56 32.8H39.2V16h-6.4v16.8H16v6.4h16.8V56h6.4V39.2H56z"/>`),
		g.Group(children),
	)
}