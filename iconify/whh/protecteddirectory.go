package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Protecteddirectory(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.38 896h-768q-53 0-90.5-37.5T.38 768V256q0-26 18.5-45t45.5-19h480q0-7 11.5-28t22.5-39l12-17q7-19 28-31.5t44-12.5h240q24 0 45 12.5t28 31.5q49 64 49 84v576q0 53-37.5 90.5t-90.5 37.5zm-256-448q0-53-37.5-90.5t-90.5-37.5t-90.5 37.5t-37.5 90.5q0 39 22.5 71.5t58.5 46.5l-81 202h256l-81-202q36-14 58.5-46.5t22.5-71.5zm-594-404q7-19 28-31.5t44-12.5h240q24 0 45 12.5t29 31.5l48 84H.38z"/>`),
		g.Group(children),
	)
}