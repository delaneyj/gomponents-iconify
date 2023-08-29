package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 6H4a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h24a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2ZM4 24V8h24v16Z"/><path fill="currentColor" d="M22 11h-4v10h4a3 3 0 0 0 3-3v-4a3 3 0 0 0-3-3zm1 7a1 1 0 0 1-1 1h-2v-6h2a1 1 0 0 1 1 1zm-10-7v4h-3v-4H8v10h2v-4h3v4h2V11h-2z"/>`),
		g.Group(children),
	)
}