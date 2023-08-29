package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControllerXbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path stroke-width=".2" stroke-linejoin="round" stroke="currentColor" d="M12 6.059c1.726 0 6.686-4.314 9.274.863C23.863 12.098 23.216 19 21.49 19c-4.313 0-1.725-4.314-9.49-4.314C4.235 14.686 6.824 19 2.51 19c-1.726 0-2.373-6.902.215-12.078c2.589-5.177 7.55-.863 9.275-.863zM12 7.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3z" fill="currentColor"/>`),
		g.Group(children),
	)
}