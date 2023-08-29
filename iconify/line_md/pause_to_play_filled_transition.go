package line_md

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PauseToPlayFilledTransition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 6L9 18L7 18L7 6z"><animate fill="freeze" attributeName="d" dur="0.4s" values="M9 18L7 18L7 6L9 6L9 18;M13 15L8 18L8 6L13 9L13 15"/><set attributeName="opacity" begin="0.4s" to="0"/></path><path d="M15 6L17 6L17 18L15 18L15 6"><animate fill="freeze" attributeName="d" dur="0.4s" values="M15 6L17 6L17 18L15 18L15 6;M13 9L18 12L18 12L13 15L13 9"/><set attributeName="opacity" begin="0.4s" to="0"/></path><path d="M8 6L18 12L8 18z" opacity="0"><set attributeName="opacity" begin="0.4s" to="1"/></path></g>`),
		g.Group(children),
	)
}