package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListBullet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feListBullet0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feListBullet1" fill="currentColor"><path id="feListBullet2" d="M10 4h10a1 1 0 0 1 0 2H10a1 1 0 1 1 0-2Zm0 7h10a1 1 0 0 1 0 2H10a1 1 0 0 1 0-2Zm0 7h10a1 1 0 0 1 0 2H10a1 1 0 0 1 0-2ZM5 7a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0 7a2 2 0 1 1 0-4a2 2 0 0 1 0 4Zm0 7a2 2 0 1 1 0-4a2 2 0 0 1 0 4Z"/></g></g>`),
		g.Group(children),
	)
}