package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lollipop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 6.565h-.19a6 6 0 0 0-11.62 0H6a1.5 1.5 0 1 0 0 3h.19a5.992 5.992 0 0 0 5.31 4.48v7.39a.5.5 0 0 0 .5.5a.5.5 0 0 0 .5-.5v-7.39a6.013 6.013 0 0 0 5.31-4.48H18a1.5 1.5 0 1 0 0-3Zm-6-3.5a4.991 4.991 0 0 1 4.77 3.5H7.23a4.991 4.991 0 0 1 4.77-3.5Zm0 10a4.991 4.991 0 0 1-4.77-3.5h9.54a4.991 4.991 0 0 1-4.77 3.5Zm6-4.5H6a.5.5 0 0 1-.5-.5a.5.5 0 0 1 .5-.5h12a.5.5 0 0 1 .5.5a.508.508 0 0 1-.5.5Z"/>`),
		g.Group(children),
	)
}