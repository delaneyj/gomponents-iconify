package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphonePortraitLock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M107 325q-9 0-15.5-6T85 304v-64q0-9 6.5-15t15.5-6v-22q0-17 12.5-29.5t30-12.5t30 12.5T192 197v22q9 0 15 6t6 15v64q0 9-6 15t-15 6h-85zm17-128v22h51v-22q0-10-7.5-17.5t-18-7.5t-18 7.5T124 197zM256 5q18 0 30.5 12.5T299 48v384q0 18-12.5 30.5T256 475H43q-18 0-30.5-12.5T0 432V48q0-18 12.5-30.5T43 5h213zm0 384V91H43v298h213z"/>`),
		g.Group(children),
	)
}