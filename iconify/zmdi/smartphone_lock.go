package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphoneLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M320 5q18 0 30.5 12.5T363 48v384q0 18-12.5 30.5T320 475H107q-18 0-30.5-12.5T64 432v-64h43v43h213V69H107v43H64V48q0-18 12.5-30.5T107 5h213zM145 219q10 0 18 8t8 19v75q0 10-8.5 18t-19.5 8H26q-10 0-18-8.5T0 319v-75q0-9 8-17t18-8v-32q0-22 18-38t41-16t41.5 16t18.5 38v32zm-28 0v-32q0-13-9-20.5T85.5 159t-23 7.5T53 187v32h64z"/>`),
		g.Group(children),
	)
}