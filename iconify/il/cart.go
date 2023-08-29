package il

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 750 850"),
		g.Raw(`<path fill="currentColor" d="M278 557q29 0 49 20t20 50t-20 49t-49 20t-49-20t-21-49t21-50t49-20zm324 0q29 0 50 20t20 50t-20 49t-50 20t-49-20t-20-49t20-50t49-20zM255 418v24h452q11 0 11 11v46q0 12-11 12H174q-12 0-12-12q0-81-21-159L72 79q-3-8-11-8H5q-5 0-5-5V6q0-5 5-5h129q5 0 5 5v35q0 7 7 7h561q11 0 11 11v186q0 15-9 27t-24 17z"/>`),
		g.Group(children),
	)
}