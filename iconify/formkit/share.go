package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Share(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7.5 10c-.28 0-.5-.22-.5-.5v-8c0-.28.22-.5.5-.5s.5.22.5.5v8c0 .28-.22.5-.5.5Z"/><path fill="currentColor" d="M9.5 4a.47.47 0 0 1-.35-.15L7.5 2.2L5.85 3.85c-.2.2-.51.2-.71 0c-.2-.2-.2-.51 0-.71l2.01-1.99c.2-.2.51-.2.71 0l2 2c.2.2.2.51 0 .71c-.1.1-.23.15-.35.15Zm3 11h-10c-.83 0-1.5-.67-1.5-1.5v-7C1 5.67 1.67 5 2.5 5h2c.28 0 .5.22.5.5s-.22.5-.5.5h-2c-.28 0-.5.22-.5.5v7c0 .28.22.5.5.5h10c.28 0 .5-.22.5-.5v-7c0-.28-.22-.5-.5-.5h-2c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h2c.83 0 1.5.67 1.5 1.5v7c0 .83-.67 1.5-1.5 1.5Z"/>`),
		g.Group(children),
	)
}