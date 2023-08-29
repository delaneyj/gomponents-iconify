package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoWikidata(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M0 4v12.258h.742V4zm1.482 0v12.258h2.223V4zm2.96 0v12.258H6.67V4zm2.964 0v12.258h.744V4zm1.48 0v12.258h.745V4zm1.483 0v12.258h2.224V4zm2.962 0v12.258h.742V4zm1.482 0v12.258h2.223V4zm2.96 0v12.258h.744V4zm1.484 0v12.258H20V4z"/>`),
		g.Group(children),
	)
}