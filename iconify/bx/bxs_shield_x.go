package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsShieldX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M11.492 21.771c.294.157.663.157.957-.001c8.012-4.304 8.581-12.713 8.574-15.104a.988.988 0 0 0-.596-.903l-8.051-3.565a1.005 1.005 0 0 0-.813.001L3.57 5.765a.988.988 0 0 0-.592.891c-.034 2.379.445 10.806 8.514 15.115zM8.293 9.707l1.414-1.414L12 10.586l2.293-2.293l1.414 1.414L13.414 12l2.293 2.293l-1.414 1.414L12 13.414l-2.293 2.293l-1.414-1.414L10.586 12L8.293 9.707z" fill="currentColor"/>`),
		g.Group(children),
	)
}