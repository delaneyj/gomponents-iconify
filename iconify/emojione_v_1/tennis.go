package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tennis(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#c3d82e" d="M57.55 14.457c9.694 14.11 6.111 33.405-7.998 43.09c-14.11 9.694-33.406 6.113-43.1-8c-9.694-14.11-6.111-33.407 8-43.1c14.11-9.69 33.404-6.108 43.09 8"/><path fill="#fff" d="M45.336 59.995s-14.8-32.978.966-55.494l1.323.73s-16.514 22.624.848 53.076l-3.137 1.69M6.46 49.552S21.043 27.893 17.024 4.869l-1.796 1.235s7.87 6.426-9.891 41.812"/>`),
		g.Group(children),
	)
}