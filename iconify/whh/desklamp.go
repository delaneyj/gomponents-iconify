package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Desklamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 576q-128-93-128-223q0-50 17-97l-18-18l-267 267l206 391h254q26 0 45 19t19 45.5t-19 45t-45 18.5H64q-26 0-45-18.5t-19-45T19 915t45-19h186L68 553l-13 13q-9 10-22.5 10t-23-9.5t-9.5-23T9 521l28-28l-32-61q-7-11-3.5-24t15-20t24-3.5T60 400l24 46l256-255q-24-36-20-80t35-75q36-36 87-36t87 36l46 46q48-17 97-17q131 0 224 128z"/>`),
		g.Group(children),
	)
}