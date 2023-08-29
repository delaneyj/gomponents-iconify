package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Uikit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23.594 4.391L18.12 7.709l6.313 3.594v9.438l-8.484 4.719l-8.344-4.719v-7.313L2.131 10.61v13.391l13.625 8l14.115-8V8.006zm-2.537-1.318L15.75 0l-5.531 3.422l5.375 2.958z"/>`),
		g.Group(children),
	)
}