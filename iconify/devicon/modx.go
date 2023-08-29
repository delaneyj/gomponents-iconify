package devicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Modx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#00b5de" d="m63.864 14.059l-8.29 13.317l43.215 26.5l24.869-39.817z"/><path fill="#ff5529" d="m99.06 58.089l-27.178 42.806L111.97 125.9V66.106z" class="modx-original-st2"/><path fill="#00decc" d="m29.483 69.911l69.306-16.035L15.622 2.1v59.794z" class="modx-original-st3"/><path fill="#ff9640" d="M64.136 113.67L99.06 58.088L29.21 74.532L4.342 113.669z" class="modx-original-st4"/>`),
		g.Group(children),
	)
}