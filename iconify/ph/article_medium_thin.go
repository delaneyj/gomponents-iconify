package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArticleMediumThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M52 136a4 4 0 0 1-4 4H24a4 4 0 0 1 0-8h12V60H24a4 4 0 0 1 0-8h16a4 4 0 0 1 3.39 1.87L80 112.45l36.61-58.57A4 4 0 0 1 120 52h16a4 4 0 0 1 0 8h-12v72h12a4 4 0 0 1 0 8h-24a4 4 0 0 1 0-8h4V70l-32.61 52.12a4 4 0 0 1-6.78 0L44 70v62h4a4 4 0 0 1 4 4Zm116-28h72a4 4 0 0 0 0-8h-72a4 4 0 0 0 0 8Zm72 24h-72a4 4 0 0 0 0 8h72a4 4 0 0 0 0-8Zm0 32H72a4 4 0 0 0 0 8h168a4 4 0 0 0 0-8Zm0 32H72a4 4 0 0 0 0 8h168a4 4 0 0 0 0-8Z"/>`),
		g.Group(children),
	)
}