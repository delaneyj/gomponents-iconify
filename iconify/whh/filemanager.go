package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Filemanager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024.526 768q0 53-37.5 90.5t-90.5 37.5h-768q-53 0-90.5-37.5T.526 768V256q0-26 18.5-45t45.5-19h480q0-12 11.5-33t22.5-36l12-16q7-18 28-30.5t44-12.5h240q24 0 45 12.5t28 30.5q49 58 49 85v576zm-978-725q7-18 28-30.5t44-12.5h240q24 0 45 12.5t29 30.5l48 85h-480z"/>`),
		g.Group(children),
	)
}