package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileLockLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M120 178h-10v-6a26 26 0 0 0-52 0v6H48a6 6 0 0 0-6 6v40a6 6 0 0 0 6 6h72a6 6 0 0 0 6-6v-40a6 6 0 0 0-6-6Zm-50-6a14 14 0 0 1 28 0v6H70Zm44 46H54v-28h60Zm98.24-134.24l-56-56A6 6 0 0 0 152 26H56a14 14 0 0 0-14 14v88a6 6 0 0 0 12 0V40a2 2 0 0 1 2-2h90v50a6 6 0 0 0 6 6h50v122a2 2 0 0 1-2 2h-40a6 6 0 0 0 0 12h40a14 14 0 0 0 14-14V88a6 6 0 0 0-1.76-4.24ZM158 46.48L193.52 82H158Z"/>`),
		g.Group(children),
	)
}