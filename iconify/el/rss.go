package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M1200 1200C1200 537.258 662.742 0 0 0v240.015c530.193 0 959.985 429.792 959.985 959.985H1200zm-480.029 0c0-397.646-322.324-719.971-719.971-719.971v239.94c265.097 0 480.029 214.934 480.029 480.029l239.942.002zm-479.956 0c0-132.549-107.466-240.015-240.015-240.015V1200h240.015z"/>`),
		g.Group(children),
	)
}