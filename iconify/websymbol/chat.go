package websymbol

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1000 1000"),
		g.Raw(`<path fill="currentColor" d="M1179 465q0 126-79 233.5T885 868t-296 62q-122 0-234-39l2 3L0 1001q44-59 70.5-126.5T102 769l4-38Q0 611 0 465q0-126 79-233T293.5 63T589 1t296 62t215 169t79 233zm-250 0q0-29-21-50t-51-21q-29 0-50 21t-21 50q0 30 21 51t50 21q30 0 51-21t21-51zm-250 0q0-29-21-50t-51-21q-29 0-50 21t-21 50q0 30 21 51t50 21q30 0 51-21t21-51zm-250 0q0-29-21-50t-51-21q-29 0-50 21t-21 50q0 30 21 51t50 21q30 0 51-21t21-51z"/>`),
		g.Group(children),
	)
}