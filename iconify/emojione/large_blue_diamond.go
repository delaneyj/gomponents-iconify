package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LargeBlueDiamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#0071b8" d="M2.018 32L32 2.019l29.981 29.98L32 61.982z"/>`),
		g.Group(children),
	)
}