package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowRightCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M0 7.5a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Zm8.5-2.707L11.207 7.5L8.5 10.207L7.793 9.5l1.5-1.5H4V7h5.293l-1.5-1.5l.707-.707Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}