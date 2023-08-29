package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Extension(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 4a3 3 0 0 1 6 0h6v6a3 3 0 1 1 0 6v6h-6.403l-.122-.858a2.5 2.5 0 0 0-4.95 0L8.403 22H2v-6.403l.858-.122a2.5 2.5 0 0 0 0-4.95L2 10.403V4h6Zm3-1a1 1 0 0 0-1 1v2H4v2.756a4.501 4.501 0 0 1 0 8.488V20h2.756a4.501 4.501 0 0 1 8.488 0H18v-6h2a1 1 0 1 0 0-2h-2V6h-6V4a1 1 0 0 0-1-1Z"/>`),
		g.Group(children),
	)
}