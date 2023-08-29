package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GridBigRound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16 19a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm-8 0a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm8-8a3 3 0 1 1 0-6a3 3 0 0 1 0 6Zm-8 0a3 3 0 1 1 0-6a3 3 0 0 1 0 6Z"/>`),
		g.Group(children),
	)
}