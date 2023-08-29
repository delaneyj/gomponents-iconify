package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GiftEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M0 5h4.5v1H0V5zm1 4.79c0 .668.542 1.21 1.21 1.21H4.52V7H1v2.79zM7.64 4H3.36A1.26 1.26 0 0 1 2 2.5A1.41 1.41 0 0 1 3.42 1a2.12 2.12 0 0 1 2.1 1.69A2.13 2.13 0 0 1 7.62 1A1.43 1.43 0 0 1 9 2.52A1.27 1.27 0 0 1 7.64 4zm-2.82-.75a1.43 1.43 0 0 0-1.4-1.5a.68.68 0 0 0-.7.75a.68.68 0 0 0 .606.747c.031.003.063.004.094.003h1.4zm2.8 0a.68.68 0 0 0 .7-.75a.71.71 0 0 0-.7-.75a1.43 1.43 0 0 0-1.4 1.5h1.4zM6.52 5v1H11V5H6.52zm0 6h2.29A1.21 1.21 0 0 0 10 9.79V7H6.52v4z" fill="currentColor"/>`),
		g.Group(children),
	)
}