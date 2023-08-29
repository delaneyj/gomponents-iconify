package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VisualStudio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.858 23.998l-9.771-9.484l-5.866 4.465L0 17.864V6.145l2.234-1.121l5.87 4.469L17.851 0l5.587 2.239V21.77L17.859 24zm-.563-16.186l-5.577 4.173l5.58 4.202zM2.788 9.497v5.016l2.787-2.525z"/>`),
		g.Group(children),
	)
}