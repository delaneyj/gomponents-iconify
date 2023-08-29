package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AndroidPhoneSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M11.29 17.71a1 1 0 0 0 1.42 0a1.15 1.15 0 0 0 .21-.33a.94.94 0 0 0 0-.76a1.15 1.15 0 0 0-.21-.33a1 1 0 0 0-1.42 0a1.15 1.15 0 0 0-.21.33a.94.94 0 0 0 0 .76a1.15 1.15 0 0 0 .21.33ZM8.66 4H16a1 1 0 0 1 1 1v7.34a1 1 0 0 0 2 0V5a3 3 0 0 0-3-3H8.66a1 1 0 0 0 0 2Zm13.05 16.29l-18-18a1 1 0 0 0-1.42 1.42L5 6.41V19a3 3 0 0 0 3 3h8a3 3 0 0 0 2.76-1.83l1.53 1.54a1 1 0 0 0 1.42 0a1 1 0 0 0 0-1.42ZM17 19a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1V8.41l10 10Z"/>`),
		g.Group(children),
	)
}