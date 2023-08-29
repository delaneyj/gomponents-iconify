package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tzc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm1.7-16.6h3.5v-2.8H10.8v2.8h3.5v10.456a10.065 10.065 0 0 0 4-.122v-2.91a7.144 7.144 0 0 1-.6.174zm1.4 10.11C23.105 24.205 26 20.44 26 16c0-5.523-4.477-10-10-10S6 10.477 6 16c0 4.44 2.895 8.205 6.9 9.51V22.5a7.2 7.2 0 1 1 6.2 0z"/>`),
		g.Group(children),
	)
}