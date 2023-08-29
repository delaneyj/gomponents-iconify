package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M4 13.25A7.25 7.25 0 0 1 11.25 6h25.5A7.25 7.25 0 0 1 44 13.25v15.5A7.25 7.25 0 0 1 36.75 36H26.397l-10.85 7.658C14.057 44.71 12 43.644 12 41.82V36h-.75A7.25 7.25 0 0 1 4 28.75v-15.5Z"/>`),
		g.Group(children),
	)
}