package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Picasa(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M30.5 4.45c6 3.727 9.967 10.442 10 17.238c0 2.275-.094 3.812-.863 6.812H30.5V4.45zm8.24 26.05c-3.188 6-9.398 10.268-16.2 10.967c-3.489.22-7.04-.662-10.04-1.893V30.5h26.24zM.5 21.413c.088-7.625 4.462-14.354 10.951-17.69l7.655 6.837l-17.88 16.023C.74 24.817.5 23.093.5 21.413zm10 17.007c-4-2.385-6.71-5.84-8.3-9.656l8.3-7.48V38.42zm18-21.836L13.367 2.78c2.346-.787 4.651-1.273 6.947-1.281c2.906 0 6.186.598 8.186 1.793v13.29z"/>`),
		g.Group(children),
	)
}