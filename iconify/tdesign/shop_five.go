package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShopFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.28 2h17.44l.972 2.914c.181.543.391 1.332.163 2.154A3.993 3.993 0 0 1 21 8.646V20h1v2H2v-2h1V8.646A3.988 3.988 0 0 1 2 6v-.162L3.28 2ZM5 9.874V20h3v-7h8v7h3V9.874a4.004 4.004 0 0 1-4-1.228A3.99 3.99 0 0 1 12 10a3.99 3.99 0 0 1-3-1.354a3.99 3.99 0 0 1-4 1.228ZM10 6a2 2 0 1 0 4 0V4h-4v2ZM8 4H4.72l-.715 2.146c.039.533.285 1.008.662 1.345A2 2 0 0 0 8 6V4Zm8 0v2a2 2 0 0 0 3.928.535c.059-.213.026-.512-.133-.989L19.279 4H16Zm-2 16v-5h-4v5h4Z"/>`),
		g.Group(children),
	)
}