package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightningABold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M179.76 110.75a12 12 0 0 0-6.85-8.56L126 81.12l12.9-54.35a12 12 0 0 0-20.48-10.92L23.19 118.7a12 12 0 0 0 3.89 19.1l47 21.08l-12.96 54.34a12 12 0 0 0 20.49 10.93l95.19-102.86a12 12 0 0 0 2.96-10.54ZM95 174.31l4.64-19.54a12 12 0 0 0-6.76-13.72l-40.76-18.3L105 65.69l-4.64 19.54A12 12 0 0 0 107.08 99l40.77 18.3Zm147.7 36.32l-36-72a12 12 0 0 0-21.47 0l-36 72a12 12 0 1 0 21.46 10.73l6.72-13.36h37.17l6.68 13.36a12 12 0 1 0 21.47-10.73ZM189.41 184l6.59-13.17l6.58 13.17Z"/>`),
		g.Group(children),
	)
}