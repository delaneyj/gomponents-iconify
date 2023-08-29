package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArchiveBoxLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m221.37 69.32l-16-32A6 6 0 0 0 200 34H56a6 6 0 0 0-5.37 3.32l-16 32A6.07 6.07 0 0 0 34 72v136a14 14 0 0 0 14 14h160a14 14 0 0 0 14-14V72a6.07 6.07 0 0 0-.63-2.68ZM59.71 46h136.58l10 20H49.71ZM208 210H48a2 2 0 0 1-2-2V78h164v130a2 2 0 0 1-2 2Zm-43.76-62.24a6 6 0 0 1 0 8.48l-32 32a6 6 0 0 1-8.48 0l-32-32a6 6 0 0 1 8.48-8.48L122 169.51V104a6 6 0 0 1 12 0v65.51l21.76-21.75a6 6 0 0 1 8.48 0Z"/>`),
		g.Group(children),
	)
}