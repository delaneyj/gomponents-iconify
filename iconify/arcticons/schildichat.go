package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Schildichat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13 33.54l-3.3.4l-2.4-.6l-2.8-.4l.6-2l2.1-6.2l2.3-4.4l.9-1.4l4.6-1.7m-.83-3.9L10.47 19"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m33 17.34l4.6 1.7l.9 1.4l2.3 4.4l2.1 6.2l.6 2.1l-2.8.4l-2.4.6l-3.2-.5"/><path d="M24.07 36.74c1.1 0 1.4.2 3.3.2c.5 0 1.8-.9 3.2-1.6a28.83 28.83 0 0 1 4.6-1.7c.4-.8-5.8-12.7-6.5-13h0a12.58 12.58 0 0 0-4.6-.7a13.22 13.22 0 0 0-4.6.7h0c-.8.3-6.9 12.2-6.5 13a33 33 0 0 1 4.6 1.7a14.86 14.86 0 0 0 3.2 1.6c1.8 0 2.1-.2 3.3-.2Z"/><path d="M28.47 20.54a35.7 35.7 0 0 1 2.8-2.4a5.29 5.29 0 0 1 1.6-.9l.9-3.9h0l-5.1-2S25.57 11 24 11h0c-1.6 0-4.7.3-4.7.3l-5.2 2h0l.9 3.9a11.1 11.1 0 0 1 1.6.9a35.7 35.7 0 0 1 2.8 2.4M37.57 19l-3.7-5.7"/></g>`),
		g.Group(children),
	)
}