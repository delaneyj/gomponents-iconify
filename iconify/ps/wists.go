package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wists(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M2 0h123q2 43 11 109.5t11 95.5q5-31 13-109t10-96h124q1 43 5.5 105t5.5 99q14-65 34-204h123q-7 53-25 204.5T406 436H260q0-3-9-66.5T237 269q-4 18-15.5 77.5T204 436H58Q19 193 2 0z"/>`),
		g.Group(children),
	)
}