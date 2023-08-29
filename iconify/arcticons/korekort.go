package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Korekort(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 31.25V16.73h3.27c3.51 0 6.35 2.84 6.35 6.35v1.81c0 3.51-2.84 6.35-6.35 6.35H13.5Zm13.2-14.5v14.52m0-5.06l7.8-9.41m0 14.47l-5.98-7.26"/>`),
		g.Group(children),
	)
}