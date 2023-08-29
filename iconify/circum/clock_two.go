package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClockTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21.933A9.933 9.933 0 1 1 21.933 12A9.944 9.944 0 0 1 12 21.933Zm0-18.866A8.933 8.933 0 1 0 20.933 12A8.943 8.943 0 0 0 12 3.067Z"/><path fill="currentColor" d="M18 12.5h-6a.429.429 0 0 1-.34-.14c-.01 0-.01-.01-.02-.02a.429.429 0 0 1-.14-.34V6a.5.5 0 0 1 1 0v5.5H18a.5.5 0 0 1 0 1Z"/>`),
		g.Group(children),
	)
}