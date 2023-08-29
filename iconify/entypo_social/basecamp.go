package entypo_social

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Basecamp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10 2C5.327 2 .7 8.481.7 14.422C.7 15.799 5.234 18 10 18s9.3-2.201 9.3-3.578C19.3 8.481 14.673 2 10 2zm.006 13.615c-5.198 0-6.673-2.068-6.673-2.722c0-1.287 2.13-4.485 2.906-4.485c.719 0 1.542 1.811 2.314 1.811c1.241 0 2.567-3.954 3.579-3.954s4.601 5.178 4.601 6.749c0 .271-1.084 2.601-6.727 2.601z"/>`),
		g.Group(children),
	)
}