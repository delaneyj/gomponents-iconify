package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeployRules(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m18 4l-6 6l1.41 1.41L17 7.83V20h2V7.83l3.59 3.58L24 10l-6-6zM8 18h7v2H8zm0 4h16v2H8zm0 4h16v2H8z"/>`),
		g.Group(children),
	)
}