package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Oracle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.957 18.912A6.953 6.953 0 0 1 1 11.962A6.963 6.963 0 0 1 7.957 5h8.087A6.961 6.961 0 0 1 23 11.962a6.952 6.952 0 0 1-6.956 6.95H7.957Zm7.907-2.453a4.497 4.497 0 0 0 4.503-4.497a4.507 4.507 0 0 0-4.503-4.508H8.136a4.507 4.507 0 0 0-4.503 4.508a4.498 4.498 0 0 0 4.503 4.497h7.728Z"/>`),
		g.Group(children),
	)
}