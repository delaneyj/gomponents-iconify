package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eyedropper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15 1c-1.8-1.8-3.7-.7-4.6.1c-.4.4-.7.9-.7 1.5c0 1.1-1.1 1.8-2.1 1.5L7.5 4l-.7.8l.7.7l-6 6l-.8 2.3l-.7.7L1.5 16l.8-.8l2.3-.8l6-6l.7.7l.7-.6l-.1-.2c-.3-1 .4-2.1 1.5-2.1c.6 0 1.1-.2 1.4-.6c.9-.9 2-2.8.2-4.6zM3.9 13.6l-2 .7l-.2.1l.1-.2l.7-2l5.8-5.8l1.5 1.5l-5.9 5.7z"/>`),
		g.Group(children),
	)
}