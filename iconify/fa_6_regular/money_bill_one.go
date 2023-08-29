package fa_6_regular

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoneyBillOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M112 112c0 35.3-28.7 64-64 64v160c35.3 0 64 28.7 64 64h352c0-35.3 28.7-64 64-64V176c-35.3 0-64-28.7-64-64H112zM0 128c0-35.3 28.7-64 64-64h448c35.3 0 64 28.7 64 64v256c0 35.3-28.7 64-64 64H64c-35.3 0-64-28.7-64-64V128zm176 128a112 112 0 1 1 224 0a112 112 0 1 1-224 0zm80-48c0 8.8 7.2 16 16 16v64h-8c-8.8 0-16 7.2-16 16s7.2 16 16 16h48c8.8 0 16-7.2 16-16s-7.2-16-16-16h-8v-80c0-8.8-7.2-16-16-16h-16c-8.8 0-16 7.2-16 16z"/>`),
		g.Group(children),
	)
}