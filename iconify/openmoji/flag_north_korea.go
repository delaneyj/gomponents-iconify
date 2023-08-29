package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagNorthKorea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" d="M5 17h62v38H5z"/><circle cx="22.587" cy="36.013" r="8.5" fill="#fff"/><path fill="#d22f27" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round" d="m19.122 41.513l3.506-11l3.293 10.929l-8.834-6.684l11-.115l-8.965 6.87z"/><path fill="#1e50a0" stroke="#fff" stroke-miterlimit="10" d="M5 49h62v6H5zm0-32h62v6H5z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}