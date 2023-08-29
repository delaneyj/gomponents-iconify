package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HuaweiNotepad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.08 9.897h6.764a3 3 0 0 1 3 3V39.5a3 3 0 0 1-3 3H9.157a3 3 0 0 1-3-3V12.897a3 3 0 0 1 3-3h5.507"/><rect width="17.415" height="9.307" x="14.665" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.981 26.455h13.076"/>`),
		g.Group(children),
	)
}