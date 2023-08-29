package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Copy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M10 4V0H3L0 3v9h6v4h10V4h-6zM3 1.414V3H1.414L3 1.414zM1 11V4h3V1h5v3L6 7v4H1zm8-5.586V7H7.414L9 5.414zM15 15H7V8h3V5h5v10z"/>`),
		g.Group(children),
	)
}