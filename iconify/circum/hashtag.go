package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hashtag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20.435 15.506H16.2l.61-7h3.63a.5.5 0 0 0 .5-.5a.5.5 0 0 0-.5-.5H16.9l.34-3.87a.509.509 0 0 0-.46-.54a.5.5 0 0 0-.54.46l-.35 3.95H8.9l.34-3.87a.509.509 0 0 0-.46-.54a.491.491 0 0 0-.54.46l-.35 3.95H3.565a.5.5 0 0 0-.5.5a.5.5 0 0 0 .5.5h4.24l-.62 7h-3.62a.5.5 0 0 0-.5.5a.5.5 0 0 0 .5.5h3.54l-.34 3.86a.508.508 0 0 0 .45.54h.05a.516.516 0 0 0 .5-.46l.34-3.94h7l-.34 3.86a.508.508 0 0 0 .45.54h.05a.516.516 0 0 0 .5-.46l.34-3.94h4.33a.5.5 0 0 0 .5-.5a.5.5 0 0 0-.5-.5Zm-5.25 0H8.2l.61-7h7Z"/>`),
		g.Group(children),
	)
}