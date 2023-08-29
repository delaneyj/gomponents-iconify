package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Timeout(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7.774 5.215L6.326 3.767A9.995 9.995 0 0 1 20.23 17.67l-1.445-1.445a7.99 7.99 0 0 0-11.01-11.01Zm12.588 15.149l-1.279 1.279l-1.413-1.414A9.995 9.995 0 0 1 3.768 6.327l-1.41-1.41l1.278-1.28l1.29 1.292l1.417 1.413L11 11.002V7h1.5v5.25l4.5 2.67l-.75 1.23l-.254-.152Zm-4.136-1.58l-5.793-5.792l-5.218-5.218a7.99 7.99 0 0 0 11.01 11.01Z"/>`),
		g.Group(children),
	)
}