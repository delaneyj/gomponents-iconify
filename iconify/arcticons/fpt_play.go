package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FptPlay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.5C39 5.5 42.5 9 42.5 24S39 42.5 24 42.5S5.5 39 5.5 24S9 5.5 24 5.5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.77 24c0-10.376-.264-10.224 8.723-5.036c8.986 5.189 8.986 4.884 0 10.072c-8.987 5.188-8.723 5.34-8.723-5.036v0Z"/>`),
		g.Group(children),
	)
}