package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AndroidAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M0 317v-85h299v85q0 62-44 106t-106 44t-105.5-44T0 317zM237 69q29 21 45.5 52.5T299 189v22H0v-22q0-36 16.5-67.5T61 69L17 24L34 7l49 49q32-16 66-16t66 16l50-49l17 17zM85.5 168q8.5 0 15-6.5t6.5-15t-6.5-15t-15-6.5t-15 6.5t-6.5 15t6.5 15t15 6.5zm128 0q8.5 0 15-6.5t6.5-15t-6.5-15t-15-6.5t-15 6.5t-6.5 15t6.5 15t15 6.5z"/>`),
		g.Group(children),
	)
}