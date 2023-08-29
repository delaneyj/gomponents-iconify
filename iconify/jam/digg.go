package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Digg(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.548.006v3.5H0v7h5.544V.006H3.548zm-1.552 8.75v-3.5h1.552v3.5H1.996zm3.992-5.25h1.996v7H5.988v-7zm0-3.5h1.996v1.969H5.988V.005zm8.205 3.5H8.871v7h3.326v1.531H8.871v1.969h5.322v-10.5zm-3.326 5.25v-3.5h1.33v3.5h-1.33zm9.092-5.25h-5.322v7h3.548v1.531h-3.548v1.969h5.322v-10.5zm-1.774 5.25h-1.552v-3.5h1.552v3.5z"/>`),
		g.Group(children),
	)
}