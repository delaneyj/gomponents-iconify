package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayToPauseTransition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 15L8 18L8 6L13 9L13 9"><animate fill="freeze" attributeName="d" dur="0.6s" keyTimes="0;0.33;1" values="M13 15L8 18L8 6L13 9L13 9;M13 15L8 18L8 6L13 9L13 15;M9 18L7 18L7 6L9 6L9 18"/></path><path d="M13 9L18 12L18 12L13 15L13 15"><animate fill="freeze" attributeName="d" dur="0.6s" keyTimes="0;0.33;1" values="M13 9L18 12L18 12L13 15L13 15;M13 9L18 12L18 12L13 15L13 9;M15 6L17 6L17 18L15 18L15 6"/></path></g>`),
		g.Group(children),
	)
}