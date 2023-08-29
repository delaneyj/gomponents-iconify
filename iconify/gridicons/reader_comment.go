package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReaderComment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" stroke="currentColor" stroke-width="1.5" d="M12.528 14.556v-.75h-8.75c-.568 0-1.028-.46-1.028-1.028v-8c0-.568.46-1.028 1.028-1.028h12.444c.568 0 1.028.46 1.028 1.028v7.948c0 .905-.438 1.756-1.175 2.282l-3.547 2.534z"/>`),
		g.Group(children),
	)
}