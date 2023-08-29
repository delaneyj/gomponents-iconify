package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagSuriname(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" d="M5 17h62v38H5z"/><path fill="#5c9e31" d="M5 50h62v5H5z"/><path fill="#fff" d="M5 46h62v4H5z"/><path fill="#5c9e31" d="M5 17h62v5H5z"/><path fill="#fff" d="M5 22h62v4H5z"/><path fill="#fcea2b" stroke="#fcea2b" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m32.122 42l3.983-12l3.434 11.816L30 34.696l12-.296l-9.878 7.6z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}