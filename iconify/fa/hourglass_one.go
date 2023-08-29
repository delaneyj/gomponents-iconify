package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HourglassOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1408 128q0 261-106.5 461.5T1035 896q160 106 266.5 306.5T1408 1664h96q14 0 23 9t9 23v64q0 14-9 23t-23 9H32q-14 0-23-9t-9-23v-64q0-14 9-23t23-9h96q0-261 106.5-461.5T501 896Q341 790 234.5 589.5T128 128H32q-14 0-23-9T0 96V32Q0 18 9 9t23-9h1472q14 0 23 9t9 23v64q0 14-9 23t-23 9h-96zm-128 0H256q0 66 9 128h1006q9-61 9-128zm0 1536q0-130-34-249.5t-90.5-208t-126.5-152T883 960H653q-76 31-146 94.5t-126.5 152t-90.5 208t-34 249.5h1024z"/>`),
		g.Group(children),
	)
}