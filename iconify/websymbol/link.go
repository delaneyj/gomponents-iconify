package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Link(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1000 267q0 112-78 188L747 631q-78 78-189 78q-97 0-171-63l115-115q26 17 56 17q44 0 75-31l175-176q31-29 31-74q0-44-30.5-74.5T734 162q-25 0-49.5 12T652 208H414L546 79Q626 1 734 1q110 0 188 78t78 188zm-387 89L498 471q-26-17-56-17q-44 0-75 31L192 661q-31 29-31 74q0 44 30.5 74.5T266 840q25 0 49.5-12t32.5-34h238L454 923q-80 78-188 78q-110 0-188-78T0 735q0-112 78-188l175-176q78-78 189-78q97 0 171 63z"/>`),
		g.Group(children),
	)
}