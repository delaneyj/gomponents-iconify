package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ElasticCloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M30.161 25.948a15.947 15.947 0 0 1-12.427 5.974a15.947 15.947 0 0 1-12.427-5.974A15.922 15.922 0 0 1 17.734 20c5.01 0 9.521 2.349 12.427 5.974zM13.76 16c0-2.026 1.667-3.984 3.495-3.984h.557c5.214 0 9.469-2.349 12.385-5.974A16.051 16.051 0 0 0 17.733.068A15.916 15.916 0 0 0 5.306 6.042A15.744 15.744 0 0 0 1.801 16c0 2.589.599 5.016 1.719 7.172a19.079 19.079 0 0 1 10.464-5.932c-.12-.401-.24-.839-.24-1.24z"/>`),
		g.Group(children),
	)
}