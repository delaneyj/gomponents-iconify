package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Settings(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g transform="translate(10 10)"><path id="oouiSettings0" fill="currentColor" d="M1.5-10h-3l-1 6.5h5m0 7h-5l1 6.5h3"/><use href="#oouiSettings0" transform="rotate(45)"/><use href="#oouiSettings0" transform="rotate(90)"/><use href="#oouiSettings0" transform="rotate(135)"/></g><path fill="currentColor" d="M10 2.5a7.5 7.5 0 0 0 0 15a7.5 7.5 0 0 0 0-15v4a3.5 3.5 0 0 1 0 7a3.5 3.5 0 0 1 0-7"/>`),
		g.Group(children),
	)
}