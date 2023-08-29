package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Linkhub(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.865 12.73l3.942-3.942a8.781 8.781 0 0 1 12.405 12.404l-5.635 5.615a8.742 8.742 0 0 1-12.362.023l-.023-.023"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m25.135 35.27l-3.942 3.942A8.781 8.781 0 0 1 8.788 26.808l5.635-5.615a8.742 8.742 0 0 1 12.362-.023l.023.023"/>`),
		g.Group(children),
	)
}