package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KeySecurity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M8 .188a7.813 7.813 0 1 0 0 15.625A7.813 7.813 0 0 0 8 .187zM5.5 2.905A2.589 2.589 0 0 1 8.094 5.5A2.587 2.587 0 0 1 5.5 8.094A2.587 2.587 0 0 1 2.906 5.5A2.589 2.589 0 0 1 5.5 2.906zm11.094 8.719a9.232 9.232 0 0 1-1.032 1.813l7.813 7.812a.89.89 0 0 1 0 1.25a.892.892 0 0 1-1.25 0l-7.719-7.75A9.396 9.396 0 0 1 11 16.813v1.375c0 .44.371.812.813.812H14v2.188c0 .44.371.812.813.812H17v2.188c0 .44.372.812.813.812h2.218l.563.563c.342.341 3.768.512 4.625-.344c.857-.858.684-4.284.343-4.625l-8.968-8.969z"/>`),
		g.Group(children),
	)
}