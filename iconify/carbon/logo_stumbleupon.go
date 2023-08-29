package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoStumbleupon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 2a14 14 0 1 0 14 14A14 14 0 0 0 16 2Zm-.09 10.45a.84.84 0 0 0-.84.84v5.14a3.55 3.55 0 0 1-7.1 0v-2.34h2.71v2.24a.84.84 0 0 0 1.68 0v-5a3.55 3.55 0 0 1 7.09 0v1l-1.58.51l-1.12-.51v-1a.85.85 0 0 0-.84-.88Zm7.93 6a3.55 3.55 0 0 1-7.09 0v-2.31l1.12.51l1.58-.51v2.29a.84.84 0 0 0 1.68 0v-2.24h2.71Z"/>`),
		g.Group(children),
	)
}