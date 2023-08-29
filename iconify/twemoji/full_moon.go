package twemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FullMoon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<circle cx="18" cy="18" r="18" fill="#FFD983"/><g fill="#FFCC4D"><circle cx="10.5" cy="8.5" r="3.5"/><circle cx="20" cy="17" r="3"/><circle cx="24.5" cy="28.5" r="3.5"/><circle cx="22" cy="5" r="2"/><circle cx="3" cy="18" r="1"/><circle cx="30" cy="9" r="1"/><circle cx="15" cy="31" r="1"/><circle cx="32" cy="19" r="2"/><circle cx="10" cy="23" r="2"/></g>`),
		g.Group(children),
	)
}