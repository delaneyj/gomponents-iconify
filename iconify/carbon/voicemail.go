package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Voicemail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24 10a6 6 0 0 0-4.46 10h-7.08A6 6 0 1 0 8 22h16a6 6 0 0 0 0-12ZM4 16a4 4 0 1 1 4 4a4 4 0 0 1-4-4Zm20 4a4 4 0 1 1 4-4a4 4 0 0 1-4 4Z"/>`),
		g.Group(children),
	)
}