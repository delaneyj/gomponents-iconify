package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RadioBtnPassive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 24C5.373 24 0 18.627 0 12S5.373 0 12 0s12 5.373 12 12c-.008 6.624-5.376 11.992-11.999 12zm0-21.6c-5.302 0-9.6 4.298-9.6 9.6s4.298 9.6 9.6 9.6s9.6-4.298 9.6-9.6c-.006-5.299-4.301-9.594-9.599-9.6H12z"/>`),
		g.Group(children),
	)
}