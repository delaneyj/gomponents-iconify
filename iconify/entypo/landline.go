package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Landline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4.902.25C3.498-.027 2.115.875 1.833 2.273c-1.105 5.455-1.105 9.997 0 15.454C2.08 18.952 3.17 19.8 4.388 19.8c.17 0 .342-.016.515-.05c1.412-.279 2.329-1.638 2.046-3.036c-.978-4.832-.978-8.598 0-13.43C7.231 1.888 6.314.529 4.902.25zM17 2H8.436a4.08 4.08 0 0 1-.017 1.44c-.936 4.72-.936 8.398 0 13.12c.098.49.09.973.017 1.44H17c1.1 0 2-.9 2-2V4c0-1.1-.899-2-2-2zm-5 12.5a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 12 14.5zM17 9h-7V5h7v4z"/>`),
		g.Group(children),
	)
}