package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Suicide(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M241 16a15 15 0 0 0-15 15v60h-15a15 15 0 0 0-15 15v90a15 15 0 0 0 13.844 14.938C158.366 301.06 106 326.67 106 361c0 90 72.837 135 150 135s150-45 150-135c0-34.326-52.37-59.927-103.844-150.03A15 15 0 0 0 316 196v-90a15 15 0 0 0-15-15h-15V31a15 15 0 0 0-15-15h-30zm15 210c15 45 90 120 90 135c0 45-45.033 90-90 90c-45.033 0-90-45-90-90c0-15 75-90 90-135z"/>`),
		g.Group(children),
	)
}