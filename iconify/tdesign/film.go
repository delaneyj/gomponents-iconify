package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Film(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 1 9 9c0 2.453-.98 4.676-2.573 6.3A8.967 8.967 0 0 1 12 21a9 9 0 1 1 0-18Zm6.326 18c.55-.387 1.061-.822 1.53-1.3A10.967 10.967 0 0 0 23 12c0-6.075-4.925-11-11-11S1 5.925 1 12s4.925 11 11 11h12v-2h-5.674ZM12 6a1 1 0 1 0 0 2a1 1 0 0 0 0-2ZM9 7a3 3 0 1 1 6 0a3 3 0 0 1-6 0Zm-2 4a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0Zm13-1a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0Zm-2 4a1 1 0 1 0 0 2a1 1 0 0 0 0-2Zm-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0Z"/>`),
		g.Group(children),
	)
}