package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserGroup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M12 5a4 4 0 0 0-2.906 6.75a5.998 5.998 0 0 0-2.974 6.449A1 1 0 0 0 7.1 19h9.8a1 1 0 0 0 .98-.801a5.998 5.998 0 0 0-2.974-6.45A4 4 0 0 0 12 5zm4.584 5.779C16.85 10.243 17 9.639 17 9s-.15-1.243-.416-1.779A3 3 0 1 1 21.4 10.8A4 4 0 0 1 23 14v1a1 1 0 0 1-1 1h-3.083a6.006 6.006 0 0 0-3.012-4.25a4.01 4.01 0 0 0 .651-.917a3.56 3.56 0 0 1 .044-.033l-.016-.021zm-8.49.97a4.01 4.01 0 0 1-.65-.916A3.902 3.902 0 0 0 7.4 10.8l.016-.021A3.984 3.984 0 0 1 7 9c0-.639.15-1.243.416-1.779A3 3 0 1 0 2.6 10.8a3.994 3.994 0 0 0-1.372 4.534A1 1 0 0 0 2.17 16h2.912a6.006 6.006 0 0 1 3.011-4.25z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}