package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NfcVariantOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m1.25 2.05l20.7 20.7L20.7 24l-2-2H4a2 2 0 0 1-2-2V5.3l-2-2l1.25-1.25M3.81 2H20c1.11 0 2 .89 2 2v16.19l-2-1.99V4H5.8L3.81 2M6 9.3l-2-2V20h12.7l-2-2H6V9.3m12 6.9l-2-2V8h-3v2.28c.6.34 1 .98 1 1.72v.19L11 9.2V8a2 2 0 0 1 2-2h5v10.2M8 16h4.7L8 11.3V16m2-8h-.2l-2-2H10v2Z"/>`),
		g.Group(children),
	)
}