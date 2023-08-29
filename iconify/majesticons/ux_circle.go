package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UxCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12zm13.8-2.6a1 1 0 0 0-1.6 1.2l1.05 1.4l-1.05 1.4a1 1 0 0 0 1.6 1.2l.7-.933l.7.933a1 1 0 0 0 1.6-1.2L16.75 12l1.05-1.4a1 1 0 0 0-1.6-1.2l-.7.933l-.7-.933zM8 10a1 1 0 1 0-2 0v2c0 .493.14 1.211.588 1.834C7.074 14.51 7.874 15 9 15s1.926-.492 2.412-1.166c.448-.623.588-1.34.588-1.834v-2a1 1 0 1 0-2 0v2c0 .173-.06.456-.212.666c-.114.159-.314.334-.788.334c-.474 0-.674-.175-.788-.334A1.239 1.239 0 0 1 8 12v-2z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}