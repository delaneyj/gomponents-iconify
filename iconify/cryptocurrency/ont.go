package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ont(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm9-7.783V15.61C25 10.854 21.06 7 16.2 7c-2.43 0-4.63.964-6.223 2.521zM7 7.783v8.608C7 21.146 10.94 25 15.8 25c2.43 0 4.63-.964 6.223-2.521z"/>`),
		g.Group(children),
	)
}