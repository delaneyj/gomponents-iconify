package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagCaribbeanNetherlands(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#1e50a0" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round" d="M5 17h62v38H5z"/><path fill="#fff" d="M5 17v38l62-38H5z"/><g fill="#d22f27" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round"><path d="m19.518 31l3.614 6.259h-7.228L19.518 31z"/><path d="m19.518 39.345l-3.614-6.259h7.228l-3.614 6.259z"/></g><circle cx="19.518" cy="35.173" r="6" fill="none" stroke="#000" stroke-miterlimit="10"/><path fill="#f1b31c" d="M5 17v16l25-16H5z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}