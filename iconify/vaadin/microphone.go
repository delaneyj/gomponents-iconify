package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Microphone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M8 10c-1.7 0-3-1.3-3-3V3c0-1.6 1.3-3 3-3c1.6 0 3 1.3 3 3v4c0 1.6-1.4 3-3 3z"/><path fill="currentColor" d="M12 5v2.5c0 1.9-1.8 3.5-3.8 3.5h-.4C5.8 11 4 9.4 4 7.5V5c-.6 0-1 .4-1 1v1.5c0 2.2 1.8 4.1 4 4.4V14c-3 0-2.5 2-2.5 2h7s.5-2-2.5-2v-2.1c2.2-.4 4-2.2 4-4.4V6c0-.6-.4-1-1-1z"/>`),
		g.Group(children),
	)
}