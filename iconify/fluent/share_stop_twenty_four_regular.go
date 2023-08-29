package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareStopTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M19.75 4A2.25 2.25 0 0 1 22 6.25v11.5A2.25 2.25 0 0 1 19.75 20H4.25A2.25 2.25 0 0 1 2 17.75V6.25A2.25 2.25 0 0 1 4.25 4h15.5zm0 1.5H4.25a.75.75 0 0 0-.75.75v11.5c0 .414.336.75.75.75h15.5a.75.75 0 0 0 .75-.75V6.25a.75.75 0 0 0-.75-.75zM9.28 8.215l2.72 2.72l2.725-2.716a.75.75 0 0 1 1.06 1.062l-2.724 2.715l2.724 2.724a.75.75 0 1 1-1.06 1.06L12 13.057L9.28 15.78a.75.75 0 1 1-1.06-1.06l2.72-2.724l-2.72-2.72a.75.75 0 0 1 1.06-1.06z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}