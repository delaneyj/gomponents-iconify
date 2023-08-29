package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bing(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m0 0l4.795 1.686V18.56l6.753-3.9l-3.31-1.555l-2.09-5.2l10.64 3.738v5.435L4.795 24l-4.8-2.67V0z"/>`),
		g.Group(children),
	)
}