package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Restaurant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9 2h2v10H9z"/><path fill="currentColor" d="M14 11a4 4 0 0 1-8 0V2H4v9a6 6 0 0 0 5 5.91V30h2V16.91A6 6 0 0 0 16 11V2h-2zm8-9h-1v28h2V20h3a2 2 0 0 0 2-2V8a5.78 5.78 0 0 0-6-6zm4 16h-3V4.09c2.88.56 3 3.54 3 3.91z"/>`),
		g.Group(children),
	)
}