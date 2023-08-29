package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1408 128q0 261-106.5 461.5T1035 896q160 106 266.5 306.5T1408 1664h96q14 0 23 9t9 23v64q0 14-9 23t-23 9H32q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h96q0-261 106.5-461.5T501 896Q341 790 234.5 589.5T128 128H32q-14 0-23-9T0 96V32Q0 18 9 9t23-9h1472q14 0 23 9t9 23v64q0 14-9 23t-23 9h-96zm-128 0H256q0 206 85 384h854q85-178 85-384zm-57 1216q-54-141-145.5-241.5T883 960H653q-103 42-194.5 142.5T313 1344h910z"/>`),
		g.Group(children),
	)
}