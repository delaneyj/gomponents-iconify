package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulbFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaBulbFill0"><g id="evaBulbFill1"><path id="evaBulbFill2" fill="currentColor" d="M12 7a5 5 0 0 0-3 9v4a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-4a5 5 0 0 0-3-9Zm0-1a1 1 0 0 0 1-1V3a1 1 0 0 0-2 0v2a1 1 0 0 0 1 1Zm9 5h-2a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2ZM5 11H3a1 1 0 0 0 0 2h2a1 1 0 0 0 0-2Zm2.66-4.58L6.22 5a1 1 0 0 0-1.39 1.47l1.44 1.39a1 1 0 0 0 .73.28a1 1 0 0 0 .72-.31a1 1 0 0 0-.06-1.41Zm11.53-1.37a1 1 0 0 0-1.41 0l-1.44 1.37a1 1 0 0 0 0 1.41a1 1 0 0 0 .72.31a1 1 0 0 0 .69-.28l1.44-1.39a1 1 0 0 0 0-1.42Z"/></g></g>`),
		g.Group(children),
	)
}