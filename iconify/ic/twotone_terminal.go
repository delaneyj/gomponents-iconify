package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneTerminal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4 18h16V8H4v10zm8-3h6v2h-6v-2zm-5.91-4.59L7.5 9l4 4l-4 4l-1.41-1.41L8.67 13l-2.58-2.59z" opacity=".3"/><path fill="currentColor" d="M12 15h6v2h-6z"/><path fill="currentColor" d="M20 4H4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16c1.1 0 2-.9 2-2V6a2 2 0 0 0-2-2zm0 14H4V8h16v10z"/><path fill="currentColor" d="m7.5 17l4-4l-4-4l-1.41 1.41L8.67 13l-2.58 2.59z"/>`),
		g.Group(children),
	)
}