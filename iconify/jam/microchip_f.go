package jam

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrochipF(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11 17.858v2a1 1 0 0 1-2 0v-2H7v2a1 1 0 0 1-2 0v-2c0-.057.005-.112.014-.166a3.008 3.008 0 0 1-1.848-1.848a.915.915 0 0 1-.166.014H1a1 1 0 0 1 0-2h2v-2H1a1 1 0 0 1 0-2h2v-2H1a1 1 0 1 1 0-2h2c.057 0 .112.005.166.014c.3-.864.984-1.548 1.848-1.848A1.007 1.007 0 0 1 5 3.858v-2a1 1 0 1 1 2 0v2h2v-2a1 1 0 1 1 2 0v2h2v-2a1 1 0 0 1 2 0v2c0 .056-.005.112-.014.166c.864.3 1.548.984 1.848 1.848A.915.915 0 0 1 17 5.858h2a1 1 0 0 1 0 2h-2v2h2a1 1 0 0 1 0 2h-2v2h2a1 1 0 0 1 0 2h-2c-.057 0-.112-.005-.166-.014a3.008 3.008 0 0 1-1.848 1.848c.01.054.014.11.014.166v2a1 1 0 0 1-2 0v-2h-2zm-3-10a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1H8z"/>`),
		g.Group(children),
	)
}