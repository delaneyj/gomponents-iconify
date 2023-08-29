package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refresh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.5 16h-4a.5.5 0 0 0 0 1h2.976a8.996 8.996 0 0 1-7.459 4a9.009 9.009 0 0 1-8.904-10.419a.5.5 0 0 0-.986-.162A9.847 9.847 0 0 0 2 12c.006 5.52 4.48 9.994 10 10a10.009 10.009 0 0 0 8-3.999v2.5a.5.5 0 1 0 1 0v-4.002a.5.5 0 0 0-.5-.499zM16.737 3.196C12.32.818 6.921 2.107 4 6.003V3.5a.5.5 0 0 0-1 0v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 0-1H4.524a8.996 8.996 0 0 1 7.459-4a9.009 9.009 0 0 1 8.903 10.421a.499.499 0 1 0 .987.16A9.847 9.847 0 0 0 22 12a10.008 10.008 0 0 0-5.263-8.804zM15 12a3 3 0 1 0-6 0a3 3 0 0 0 6 0zm-5 0a2 2 0 1 1 4 0a2 2 0 0 1-4 0z"/>`),
		g.Group(children),
	)
}