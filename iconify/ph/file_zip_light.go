package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileZipLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M184 146h-16a6 6 0 0 0-6 6v56a6 6 0 0 0 12 0v-10h10a26 26 0 0 0 0-52Zm0 40h-10v-28h10a14 14 0 0 1 0 28Zm-50-34v56a6 6 0 0 1-12 0v-56a6 6 0 0 1 12 0Zm-40 56a6 6 0 0 1-6 6H56a6 6 0 0 1-5.21-9l26.87-47H56a6 6 0 0 1 0-12h32a6 6 0 0 1 5.21 9l-26.87 47H88a6 6 0 0 1 6 6ZM212.24 83.76l-56-56A6 6 0 0 0 152 26H56a14 14 0 0 0-14 14v72a6 6 0 0 0 12 0V40a2 2 0 0 1 2-2h90v50a6 6 0 0 0 6 6h50v18a6 6 0 0 0 12 0V88a6 6 0 0 0-1.76-4.24ZM158 82V46.48L193.52 82Z"/>`),
		g.Group(children),
	)
}