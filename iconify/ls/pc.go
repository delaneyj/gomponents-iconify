package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M59 4h598c33 0 60 26 60 59v382c0 33-27 61-60 61H424v99h170c25 0 27 82 30 85H93s0-85 30-85h170v-99H59c-33 0-59-28-59-61V63C0 30 26 4 59 4zm0 441h598V63H59v382z"/>`),
		g.Group(children),
	)
}