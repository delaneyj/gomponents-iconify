package dashicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Visibility(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M18.3 9.5C15 4.9 8.5 3.8 3.9 7.2c-1.2.9-2.2 2.1-3 3.4c.2.4.5.8.8 1.2c3.3 4.6 9.6 5.6 14.2 2.4c.9-.7 1.7-1.4 2.4-2.4c.3-.4.5-.8.8-1.2c-.3-.4-.5-.8-.8-1.1zm-8.2-2.3c.5-.5 1.3-.5 1.8 0s.5 1.3 0 1.8s-1.3.5-1.8 0s-.5-1.3 0-1.8zm-.1 7.7c-3.1 0-6-1.6-7.7-4.2C3.5 9 5.1 7.8 7 7.2c-.7.8-1 1.7-1 2.7c0 2.2 1.7 4.1 4 4.1c2.2 0 4.1-1.7 4.1-4v-.1c0-1-.4-2-1.1-2.7c1.9.6 3.5 1.8 4.7 3.5c-1.7 2.6-4.6 4.2-7.7 4.2z"/>`),
		g.Group(children),
	)
}