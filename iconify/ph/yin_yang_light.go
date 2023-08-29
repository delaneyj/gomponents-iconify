package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YinYangLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 26a102 102 0 1 0 102 102A102.12 102.12 0 0 0 128 26ZM38 128a90.1 90.1 0 0 1 90-90a42 42 0 0 1 0 84a54 54 0 0 0-44.88 84A90.06 90.06 0 0 1 38 128Zm90 90a42 42 0 0 1 0-84a54 54 0 0 0 44.88-84A90 90 0 0 1 128 218Zm12-42a12 12 0 1 1-12-12a12 12 0 0 1 12 12Zm-22-96a10 10 0 1 1 10 10a10 10 0 0 1-10-10Z"/>`),
		g.Group(children),
	)
}