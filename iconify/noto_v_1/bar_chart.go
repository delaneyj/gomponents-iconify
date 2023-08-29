package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BarChart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<defs><path id="notoV1BarChart0" d="M37.4 13.08c0-1.72-1.4-3.13-3.12-3.13H22.65c-1.72 0-3.12 1.41-3.12 3.13v112.17H37.4V13.08z"/></defs><use fill="#ed6c30" href="#notoV1BarChart0"/><defs><path id="notoV1BarChart1" d="M73.09 50.56c0-1.72-1.4-3.13-3.12-3.13H58.35c-1.72 0-3.13 1.41-3.13 3.13v74.69h17.87V50.56z"/></defs><use fill="#006ca2" href="#notoV1BarChart1"/><defs><path id="notoV1BarChart2" d="M108.79 30.71c0-1.72-1.41-3.13-3.13-3.13H94.03c-1.71 0-3.12 1.41-3.12 3.13v94.53h17.87V30.71z"/></defs><use fill="#fcc21b" href="#notoV1BarChart2"/><path fill="#2f2f2f" d="M122.63 128H2.76C1.23 128 0 126.77 0 125.24V6.18a2.76 2.76 0 0 1 2.76-2.76a2.75 2.75 0 0 1 2.75 2.76v116.31h117.12a2.76 2.76 0 0 1 2.76 2.76c0 1.52-1.24 2.75-2.76 2.75z"/>`),
		g.Group(children),
	)
}