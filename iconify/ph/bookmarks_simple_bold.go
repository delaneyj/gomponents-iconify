package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarksSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M156 56H60a20 20 0 0 0-20 20v152a12 12 0 0 0 19 9.76l49-35l49 35a12 12 0 0 0 19-9.76V76a20 20 0 0 0-20-20Zm-4 148.68l-37-26.45a12 12 0 0 0-14 0l-37 26.45V80h88ZM216 36v152a12 12 0 0 1-24 0V40H92a12 12 0 0 1 0-24h104a20 20 0 0 1 20 20Z"/>`),
		g.Group(children),
	)
}