package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Prohibited(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M50 12.5c-20.712 0-37.5 16.793-37.5 37.502C12.5 70.712 29.288 87.5 50 87.5s37.5-16.788 37.5-37.498C87.5 29.293 70.712 12.5 50 12.5zm0 65.061c-15.199 0-27.561-12.362-27.561-27.559a27.435 27.435 0 0 1 6.4-17.636l38.795 38.795A27.431 27.431 0 0 1 50 77.561zm21.161-9.926L32.366 28.839A27.431 27.431 0 0 1 50 22.439c15.198 0 27.56 12.367 27.56 27.562a27.43 27.43 0 0 1-6.399 17.634z"/>`),
		g.Group(children),
	)
}