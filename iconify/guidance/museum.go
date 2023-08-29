package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Museum(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M19.5 10.5V19m-10-8.5V19m-5-8.5V19m10-8.5V19M2 21h20M0 23.5h24m-.5-15.75v.75H.5v-.75C5 6 9.186 3.577 11.438.875L11.75.5h.5l.312.375C14.814 3.577 19 6 23.5 7.75Z"/>`),
		g.Group(children),
	)
}