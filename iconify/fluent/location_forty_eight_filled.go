package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LocationFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M39.014 28.98A16.925 16.925 0 0 0 41 21c0-9.389-7.611-17-17-17S7 11.611 7 21a16.922 16.922 0 0 0 4 10.955l.02.025c.007.006.013.014.018.02H11l10.088 10.71a4 4 0 0 0 5.823 0L37 32h-.038l.016-.019l.002-.002c.072-.086.144-.172.215-.26a17.038 17.038 0 0 0 1.82-2.74Zm-15.01-1.48a6 6 0 1 1 0-12a6 6 0 0 1 0 12Z"/>`),
		g.Group(children),
	)
}