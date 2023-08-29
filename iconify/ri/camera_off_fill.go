package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CameraOffFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.587 21H3a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h.586L1.395 2.808l1.414-1.414l19.799 19.798l-1.415 1.415L19.587 21ZM7.557 8.97a6 6 0 0 0 8.475 8.475l-1.417-1.417a4 4 0 0 1-5.642-5.642L7.556 8.97ZM22 17.786l-4.045-4.045a6 6 0 0 0-6.695-6.695L8.108 3.892L9.001 3h6l2 2h4a1 1 0 0 1 1 1v11.786Zm-8.49-8.492a4.014 4.014 0 0 1 2.198 2.198L13.51 9.294Z"/>`),
		g.Group(children),
	)
}