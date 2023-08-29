package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForFlagSouthSudan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="#141414" d="M32 5H4a4 4 0 0 0-4 4v3h36V9a4 4 0 0 0-4-4z"/><path fill="#078930" d="M0 27a4 4 0 0 0 4 4h28a4 4 0 0 0 4-4v-3H0v3z"/><path fill="#DA121A" d="M0 12h36v12H0z"/><path d="M0 12h36v2H0zm0 10h36v2H0z" fill="#EEE"/><path fill="#0F47AF" d="M1.351 6.004H1.35C.522 6.737 0 7.808 0 9.267v18c0 .926.522 1.997 1.351 2.729L17.5 18L1.351 6.004z"/><path fill="#FCDD09" d="M8.249 17.917l1.777-2.446l-2.875.934l-1.776-2.445v3.023l-2.875.934l2.875.934v3.022l1.776-2.445l2.875.934z"/>`),
		g.Group(children),
	)
}