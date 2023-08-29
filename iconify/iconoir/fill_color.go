package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FillColor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m2.636 10.293l7.071-7.071l8.485 8.485l-7.07 7.071a2 2 0 0 1-2.83 0l-5.656-5.657a2 2 0 0 1 0-2.828Zm5.657-8.485l1.414 1.414"/><path stroke-miterlimit="1.5" d="M20 15s3 2.993 3 4.887c0 1.655-1.345 3-3 3c-1.656 0-2.988-1.345-3-3C17.01 17.992 20 15 20 15Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}