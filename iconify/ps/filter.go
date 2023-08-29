package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Filter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m169 399l43 32q10 9 25 9q6 0 20-4q23-12 23-39V246L419 73q17-21 4-45q-15-25-38-25H47Q18 3 9 26q-10 26 4 45l139 175v119q0 22 17 34zM47 45h338L237 229v168l-42-32V229z"/>`),
		g.Group(children),
	)
}