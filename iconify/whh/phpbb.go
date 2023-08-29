package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phpbb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm0-704q0-53-37.5-90.5t-90.5-37.5t-90.5 37.5t-37.5 90.5t37.5 90.5t90.5 37.5q-68 0-116 46.5t-66 127.5q-70-110-202-110q79 0 135.5-56t56.5-135.5t-56.5-136t-136-56.5t-135.5 56.5t-56 136t56 135.5t136 56q-116 0-186 88t-70 232q64 36 116 50t140 14t140-14t116-50q0-44-8-87q53 23 136 23q66 0 105-11.5t87-41.5q0-121-52.5-194t-139.5-73q53 0 90.5-37.5t37.5-90.5z"/>`),
		g.Group(children),
	)
}