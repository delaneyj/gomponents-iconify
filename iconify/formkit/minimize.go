package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minimize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M5.5 10a.47.47 0 0 1-.35-.15c-.2-.2-.2-.51 0-.71l6.5-6.49c.2-.2.51-.2.71 0c.2.2.2.51 0 .71l-6.5 6.5c-.1.1-.23.15-.35.15Z"/><path fill="currentColor" d="M9.5 10h-4c-.28 0-.5-.22-.5-.5v-4c0-.28.22-.5.5-.5s.5.22.5.5V9h3.5c.28 0 .5.22.5.5s-.22.5-.5.5Z"/><path fill="currentColor" d="M11.5 14h-9c-.83 0-1.5-.67-1.5-1.5v-9C1 2.67 1.67 2 2.5 2h5c.28 0 .5.22.5.5s-.22.5-.5.5h-5c-.28 0-.5.22-.5.5v9c0 .28.22.5.5.5h9c.28 0 .5-.22.5-.5v-5c0-.28.22-.5.5-.5s.5.22.5.5v5c0 .83-.67 1.5-1.5 1.5Z"/>`),
		g.Group(children),
	)
}