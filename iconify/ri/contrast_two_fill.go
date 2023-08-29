package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContrastTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21.997c-5.523 0-10-4.478-10-10c0-5.523 4.477-10 10-10s10 4.477 10 10c0 5.522-4.477 10-10 10Zm-6.671-5.575A8 8 0 1 0 16.425 5.325a8.997 8.997 0 0 1-2.304 8.793a8.997 8.997 0 0 1-8.792 2.304Z"/>`),
		g.Group(children),
	)
}