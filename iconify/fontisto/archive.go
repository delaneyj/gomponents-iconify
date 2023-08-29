package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24.975 5.143H1.227A1.258 1.258 0 0 1 0 3.886v-.033v.002v-2.597C0 .574.546.018 1.225.001h23.75c.681.015 1.227.57 1.227 1.253v.034v-.002v2.569l.001.039a1.25 1.25 0 0 1-1.225 1.25h-.001zm.129 17.571V8.138a1.28 1.28 0 0 0-1.28-1.28H2.384a1.28 1.28 0 0 0-1.28 1.28V22.72c0 .707.573 1.28 1.28 1.28h21.434c.71 0 1.285-.576 1.286-1.286zm-9.214-9h-5.569a.646.646 0 0 1-.64-.64v-.435a.646.646 0 0 1 .64-.64h5.571a.646.646 0 0 1 .64.64v.43a.645.645 0 0 1-.642.645z"/>`),
		g.Group(children),
	)
}