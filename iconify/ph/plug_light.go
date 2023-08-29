package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlugLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M236.24 67.76a6 6 0 0 0-8.48 0L192 103.51L152.49 64l35.75-35.76a6 6 0 0 0-8.48-8.48L144 55.51l-27.76-27.75a6 6 0 1 0-8.48 8.48l7.75 7.76l-54.38 54.38a38 38 0 0 0 0 53.75l17.13 17.12l-50.5 50.51a6 6 0 1 0 8.48 8.48l50.51-50.5l17.13 17.13a38 38 0 0 0 53.74 0L212 140.49l7.76 7.75a6 6 0 0 0 8.48-8.48L200.49 112l35.75-35.76a6 6 0 0 0 0-8.48Zm-87.11 118.62a26 26 0 0 1-36.77 0l-42.74-42.74a26 26 0 0 1 0-36.77L124 52.49L203.51 132Z"/>`),
		g.Group(children),
	)
}