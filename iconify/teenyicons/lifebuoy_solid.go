package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LifebuoySolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M5.404 4.697L2.562 1.855a7.501 7.501 0 0 1 9.876 0L9.596 4.697a3.502 3.502 0 0 0-4.192 0Zm4.899.707a3.502 3.502 0 0 1 0 4.192l2.842 2.842a7.501 7.501 0 0 0 0-9.876l-2.842 2.842Zm-.707 4.899a3.502 3.502 0 0 1-4.192 0l-2.842 2.842a7.501 7.501 0 0 0 9.876 0l-2.842-2.842ZM4.697 5.404a3.502 3.502 0 0 0 0 4.192l-2.842 2.842a7.501 7.501 0 0 1 0-9.876l2.842 2.842Z"/>`),
		g.Group(children),
	)
}