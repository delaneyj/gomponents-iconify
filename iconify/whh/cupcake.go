package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cupcake(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 640q-139 0-257-34.5T68.5 512T0 384q0-80 89-144.5T322 146q7 74 61.5 124T512 320t128.5-50T702 146q144 29 233 93.5t89 144.5q0 69-68.5 128T769 605.5T512 640zm0-384q-53 0-90.5-37.5T384 128t37.5-90.5T512 0t90.5 37.5T640 128t-37.5 90.5T512 256zM220 979q-92-36-92-83L64 576q53 53 130 84zm39-298q56 14 125 20v316q-52-7-98-18zm317 22v319q-34 2-64 2t-64-2V703q25 1 64 1q40 0 64-1zm162 296q-46 11-98 18V701q69-6 125-20zm222-423l-64 320q0 47-92 83l26-319q77-31 130-84z"/>`),
		g.Group(children),
	)
}