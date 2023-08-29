package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Businesswoman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M16 11c-.9-.8-2.2-.9-3.4-1l1 2.1l-3.6 3.7l-3.6-3.6l1-2.2c-1.2 0-2.5.2-3.4 1c-.8.7-1 1.9-1 3.1v2.8s3.4 1.2 7 1.1c3.6.1 7-1.1 7-1.1v-2.8c0-1.1-.2-2.3-1-3.1zM6.6 9.3c.8 0 2-.4 2.2-.7c-.8-1-1.5-2-.8-3.9c0 0 1.1 1.2 4.3 1.5c0 1-.5 1.7-1.1 2.4c.2.3 1.4.7 2.2.7s1.4-.2 1.4-.5s-1.3-1.3-1.6-2.2c-.3-.9-.1-1.9-.5-3.1c-.6-1.4-2-1.5-2.7-1.5c-.7 0-2.1.1-2.7 1.5c-.4 1.2-.2 2.2-.5 3.1c-.3.9-1.6 1.9-1.6 2.2c0 .3.6.5 1.4.5z"/><path fill="currentColor" d="m10 11l-2.3-1l2.3 5.8l2.3-5.8z"/>`),
		g.Group(children),
	)
}