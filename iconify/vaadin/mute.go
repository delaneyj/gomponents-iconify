package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mute(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15.2 0L11 4.2V3c0-1.7-1.3-3-3-3S5 1.3 5 3v4c0 .9.4 1.7 1 2.2l-.8.8C4.5 9.4 4 8.5 4 7.5V5c-.6 0-1 .4-1 1v1.5c0 1.3.6 2.4 1.5 3.2L0 15.3v.7h.7L16 .6V0h-.8zm-2.7 5.1l-.5.5v1.9c0 1.9-1.8 3.5-3.8 3.5h-.4c-.3 0-.6-.1-.9-.1l-.9.7c.3.1.6.2 1 .3V14c-3 0-2.5 2-2.5 2h7s.5-2-2.5-2v-2.1c2.2-.4 4-2.2 4-4.4V6c0-.4-.2-.7-.5-.9z"/><path fill="currentColor" d="M11 7v-.4L7.7 10H8c1.7 0 3-1.4 3-3z"/>`),
		g.Group(children),
	)
}