package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeDownOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.967 5.69c-.17-1.359-1.776-1.992-2.828-1.115L8.798 8.192a.25.25 0 0 1-.16.058H5A1.75 1.75 0 0 0 3.25 10v4c0 .966.784 1.75 1.75 1.75h3.638a.25.25 0 0 1 .16.058l4.34 3.617c1.053.877 2.66.244 2.83-1.116c.523-4.19.523-8.428 0-12.618Zm-1.868.037a.233.233 0 0 1 .38.15c.508 4.066.508 8.18 0 12.246a.233.233 0 0 1-.38.15l-4.34-3.617a1.75 1.75 0 0 0-1.121-.406H5a.25.25 0 0 1-.25-.25v-4A.25.25 0 0 1 5 9.75h3.638c.41 0 .806-.144 1.12-.406L14.1 5.727Z" clip-rule="evenodd"/><path fill="currentColor" d="M19.032 7.948a.75.75 0 1 0-1.408.517c.405 1.101.626 2.291.626 3.535a10.25 10.25 0 0 1-.626 3.535a.75.75 0 0 0 1.408.517A11.667 11.667 0 0 0 19.75 12c0-1.423-.253-2.788-.718-4.052Z"/>`),
		g.Group(children),
	)
}