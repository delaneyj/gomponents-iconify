package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Six(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M373 39L186 331c14-2 28-4 42-4c125 0 227 102 227 227c0 126-102 227-227 227c-106 0-196-73-221-171c-5-17-7-37-7-56c0-43 13-83 33-117l6-10L311 0zM106 456l-19 28c-10 22-16 45-16 70c0 86 71 157 157 157s155-71 155-157s-69-156-155-156c-49 0-93 23-122 58z"/>`),
		g.Group(children),
	)
}