package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlaskOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 15a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-3 2a1 1 0 1 0 0 2a1 1 0 0 0 0-2Z"/><path fill="currentColor" fill-rule="evenodd" d="M14.493 3.25H8.5a.75.75 0 1 0 0 1.5h.26v5.087a7.25 7.25 0 0 1-1.256 4.078l-3.093 4.548a1.855 1.855 0 0 0 1.326 2.887c4.162.468 8.364.468 12.526 0a1.855 1.855 0 0 0 1.326-2.887l-3.093-4.548a7.25 7.25 0 0 1-1.256-4.078V4.75h.26a.75.75 0 0 0 0-1.5h-1.007ZM5.905 19.86c4.05.455 8.14.455 12.19 0a.355.355 0 0 0 .254-.553l-3.094-4.549a8.73 8.73 0 0 1-.591-1.008H9.336a8.745 8.745 0 0 1-.591 1.008L5.65 19.307a.355.355 0 0 0 .254.553Zm4.015-7.61h4.16a8.748 8.748 0 0 1-.34-2.413V4.75h-3.48v5.087a8.75 8.75 0 0 1-.34 2.413Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}