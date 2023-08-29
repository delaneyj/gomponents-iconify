package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Autopilot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M23.6 5c-.2-.2-.5-.4-.8-.4c-2.3-.1-5.2-2.5-7.1-3.5c-1.2-.6-2-1-2.6-1.1h-.4c-.6.1-1.4.5-2.6 1.1c-1.9 1-4.8 3.4-7.1 3.5c-.2.1-.4.2-.6.4c-.2.3-.3.6-.3.9c.5 10 4.1 16.2 10.4 19.8c.2.1.3.1.5.1s.4 0 .5-.1c6.3-3.6 9.9-9.8 10.4-19.8c0-.3-.1-.6-.3-.9zM19 15.5l-5-2.7V16l2 1.5v.5l-3-.5l-3 .5v-.5l2-1.5v-3.2l-5 2.7v-1.1l5-4.3V6c0-.6.4-1 1-1s1 .4 1 1v4.1l5 4.3v1.1z"/>`),
		g.Group(children),
	)
}