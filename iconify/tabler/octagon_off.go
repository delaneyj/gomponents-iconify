package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OctagonOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7.647 3.653L8 3.3c.2-.2.4-.3.7-.3h6.6c.3 0 .5.1.7.3L20.7 8c.2.2.3.4.3.7v6.6c0 .3-.1.5-.3.7l-.35.35m-2 2l-2.353 2.353c-.2.2-.4.3-.7.3h-6.6c-.3 0-.5-.1-.7-.3l-4.7-4.7c-.2-.2-.3-.4-.3-.7v-6.6c0-.3.1-.5.3-.7l2.35-2.35M3 3l18 18"/>`),
		g.Group(children),
	)
}