package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M249 42Q144 42 59 97l190 237L438 97q-84-55-189-55zm-.5-39q49.5 0 96 11T425 43.5t47 27T497 88L249 397L0 88q12-9 25-17.5t47.5-27T153 14t95.5-11z"/>`),
		g.Group(children),
	)
}