package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Help(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M213.5 3q88.5 0 151 62.5T427 216t-62.5 150.5t-151 62.5t-151-62.5T0 216T62.5 65.5T213.5 3zM235 365v-42h-43v42h43zm44-165q20-20 20-48q0-35-25-60t-60.5-25T153 92t-25 60h43q0-18 12.5-30.5t30-12.5t30 12.5T256 152t-13 30l-26 27q-25 25-25 60v11h43q0-22 5.5-34.5T260 220z"/>`),
		g.Group(children),
	)
}