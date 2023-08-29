package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.506 24a1.019 1.019 0 1 1 0-2.038h6.873a1.019 1.019 0 1 1 0 2.038zm3.263-5.491a1.019 1.019 0 1 1 0-2.038h9.728a1.019 1.019 0 1 1 0 2.038zm-2.04-5.49a1.019 1.019 0 1 1 0-2.038h13.399a1.019 1.019 0 1 1 0 2.038zM3.058 7.525a1.019 1.019 0 1 1 0-2.038h15.438a1.019 1.019 0 1 1 0 2.038zM1.019 2.039a1.019 1.019 0 1 1 0-2.038h23.187a1.019 1.019 0 1 1 0 2.038z"/>`),
		g.Group(children),
	)
}