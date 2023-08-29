package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScreenRotation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSScreenRotation0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M4 24c0 11.046 8.954 20 20 20l-5-5m25-15c0-11.046-8.954-20-20-20l5 5"/><path fill="#fff" d="M30 41L7 18L18 7l23 23l-11 11Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSScreenRotation0)"/>`),
		g.Group(children),
	)
}