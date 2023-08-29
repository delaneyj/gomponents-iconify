package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaunchStudyTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="23" cy="7" r="1" fill="currentColor"/><path d="M26 4v6H6V4h20m0-2H6a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2z" fill="currentColor"/><path d="M20.59 21.41L17 17.83V30h-2V17.83l-3.59 3.58L10 20l6-6l6 6l-1.41 1.41z" fill="currentColor"/><path d="M8 6h12v2H8z" fill="currentColor"/>`),
		g.Group(children),
	)
}