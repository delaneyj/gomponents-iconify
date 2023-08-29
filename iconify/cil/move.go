package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Move(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m403.814 165.089l-22.627 22.627l51.882 51.882H272.402V78.932l51.882 51.881l22.628-22.626l-90.51-90.511l-90.51 90.511l22.628 22.626l51.882-51.881v160.666H78.932l51.882-51.882l-22.627-22.627l-90.51 90.509l90.509 90.509l22.628-22.627l-51.883-51.882h161.471v161.47l-51.882-51.881l-22.628 22.626l90.51 90.511l90.51-90.511l-22.628-22.626l-51.882 51.881v-161.47h160.667l-51.883 51.882l22.628 22.627l90.509-90.509l-90.509-90.509z"/>`),
		g.Group(children),
	)
}