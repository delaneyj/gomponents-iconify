package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CreativeCommonsNcLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2ZM7.094 5.68A8 8 0 0 1 18.32 16.905l-2.154-2.154A2.5 2.5 0 0 0 14 11h-4a.5.5 0 0 1 0-1.001h5.5V8H13V6h-2v2h-1c-.182 0-.359.02-.53.056L7.094 5.68ZM5.68 7.094L7.835 9.25A2.5 2.5 0 0 0 10 13h4a.5.5 0 0 1 0 1.001H8.5v2H11v2h2v-2h1c.182 0 .359-.02.53-.056l2.376 2.376A8 8 0 0 1 5.68 7.095Z"/>`),
		g.Group(children),
	)
}