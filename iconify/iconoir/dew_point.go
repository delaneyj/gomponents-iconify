package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DewPoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5 12a5 5 0 1 0 6 0m-6 0V3h6v9m0-9h2m-2 3h2m-2 3h2"/><path d="M8 14a2 2 0 1 0 0 4a2 2 0 0 0 0-4Zm0 0V9"/><path stroke-miterlimit="1.5" d="M19 3s3 2.993 3 4.887c0 1.655-1.345 3-3 3c-1.656 0-2.988-1.345-3-3C16.01 5.992 19 3 19 3Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}