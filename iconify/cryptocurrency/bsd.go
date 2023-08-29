package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bsd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm12-16.873L19.927 4l-1.963 6.436L4 15.127h24zM4.11 16.655l7.854 11.236l1.963-6.655l13.964-4.581H4.109z"/>`),
		g.Group(children),
	)
}