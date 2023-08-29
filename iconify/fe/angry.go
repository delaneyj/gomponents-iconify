package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Angry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feAngry0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feAngry1" fill="currentColor"><path id="feAngry2" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm0-2a8 8 0 1 0 0-16a8 8 0 0 0 0 16Zm4-3h-1.339s-.417-2.667-2.661-2.667S9.333 17 9.333 17H8a4 4 0 1 1 8 0Zm-5.494-8.01a1.5 1.5 0 1 1-1.52-.416L7.33 7.97a.5.5 0 0 1 .342-.94l2.82 1.026a.5.5 0 0 1 .015.934Zm2.897 0a.499.499 0 0 1 .016-.934l2.82-1.026a.5.5 0 1 1 .342.94l-1.659.603a1.5 1.5 0 1 1-1.519.417Z"/></g></g>`),
		g.Group(children),
	)
}