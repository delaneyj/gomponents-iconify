package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeycapPound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M11.875 12.208h48v47.834h-48z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.125 11.958h48v48h-48zm21.328 15.658L30.94 44.634m9.87-17.018l-2.513 17.018M28.066 32.357l16.717-.017m-17.816 7.507l16.717-.017"/>`),
		g.Group(children),
	)
}