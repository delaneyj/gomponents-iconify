package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Houzz(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m512 1191l512-295v591l-512 296v-592zM0 896v591l512-296zM512 9v591L0 896V305zm0 591l512-295v591z"/>`),
		g.Group(children),
	)
}