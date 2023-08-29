package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Newspapers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M18.91 20c.276-1.823.59-4.637.59-8c0-6-1-10.25-1-10.25l-.05-.25H1.5v.25S2.5 6 2.5 12s-1 10.25-1 10.25v.25h19.95l.05-.25s1-4.25 1-10.25c0-3.07-.262-5.681-.517-7.5H19m-2.685 12h-11m-.332-12a54.458 54.458 0 0 1 .496 9h11a54.453 54.453 0 0 0-.496-9h-11Z"/>`),
		g.Group(children),
	)
}