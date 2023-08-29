package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Opera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M16 8c0 2.369-1.031 4.5-2.669 5.963c-2.053 1-3.966.3-4.597-.137c2.016-.441 3.537-2.878 3.537-5.825s-1.522-5.384-3.537-5.828c.634-.438 2.547-1.137 4.597-.138A7.99 7.99 0 0 1 16 8.001z"/><path fill="currentColor" d="M5.366 3.491C4.482 4.535 3.91 6.078 3.872 7.813v.378c.038 1.731.613 3.275 1.497 4.319c1.147 1.491 2.853 2.434 4.759 2.434a5.768 5.768 0 0 0 3.206-.978a7.984 7.984 0 0 1-5.715 2.025A8 8 0 0 1 8 0h.031a7.952 7.952 0 0 1 5.303 2.038a5.773 5.773 0 0 0-3.206-.981c-1.906 0-3.612.944-4.763 2.434z"/>`),
		g.Group(children),
	)
}