package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Equalizer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSEqualizer0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path stroke-linecap="round" d="M11 16v26m13-13v13m0-23V6m13 0v26"/><path fill="#fff" d="M11 16a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm13 13a5 5 0 1 0 0-10a5 5 0 0 0 0 10Zm13 13a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSEqualizer0)"/>`),
		g.Group(children),
	)
}