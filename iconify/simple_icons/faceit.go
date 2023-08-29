package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Faceit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M23.999 2.705a.167.167 0 0 0-.312-.1a1141.27 1141.27 0 0 0-6.053 9.375H.218c-.221 0-.301.282-.11.352c7.227 2.73 17.667 6.836 23.5 9.134c.15.06.39-.08.39-.18z"/>`),
		g.Group(children),
	)
}