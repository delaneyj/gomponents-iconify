package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeleteTime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.5 1h9v3H22v2h-2.029l-.163 5.529l-1.999-.059L17.97 6H6.03l.441 15H11.5v2H4.529l-.5-17H2V4h5.5V1Zm2 3h5V3h-5v1ZM13 8v7h-2V8h2Zm5.5 7a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7ZM13 18.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0Zm6.5-2.248v1.834l1.414 1.414l-1.414 1.414l-2-2v-2.662h2Z"/>`),
		g.Group(children),
	)
}