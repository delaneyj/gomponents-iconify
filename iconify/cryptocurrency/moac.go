package cryptocurrency

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moac(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 32C7.163 32 0 24.837 0 16S7.163 0 16 0s16 7.163 16 16s-7.163 16-16 16zm-.208-16.208L9.423 9.423l-.138-.208L7.02 7v17.515h2.284V12.4l4.916 4.985l1.592 1.592l1.592-1.592L22.32 12.4v12.115h2.284V7L22.32 9.215l-6.529 6.577z"/>`),
		g.Group(children),
	)
}