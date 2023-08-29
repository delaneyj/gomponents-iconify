package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Logout(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="m912 236l276 266l-276 264V589H499V413h413V236zM746 748l106 107q-156 146-338 146q-217 0-365.5-143.5T0 499q0-135 68-250T251.5 67.5T502 1q184 0 349 148L746 255Q632 151 503 151q-149 0-251.5 104T149 509q0 140 105.5 241T502 851q131 0 244-103z"/>`),
		g.Group(children),
	)
}