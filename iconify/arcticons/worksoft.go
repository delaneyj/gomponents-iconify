package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Worksoft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.659 4.5a9.243 9.243 0 1 0 9.271 9.245v-.002A9.257 9.257 0 0 0 28.66 4.5ZM19.39 14.517v0Zm-4.648 9.91a4.665 4.665 0 0 0-4.672 4.658v0a4.672 4.672 0 1 0 4.673-4.658Zm16.002 9.317a4.886 4.886 0 0 0-4.893 4.878v0a4.886 4.886 0 0 0 4.893 4.878h0a4.885 4.885 0 0 0 4.893-4.878v0a4.886 4.886 0 0 0-4.893-4.878Zm-16.01.005h0Z"/>`),
		g.Group(children),
	)
}