package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagPrideTwentyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path fill="#745125" d="M5 4.25h18V6H5z"/><path fill="#E62C46" d="M5 6h18v2H5z"/><path fill="#F36D38" d="M5 8h18v2H5z"/><path fill="#FFD23E" d="M5 10h18v1H5z"/><path fill="#61BC51" d="M5 11h18v2H5z"/><path fill="#1793E8" d="M5 13h18v2H5z"/><path fill="#B73FBB" d="M5 15h18v2H5z"/><path d="M4.75 3a.75.75 0 0 0-.75.75v20.5a.75.75 0 0 0 1.5 0V18h17.75a.75.75 0 0 0 .75-.75V3.75a.75.75 0 0 0-.75-.75H4.75zm.75 1.5h17v12h-17v-12z" fill="#212121"/></g>`),
		g.Group(children),
	)
}