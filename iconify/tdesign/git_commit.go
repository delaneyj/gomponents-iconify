package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitCommit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 9a3 3 0 1 0 0 6a3 3 0 0 0 0-6Zm-4.9 2a5.002 5.002 0 0 1 9.8 0H23v2h-6.1a5.002 5.002 0 0 1-9.8 0H1v-2h6.1Z"/>`),
		g.Group(children),
	)
}