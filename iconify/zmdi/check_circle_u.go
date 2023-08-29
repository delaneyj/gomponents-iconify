package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckCircleU(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M213 3q88 0 151 62.5T427 216t-63 150.5T213 429T62.5 366.5T0 216T62.5 65.5T213 3zm107 341v-43H107v43h213zm-143-85l143-143l-30-30l-113 113l-40-41l-30 30z"/>`),
		g.Group(children),
	)
}