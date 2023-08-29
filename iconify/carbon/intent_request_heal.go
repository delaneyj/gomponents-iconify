package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IntentRequestHeal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27 20c-.8 0-1.5.3-2.1.9l-.4.5l-.4-.5c-.6-.6-1.4-.9-2.1-.9s-1.5.3-2.1.9c-1.2 1.2-1.2 3.1 0 4.3l4.6 4.8l4.6-4.8c1.2-1.2 1.2-3.1 0-4.3c-.5-.6-1.3-.9-2.1-.9z"/><path fill="currentColor" d="M16.6 28.6L4 16L16 4l12.6 12.6l1.4-1.4L17.5 2.6c-.8-.8-2.1-.8-2.9 0l-12 11.9c-.8.8-.8 2.1 0 2.9L15.1 30l1.5-1.4z"/>`),
		g.Group(children),
	)
}