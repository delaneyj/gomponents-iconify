package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagWesternSahara(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#5c9e31" d="M5 17h62v38H5z"/><path fill="#fff" d="M5 30h62v12H5z"/><path d="M5 17h62v13H5z"/><path fill="#d22f27" d="M26 36L5 55V17l21 19z"/><path fill="#d22f27" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round" d="M35.122 36a4.053 4.053 0 0 1 3.378-3.924a4.396 4.396 0 0 0-.81-.076a4.004 4.004 0 1 0 0 8a4.396 4.396 0 0 0 .81-.076A4.053 4.053 0 0 1 35.122 36Z"/><path fill="#d22f27" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round" d="m38.707 38l1.328-4l1.145 3.939L38 35.565l4-.098L38.707 38z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}