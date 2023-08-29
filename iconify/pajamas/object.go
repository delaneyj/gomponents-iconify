package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Object(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 14.5c.8 0 1.566-.145 2.274-.409l.665-2.046h2.15a6.472 6.472 0 0 0 1.405-4.327l-1.739-1.263l.665-2.045a6.512 6.512 0 0 0-3.68-2.674L8 3L6.26 1.736A6.512 6.512 0 0 0 2.58 4.41l.665 2.045l-1.739 1.263a6.472 6.472 0 0 0 1.406 4.327h2.15l.664 2.046A6.486 6.486 0 0 0 8 14.5ZM8 16A8 8 0 1 0 8 0a8 8 0 0 0 0 16ZM8 5.286l2.853 2.073l-1.09 3.354H6.237L5.147 7.36L8 5.286Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}