package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareStopTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path d="M19.75 4A2.25 2.25 0 0 1 22 6.25v11.5A2.25 2.25 0 0 1 19.75 20H4.25A2.25 2.25 0 0 1 2 17.75V6.25A2.25 2.25 0 0 1 4.25 4zM8.22 8.215a.75.75 0 0 0 0 1.06l2.72 2.72l-2.72 2.725a.75.75 0 0 0 1.06 1.06L12 13.057l2.724 2.723a.75.75 0 1 0 1.06-1.06l-2.723-2.724l2.723-2.715a.75.75 0 1 0-1.06-1.062l-2.723 2.717l-2.72-2.72a.75.75 0 0 0-1.061 0z" fill="currentColor" fill-rule="nonzero"/>`),
		g.Group(children),
	)
}