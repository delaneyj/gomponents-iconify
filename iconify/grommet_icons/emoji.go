package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emoji(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 1.73A10.27 10.27 0 1 0 22.24 12A10.25 10.25 0 0 0 12 1.73ZM21 12a9 9 0 1 1-9-9a9 9 0 0 1 9 9Z"/><path fill="currentColor" d="M8.8 11.05a1.55 1.55 0 1 0-1.51-1.5a1.56 1.56 0 0 0 1.51 1.5Zm6.64 0a1.55 1.55 0 1 0 0-3.09a1.53 1.53 0 0 0-1.51 1.59a1.51 1.51 0 0 0 1.51 1.5Zm-3.25 5.3A6.58 6.58 0 0 1 6.9 13.5a5.71 5.71 0 0 0 5.3 4a5.54 5.54 0 0 0 5.31-4a6.27 6.27 0 0 1-5.32 2.85Z"/>`),
		g.Group(children),
	)
}