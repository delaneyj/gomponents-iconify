package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flipboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 0h7.7v24H0zm8.5 8.5h7.8v7.8H8.5zm0-8.5H24v7.7H8.5z"/>`),
		g.Group(children),
	)
}