package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alarm(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 7v5h5m4.004-7.429L17.939 2M6.064 2L3 4.571M12 20a8 8 0 1 1 0-16a8 8 0 0 1 0 16Z"/>`),
		g.Group(children),
	)
}