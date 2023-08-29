package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MsWordSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m2.015 5.621l1 4a.5.5 0 0 0 .901.156l.584-.876l.584.876a.5.5 0 0 0 .901-.156l1-4l-.97-.242l-.726 2.903l-.373-.56a.5.5 0 0 0-.832 0l-.373.56l-.726-2.903l-.97.242Z"/><path fill="currentColor" fill-rule="evenodd" d="M3.5 0A1.5 1.5 0 0 0 2 1.5V3h-.5A1.5 1.5 0 0 0 0 4.5v6A1.5 1.5 0 0 0 1.5 12H2v1.5A1.5 1.5 0 0 0 3.5 15h10a1.5 1.5 0 0 0 1.5-1.5v-12A1.5 1.5 0 0 0 13.5 0h-10Zm-2 4a.5.5 0 0 0-.5.5v6a.5.5 0 0 0 .5.5h6a.5.5 0 0 0 .5-.5v-6a.5.5 0 0 0-.5-.5h-6Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}