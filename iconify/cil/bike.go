package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bike(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M116 468A100 100 0 1 0 16 368a100.113 100.113 0 0 0 100 100Zm0-168a68 68 0 1 1-68 68a68.077 68.077 0 0 1 68-68Zm180 68a100 100 0 1 0 100-100a100.113 100.113 0 0 0-100 100Zm100-68a68 68 0 1 1-68 68a68.077 68.077 0 0 1 68-68Z"/><circle cx="317.912" cy="94.088" r="34.088" fill="currentColor"/><path fill="currentColor" d="M190.954 266.3L240 314.69V404h32v-92.655a24.154 24.154 0 0 0-7.144-17.084l-51.274-50.588l66.453-66.453l58.165 59.551A24.14 24.14 0 0 0 355.369 244H384v-32h-25.262l-92.487-94.688l-.475.464l-8.4-8.4l-112 112l45.254 45.254Z"/>`),
		g.Group(children),
	)
}