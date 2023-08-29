package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1217 44v131q0 18-12.5 30.5T1174 218H43q-17 0-30-12.5T0 175V44q0-17 13-30T43 1h1131q18 0 30.5 13t12.5 30zm-43 305v609q0 17-13 30t-31 13H87q-18 0-31-13t-13-30V349q0-18 13-31t31-13h1043q18 0 31 13t13 31zM870 501q0-27-19.5-46T804 436H413q-27 0-46 19t-19 46t19 46t46 19h391q27 0 46.5-19t19.5-46z"/>`),
		g.Group(children),
	)
}