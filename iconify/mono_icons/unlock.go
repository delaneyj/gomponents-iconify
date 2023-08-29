package mono_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 4c-1.648 0-3 1.352-3 3v3h9a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h1V7c0-2.752 2.248-5 5-5s5 2.248 5 5a1 1 0 1 1-2 0c0-1.648-1.352-3-3-3zm-6 8v8h12v-8H6z"/>`),
		g.Group(children),
	)
}