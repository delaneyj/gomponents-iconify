package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagAruba(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#61b2e4" d="M5 17h62v38H5z"/><path fill="#f1b31c" d="M5 47h62v2H5zm0-4h62v2H5z"/><path fill="#d22f27" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round" d="M12.707 23.293L16 24l-3.293.707L12 28l-.707-3.293L8 24l3.293-.707L12 20l.707 3.293z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}