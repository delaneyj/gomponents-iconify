package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Webpack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m28.021 24.161l-11.552 6.505v-5.068l7.198-3.943zm.792-.713V9.839l-4.229 2.432v8.745zm-24.912.713l11.552 6.505v-5.068l-7.198-3.943zm-.792-.713V9.839l4.229 2.432v8.745zm.495-14.49l11.849-6.672v4.901l-7.646 4.188zm24.714 0L16.469 2.286v4.901l7.646 4.188zm-12.865 15.49l-7.099-3.891v-7.703l7.099 4.083zm1.016 0l7.099-3.891v-7.703l-7.099 4.083zM8.833 11.964l7.13-3.901l7.13 3.901l-7.13 4.099z"/>`),
		g.Group(children),
	)
}