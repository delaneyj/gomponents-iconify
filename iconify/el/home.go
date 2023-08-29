package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Home(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M600 0L56.645 422.323V1200h373.829V730.541h339.054V1200h373.828V422.323L600 0z"/>`),
		g.Group(children),
	)
}