package eva

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OptionsTwoOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="evaOptions2Outline0"><g id="evaOptions2Outline1"><path id="evaOptions2Outline2" fill="currentColor" d="M19 9a3 3 0 0 0-2.82 2H3a1 1 0 0 0 0 2h13.18A3 3 0 1 0 19 9Zm0 4a1 1 0 1 1 1-1a1 1 0 0 1-1 1ZM3 7h1.18a3 3 0 0 0 5.64 0H21a1 1 0 0 0 0-2H9.82a3 3 0 0 0-5.64 0H3a1 1 0 0 0 0 2Zm4-2a1 1 0 1 1-1 1a1 1 0 0 1 1-1Zm14 12h-7.18a3 3 0 0 0-5.64 0H3a1 1 0 0 0 0 2h5.18a3 3 0 0 0 5.64 0H21a1 1 0 0 0 0-2Zm-10 2a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/></g></g>`),
		g.Group(children),
	)
}