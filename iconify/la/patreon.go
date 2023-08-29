package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Patreon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 4v24h7V4H4zm17.5 0C16.813 4 13 7.813 13 12.5c0 4.687 3.813 8.5 8.5 8.5c4.687 0 8.5-3.813 8.5-8.5C30 7.813 26.187 4 21.5 4zM6 6h3v20H6V6zm15.5 0c3.584 0 6.5 2.916 6.5 6.5S25.084 19 21.5 19a6.508 6.508 0 0 1-6.5-6.5C15 8.916 17.916 6 21.5 6z"/>`),
		g.Group(children),
	)
}