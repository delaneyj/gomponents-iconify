package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func End(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M16 5c-.551 0-1 .449-1 1v4.875L7.906 5.25a1.339 1.339 0 0 0-1.281.063C6.239 5.545 6 5.963 6 6.405v13.188c0 .443.239.86.625 1.093c.208.127.449.188.688.188c.2 0 .408-.033.593-.125L15 15.125V20c0 .551.449 1 1 1h3c.551 0 1-.449 1-1V6c0-.551-.449-1-1-1h-3z"/>`),
		g.Group(children),
	)
}