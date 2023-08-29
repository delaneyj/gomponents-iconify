package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flask(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M1431 1320q56 89 21.5 152.5T1312 1536H160q-106 0-140.5-63.5T41 1320l503-793V128h-64q-26 0-45-19t-19-45t19-45t45-19h512q26 0 45 19t19 45t-19 45t-45 19h-64v399zM652 595l-272 429h712L820 595l-20-31V128H672v436z"/>`),
		g.Group(children),
	)
}