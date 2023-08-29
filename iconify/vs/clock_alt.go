package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="M768 416q26 0 45 19t19 45v192l266 207q22 20 22 49q0 26-19 45t-45 19t-44-18L733 758q-29-19-29-54V480q0-26 19-45t45-19zm0-416Q559 0 382.5 103T103 382.5T0 768t103 385.5T382.5 1433T768 1536t385.5-103t279.5-279.5T1536 768t-103-385.5T1153.5 103T768 0zm0 224q148 0 273 73t198 198t73 273t-73 273t-198 198t-273 73t-273.5-73t-198-198T224 768t72.5-273t198-198T768 224z"/>`),
		g.Group(children),
	)
}