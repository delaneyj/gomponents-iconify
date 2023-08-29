package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LikeSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M9.582 1.052c-.75-.209-1.337.35-1.546.872c-.24.6-.452 1.02-.705 1.523c-.157.312-.33.654-.534 1.09c-.474 1.01-.948 1.657-1.292 2.045a4.064 4.064 0 0 1-.405.403c-.047.039-.081.065-.102.08l-.016.012L3.11 8.182a2 2 0 0 0-.856 2.425l.52 1.385a2 2 0 0 0 1.273 1.204l5.356 1.682a2.5 2.5 0 0 0 3.148-1.68l1.364-4.646a2 2 0 0 0-1.919-2.564h-1.385c.066-.227.134-.479.195-.74c.131-.561.243-1.203.233-1.738c-.01-.497-.06-1.018-.264-1.462c-.22-.475-.603-.832-1.193-.996z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}