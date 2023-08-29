package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SixKey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSixKey0"><g fill="none" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#fff" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" rx="3"/><path stroke="#000" d="M24.5 33a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="M29.596 17.74C28.778 15.545 26.804 14 24.5 14c-3.038 0-5.5 2.686-5.5 6v7"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSixKey0)"/>`),
		g.Group(children),
	)
}