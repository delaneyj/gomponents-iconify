package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftAzureDevops(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m22 18l-5 4l-8-3v3l-4.19-5.75l12.91 1.05V6.34L22 5.65V18M4.81 16.25V8.96l12.91-2.62L10.6 2v2.84L3.97 6.76L2 9.38v5.69l2.81 1.18Z"/>`),
		g.Group(children),
	)
}