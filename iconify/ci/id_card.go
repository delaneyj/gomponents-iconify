package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IdCard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 20H4a1.943 1.943 0 0 1-2-1.876V5.875A1.942 1.942 0 0 1 4 4h16a1.942 1.942 0 0 1 2 1.875v12.25A1.943 1.943 0 0 1 20 20ZM4 6v11.989L20 18V6.011L4 6Zm9.43 10H6a3.21 3.21 0 0 1 1.093-2.14a3.829 3.829 0 0 1 2.622-1.11c.984.02 1.923.417 2.622 1.11A3.212 3.212 0 0 1 13.43 16ZM18 15h-3v-2h3v2Zm-8.285-3a1.934 1.934 0 0 1-2-2a1.935 1.935 0 0 1 2-2a1.935 1.935 0 0 1 2 2a1.934 1.934 0 0 1-2 2ZM18 11h-4V9h4v2Z"/>`),
		g.Group(children),
	)
}