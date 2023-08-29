package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlignCenterVerticalOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M15 7.5h-3.5m-8 0H0m8.5 0h-2m5 3.5h-3V4h3v7Zm-5 2.5h-3v-12h3v12Z"/>`),
		g.Group(children),
	)
}