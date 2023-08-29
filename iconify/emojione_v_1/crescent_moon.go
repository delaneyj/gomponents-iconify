package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CrescentMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#f5eb35" d="M44.1 31.02c1-17.417-7.735-26.423-10.353-30.549c14.396 3.137 30.09 10.526 30.09 31.28c0 17.689-14.339 32.03-32.03 32.03a31.797 31.797 0 0 1-16.11-4.381c8.664-4.876 27.4-10.957 28.4-28.373"/><g fill="#e0cf35"><path d="M56.745 40.26a2.544 2.544 0 0 0 2.548 2.545a2.548 2.548 0 1 0-2.548-2.545"/><circle cx="32.27" cy="57.38" r="3.486"/><circle cx="47.895" cy="45.18" r="5.861"/><path d="M47.667 12.575a2.07 2.07 0 1 0 4.14.001a2.07 2.07 0 0 0-4.14-.001"/></g>`),
		g.Group(children),
	)
}