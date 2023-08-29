package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zelda(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1013 896H525q-20-18-7-31l232-410q8-7 19-7t19 7l232 410q13 12-7 31zM269 448q-20-18-7-31L494 7q8-7 19-7t19 7l232 410q13 12-7 31H269zm239 417q13 12-7 31H13q-20-18-7-31l232-410q8-7 19-7t19 7z"/>`),
		g.Group(children),
	)
}