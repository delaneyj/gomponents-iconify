package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gravetwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 1024H0q0-60 90-108.5T325 845l109-378l-274 109q-13 0-22.5-9.5T128 544V224q0-13 9.5-22.5T160 192l267 107L320 32q0-13 9.5-22.5T352 0h320q13 0 22.5 9.5T704 32L597 299l267-107q13 0 22.5 9.5T896 224v320q0 13-9.5 22.5T864 576L590 467l109 378q145 22 235 70.5t90 108.5z"/>`),
		g.Group(children),
	)
}