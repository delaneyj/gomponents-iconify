package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Truck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m405 107l64 85v107h-42q0 26-19 45t-45.5 19t-45-19t-18.5-45H171q0 26-19 45t-45.5 19t-45-19T43 299H0V64q0-18 12.5-30.5T43 21h298v86h64zM106.5 331q13.5 0 23-9.5t9.5-23t-9.5-22.5t-23-9t-22.5 9t-9 22.5t9 23t22.5 9.5zM395 139h-54v53h95zm-32.5 192q13.5 0 23-9.5t9.5-23t-9.5-22.5t-23-9t-22.5 9t-9 22.5t9 23t22.5 9.5z"/>`),
		g.Group(children),
	)
}