package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RailMetro(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M5.5 0s-.75 0-1 1L3 6.5V10c0 1 1 1 1 1h7s1 0 1-1V6.5L10.5 1c-.273-1-1-1-1-1h-4zm1 1.5h2s.536 0 .75 1L10 6c.215 1.002-1 1-1 1H6s-1.215.002-1-1l.75-3.5c.214-1 .75-1 .75-1zM5 8a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm1.75 0h1.5a.25.25 0 1 1 0 .5h-1.5a.25.25 0 1 1 0-.5zM10 8a1 1 0 1 1 0 2a1 1 0 0 1 0-2zm-5.875 4L3 15h1.5l.375-1h5.25l.375 1H12l-1.125-3h-1.5l.375 1h-4.5l.375-1h-1.5z"/>`),
		g.Group(children),
	)
}