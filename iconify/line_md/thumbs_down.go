package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-dasharray="80" stroke-dashoffset="80" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 4H18L21 11V14H14L15 20L12 21L7 13H3V4H7V13"><animate fill="freeze" attributeName="stroke-dashoffset" dur="0.8s" values="80;0"/></path>`),
		g.Group(children),
	)
}