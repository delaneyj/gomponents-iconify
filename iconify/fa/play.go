package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Play(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1384 831L56 1569q-23 13-39.5 3T0 1536V64q0-26 16.5-36T56 31l1328 738q23 13 23 31t-23 31z"/>`),
		g.Group(children),
	)
}