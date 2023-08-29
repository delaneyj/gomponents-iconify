package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M23 6H3a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h20a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3zm-3 5h-2V9h2v2zm3 0h-2V9h2v2zm-9 0h-2V9h2v2zm3 0h-2V9h2v2zm-6 0H9V9h2v2zm-6 0H3V9h2v2zm3 0H6V9h2v2zm12 8H6v-2h14v2zm1-4h-2v-2h2v2zm-6 0h-2v-2h2v2zm3 0h-2v-2h2v2zm-6 0h-2v-2h2v2zm-6 0H4v-2h2v2zm3 0H7v-2h2v2z"/>`),
		g.Group(children),
	)
}