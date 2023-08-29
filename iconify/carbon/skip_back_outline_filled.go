package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipBackOutlineFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 2a14 14 0 1 0 14 14A14 14 0 0 0 16 2Zm-6 20H8V10h2Zm14-1a1 1 0 0 1-1.486.874l-9-5a1 1 0 0 1 0-1.748l9-5A1 1 0 0 1 24 11Z"/><path fill="currentColor" d="M22 19.301v-6.602L16.059 16L22 19.301z"/><path fill="none" d="M23 22a1.004 1.004 0 0 1-.486-.126l-9-5a1 1 0 0 1 0-1.748l9-5A1 1 0 0 1 24 11v10a1 1 0 0 1-1 1Zm-6.94-6L22 19.3v-6.6ZM10 22H8V10h2z"/>`),
		g.Group(children),
	)
}