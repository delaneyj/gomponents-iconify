package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Label(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M414 207q13 13 13 30.5T414 267L265 417q-13 12-30.5 12T205 417L13 225Q0 212 0 195V45q0-17 12.5-29.5T43 3h149q18 0 30 12zM74.5 109q13.5 0 23-9t9.5-22.5t-9.5-23t-23-9.5T52 54.5t-9 23t9 22.5t22.5 9z"/>`),
		g.Group(children),
	)
}