package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HtmlFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="m1130 469l16-175H262l47 534h612l-22 228l-197 53l-196-53l-13-140H318l22 278l362 100h4v-1l359-99l50-544H471l-15-181h674zM0 0h1408l-128 1438l-578 162l-574-162z"/>`),
		g.Group(children),
	)
}