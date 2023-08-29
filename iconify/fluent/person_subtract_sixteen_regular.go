package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonSubtractSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.222 13.957C3.555 13.653 2 11.803 2 10v-.5A1.5 1.5 0 0 1 3.5 8h2.1a5.463 5.463 0 0 0-.393 1H3.5a.5.5 0 0 0-.5.5v.5c0 1.128.882 2.333 2.502 2.8c.192.416.435.804.72 1.157Zm3.404-8.888a2.75 2.75 0 1 0-3.299 1.848a5.53 5.53 0 0 1 1.078-.964a1.75 1.75 0 1 1 .836-.47a5.46 5.46 0 0 1 1.385-.414ZM10.5 15a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Zm-2.121-5h4.242a.5.5 0 1 1 0 1H8.38a.5.5 0 0 1 0-1Z"/>`),
		g.Group(children),
	)
}