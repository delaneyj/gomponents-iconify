package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClothesGlovesTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSClothesGlovesTwo0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M34 22v-7m0 0v0c0-4.691 0-7.037-1.24-8.653a6 6 0 0 0-1.107-1.107C30.037 4 27.69 4 23 4h-1c-5.657 0-8.485 0-10.243 1.757C10 7.515 10 10.343 10 16v22h24v-6s7 0 7-6v-5c0-6-7-6-7-6Z"/><path fill="#fff" d="M9 38h26a3 3 0 1 1 0 6H9a3 3 0 1 1 0-6Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSClothesGlovesTwo0)"/>`),
		g.Group(children),
	)
}