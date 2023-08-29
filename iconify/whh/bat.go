package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m1005.594 881l-125 125q-19 19-46 19t-46.5-19t-19.5-46.5t19-46.5l18-18l-149-149q-38-38-107-76.5t-130.5-70t-83.5-53.5l-320-320q-23-22-12-80t47.5-94.5t94.5-47.5t80 12l320 320q22 22 53.5 83.5t70 130.5t76.5 107l149 149l18-18q19-19 46.5-19t46.5 19.5t19 46.5t-19 46z"/>`),
		g.Group(children),
	)
}