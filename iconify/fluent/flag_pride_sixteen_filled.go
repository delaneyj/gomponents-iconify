package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagPrideSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path fill="#745125" d="M3 3h10v1H3z"/><path fill="#E62C46" d="M3 4h10v1H3z"/><path fill="#F36D38" d="M3 5h10v1H3z"/><path fill="#FFD23E" d="M3 6h10v1H3z"/><path fill="#61BC51" d="M3 7h10v1H3z"/><path fill="#1793E8" d="M3 8h10v1H3z"/><path fill="#B73FBB" d="M3 9h10v1H3z"/><path d="M2.5 2h11a.5.5 0 0 1 .5.5v8a.5.5 0 0 1-.5.5H3v2.5a.5.5 0 0 1-1 0v-11a.5.5 0 0 1 .5-.5zM3 3v7h10V3H3z" fill="#212121"/></g>`),
		g.Group(children),
	)
}