package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.38 3.25H6.81c-1.09 0-2.02.78-2.21 1.86l-1.31 7.5c-.11.66.07 1.33.49 1.84c.43.51 1.05.8 1.72.8h4.03V18c0 1.52 1.23 2.75 2.83 2.75c.7 0 1.33-.42 1.61-1.07l2.54-5.93h1.88c1.31 0 2.37-1.06 2.37-2.37V5.61c0-1.3-1.06-2.36-2.37-2.36Zm-3.12 9.6l-2.67 6.25c-.04.09-.13.16-.32.16c-.69 0-1.24-.56-1.24-1.25v-4.25H5.5c-.23 0-.43-.09-.57-.26a.724.724 0 0 1-.16-.62l1.31-7.5c.06-.36.37-.62.74-.62h8.44v8.09Zm3.99-1.47c0 .48-.39.87-.87.87h-1.61v-7.5h1.61c.48 0 .87.39.87.86v5.77Z"/>`),
		g.Group(children),
	)
}