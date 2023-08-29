package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pause(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 45v604c0 17 14 31 31 31h212c16 0 29-14 29-31V45c0-17-13-31-29-31H31C14 14 0 28 0 45zm393 0v604c0 17 14 31 30 31h212c17 0 31-14 31-31V45c0-17-14-31-31-31H423c-16 0-30 14-30 31z"/>`),
		g.Group(children),
	)
}