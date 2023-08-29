package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M33.998 19.271h-5.571v10.91h5.571c3.072 0 5.572-2.447 5.572-5.455c0-3.009-2.5-5.455-5.572-5.455"/><path fill="currentColor" d="M52 2H12C6.476 2 2 6.477 2 12v40c0 5.523 4.477 10 10 10h40c5.523 0 10-4.477 10-10V12c0-5.523-4.477-10-10-10zM33.998 37.453h-5.571V52H21V12h12.999C41.168 12 47 17.707 47 24.727c0 7.017-5.832 12.726-13.002 12.726z"/>`),
		g.Group(children),
	)
}