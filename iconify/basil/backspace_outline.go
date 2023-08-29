package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BackspaceOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10.47 8.47a.75.75 0 0 1 1.06 0L14 10.94l2.47-2.47a.75.75 0 1 1 1.06 1.06L15.06 12l2.47 2.47a.75.75 0 0 1-1.06 1.06L14 13.06l-2.47 2.47a.75.75 0 1 1-1.06-1.06L12.94 12l-2.47-2.47a.75.75 0 0 1 0-1.06Z"/><path fill="currentColor" fill-rule="evenodd" d="M8.406 4.95a3.75 3.75 0 0 0-2.95 1.434l-3.56 4.535a1.75 1.75 0 0 0 0 2.162l3.56 4.535a3.75 3.75 0 0 0 2.95 1.434h10.647a2.75 2.75 0 0 0 2.75-2.75V7.7a2.75 2.75 0 0 0-2.75-2.75H8.406Zm-1.77 2.36a2.25 2.25 0 0 1 1.77-.86h10.647c.69 0 1.25.56 1.25 1.25v8.6c0 .69-.56 1.25-1.25 1.25H8.406a2.25 2.25 0 0 1-1.77-.86l-3.561-4.536a.25.25 0 0 1 0-.309L6.636 7.31Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}