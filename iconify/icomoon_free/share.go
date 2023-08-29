package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Share(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4 10s.919-3 6-3v3l6-4l-6-4v3c-4 0-6 2.495-6 5zm7 2H2V6h1.967c.158-.186.327-.365.508-.534A6.933 6.933 0 0 1 6.914 4H0v10h13V9.803l-2 1.333V12z"/>`),
		g.Group(children),
	)
}