package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QuitFullScreenBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M2 14c3.771 0 5.657 0 6.828 1.172C10 16.343 10 18.229 10 22M2 10c1.17 0 2.158 0 3-.035M10 2c0 3.771 0 5.657-1.172 6.828M14 22c0-3.771 0-5.657 1.172-6.828M22 14c-1.17 0-2.158 0-3 .035M22 10c-3.771 0-5.657 0-6.828-1.172C14 7.657 14 5.771 14 2"/>`),
		g.Group(children),
	)
}