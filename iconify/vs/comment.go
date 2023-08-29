package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M1408 64q98 0 177 77t79 179v512q0 109-72.5 182.5T1408 1088H896l-512 512v-512H256q-113 0-184.5-72T0 832V320q0-109 73.5-182.5T256 64h1152z"/>`),
		g.Group(children),
	)
}