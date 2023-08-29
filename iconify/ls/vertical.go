package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vertical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M68 3h237c37 0 68 30 68 68v551c0 38-31 69-68 69H68c-37 0-68-31-68-69V71C0 33 31 3 68 3zm597 279h-52c3-35-7-69-34-96c-20-20-47-33-75-35v69l-100-96l100-92v68c41 2 79 19 110 50c37 36 54 85 51 132zM44 556h286V133H44v423zm368-239h237c37 0 68 31 68 69v236c0 38-31 69-68 69H388c10-12 18-28 22-44h178V361H412v-44zM186 650c20 0 37-18 37-38s-17-37-37-37s-36 17-36 37s16 38 36 38z"/>`),
		g.Group(children),
	)
}