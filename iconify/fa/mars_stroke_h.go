package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MarsStrokeH(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1901 659q19 19 19 45t-19 45l-294 294q-9 10-22.5 10t-22.5-10l-45-45q-10-9-10-22.5t10-22.5l185-185h-294v224q0 14-9 23t-23 9h-64q-14 0-23-9t-9-23V768h-132q-24 217-187.5 364.5T576 1280q-167 0-306-87T58 957T4 638q15-133 88-245.5t188-182T529 130q155-12 292 52.5t224 186T1148 640h132V416q0-14 9-23t23-9h64q14 0 23 9t9 23v224h294l-185-185q-10-9-10-22.5t10-22.5l45-45q9-10 22.5-10t22.5 10zM576 1152q185 0 316.5-131.5T1024 704T892.5 387.5T576 256T259.5 387.5T128 704t131.5 316.5T576 1152z"/>`),
		g.Group(children),
	)
}