package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZoomOut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M17 6.01C17 2.725 14.33.063 11.035.063c-3.294 0-5.964 2.662-5.964 5.947c0 3.284 2.67 5.948 5.964 5.948C14.33 11.958 17 9.294 17 6.01zM5.889 6c0-2.807 2.289-5.079 5.116-5.079c2.825 0 5.114 2.272 5.114 5.079c0 2.806-2.289 5.078-5.114 5.078c-2.827 0-5.116-2.272-5.116-5.078zm-2.98 9.938l-1.823-1.823l4.148-4.148s.088.773.571 1.256c.483.484 1.252.566 1.252.566l-4.148 4.149z"/><path d="M13.938 5.996c0 .523-.326.948-.729.948H8.771c-.402 0-.729-.425-.729-.948c0-.525.326-.95.729-.95h4.438c.403 0 .729.425.729.95z"/></g>`),
		g.Group(children),
	)
}