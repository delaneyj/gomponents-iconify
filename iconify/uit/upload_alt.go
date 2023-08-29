package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 9h-1.5a.5.5 0 0 0 0 1H18a2.003 2.003 0 0 1 2 2v7a2.003 2.003 0 0 1-2 2H6a2.003 2.003 0 0 1-2-2v-7a2.003 2.003 0 0 1 2-2h2.5a.5.5 0 0 0 0-1H6a3.003 3.003 0 0 0-3 3v7a3.003 3.003 0 0 0 3 3h12a3.003 3.003 0 0 0 3-3v-7a3.003 3.003 0 0 0-3-3zM8.862 6.345L11.5 3.707v13.794a.5.5 0 1 0 1 0V3.706l2.638 2.638a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.707l-3.491-3.491a.5.5 0 0 0-.344-.145L12 2l-.007.001a.498.498 0 0 0-.347.146l-3.49 3.49a.5.5 0 1 0 .707.707z"/>`),
		g.Group(children),
	)
}