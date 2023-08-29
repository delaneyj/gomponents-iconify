package flat_color_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sms(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="#009688" d="M37 39H11l-6 6V11c0-3.3 2.7-6 6-6h26c3.3 0 6 2.7 6 6v22c0 3.3-2.7 6-6 6z"/><g fill="#fff"><circle cx="24" cy="22" r="3"/><circle cx="34" cy="22" r="3"/><circle cx="14" cy="22" r="3"/></g>`),
		g.Group(children),
	)
}