package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Copy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#9B9B9A" d="M23.825 9.208h32.128v47.765H23.825z"/><path fill="#FFF" d="M19.936 13.097h32.128v47.765H19.936z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M52.06 56.97h3.89V9.21H23.82v3.89"/><path d="M19.936 13.097h32.128v47.765H19.936zm11.235 23.924h9.658M36 41.85v-9.658"/></g>`),
		g.Group(children),
	)
}