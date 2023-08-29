package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoChromeFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 23c6.075 0 11-4.925 11-11c0-1.131-.17-2.223-.488-3.25h-7.048A4.733 4.733 0 0 1 16.75 12a4.73 4.73 0 0 1-.636 2.376l-4.96 8.592c.279.021.561.032.846.032Z"/><path fill="currentColor" d="m9.56 22.728l3.523-6.102a4.748 4.748 0 0 1-5.197-2.25L2.925 5.782A10.95 10.95 0 0 0 1 12c0 5.236 3.659 9.618 8.56 10.728Z"/><path fill="currentColor" d="M3.93 4.524A10.97 10.97 0 0 1 12 1a11 11 0 0 1 9.924 6.25H12a4.752 4.752 0 0 0-4.548 3.374l-3.521-6.1Z"/><path fill="currentColor" d="m9.185 13.625l-.004-.006a3.25 3.25 0 1 1 .003.005Z"/>`),
		g.Group(children),
	)
}