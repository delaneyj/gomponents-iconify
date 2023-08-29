package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpsideDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2C6.486 2 2 6.486 2 12s4.486 10 10 10s10-4.486 10-10S17.514 2 12 2zm0 18c-4.411 0-8-3.589-8-8s3.589-8 8-8s8 3.589 8 8s-3.589 8-8 8z"/><path fill="currentColor" d="M14.829 9.172c.181.181.346.38.488.592l1.658-1.119a6.063 6.063 0 0 0-1.621-1.62a5.963 5.963 0 0 0-2.148-.903a5.985 5.985 0 0 0-5.448 1.634a5.993 5.993 0 0 0-.733.889l1.657 1.119a4.017 4.017 0 0 1 2.51-1.683a3.989 3.989 0 0 1 3.637 1.091z"/><circle cx="15.5" cy="13.5" r="1.5" fill="currentColor"/><circle cx="8.507" cy="13.507" r="1.493" fill="currentColor"/>`),
		g.Group(children),
	)
}