package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Threshold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path d="M26 4H6a2.002 2.002 0 0 0-2 2v20a2.002 2.002 0 0 0 2 2h20a2.002 2.002 0 0 0 2-2V6a2.002 2.002 0 0 0-2-2zM6 6h20v10h-2v2h2v2h-2v2h2v2h-2v2h-2v-2h-2v2h-2v-2h-2v2h-2v-2h-2v2h-2v-2H8v2H6z" fill="currentColor"/><path d="M8 20h2v2H8z" fill="currentColor"/><path d="M12 20h2v2h-2z" fill="currentColor"/><path d="M16 20h2v2h-2z" fill="currentColor"/><path d="M20 20h2v2h-2z" fill="currentColor"/><path d="M8 16h2v2H8z" fill="currentColor"/><path d="M16 16h2v2h-2z" fill="currentColor"/><path d="M20 16h2v2h-2z" fill="currentColor"/><path d="M20 12h2v2h-2z" fill="currentColor"/><path d="M20 8h2v2h-2z" fill="currentColor"/>`),
		g.Group(children),
	)
}