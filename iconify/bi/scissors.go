package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scissors(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M3.5 3.5c-.614-.884-.074-1.962.858-2.5L8 7.226L11.642 1c.932.538 1.472 1.616.858 2.5L8.81 8.61l1.556 2.661a2.5 2.5 0 1 1-.794.637L8 9.73l-1.572 2.177a2.5 2.5 0 1 1-.794-.637L7.19 8.61L3.5 3.5zm2.5 10a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0zm7 0a1.5 1.5 0 1 0-3 0a1.5 1.5 0 0 0 3 0z"/>`),
		g.Group(children),
	)
}