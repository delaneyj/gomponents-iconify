package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Apex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zM6 19.25v4.25l10-12.75L26 23.5v-4.25L16 6.5L6 19.25zm10.5 1.25a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5z"/>`),
		g.Group(children),
	)
}