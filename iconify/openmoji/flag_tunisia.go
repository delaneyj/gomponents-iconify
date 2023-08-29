package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagTunisia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d22f27" d="M5 17h62v38H5z"/><circle cx="36" cy="36" r="9" fill="#fff"/><path fill="#d22f27" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round" d="M38.023 40.714a4.714 4.714 0 1 1 2.159-8.91a6 6 0 1 0 0 8.392a4.753 4.753 0 0 1-2.16.518Z"/><path fill="#d22f27" stroke="#d22f27" stroke-linecap="round" stroke-linejoin="round" d="m38.388 38.71l-.15-5.267l3.028 4.138l-4.75-1.429l4.657-1.826l-2.785 4.384z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17h62v38H5z"/>`),
		g.Group(children),
	)
}