package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TapeDriveSolidBadged(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path id="clarityTapeDriveSolidBadged0" fill="currentColor" d="M12.21 23a5 5 0 1 0-5-5a5 5 0 0 0 5 5Zm0-7a2 2 0 1 1-2 2a2 2 0 0 1 2-2Zm11.58 7a5 5 0 1 0-5-5a5 5 0 0 0 5 5Zm0-7a2 2 0 1 1-2 2a2 2 0 0 1 2-2Z"/><g id="clarityTapeDriveSolidBadged1" fill="currentColor"><path d="M30 13.5V24H6V12h19.51a7.49 7.49 0 0 1-3-6H4a2 2 0 0 0-2 2v20a2 2 0 0 0 2 2h28a2 2 0 0 0 2-2V12.34a7.49 7.49 0 0 1-4 1.16Z"/><circle cx="30" cy="6" r="5"/></g>`),
		g.Group(children),
	)
}