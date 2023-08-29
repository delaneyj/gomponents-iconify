package uil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonEclipse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 2a10 10 0 0 0-2.14.24h-.12a10 10 0 0 0-.1 19.44h.14A9.57 9.57 0 0 0 12 22a10 10 0 0 0 0-20Zm-2 17.74a8 8 0 0 1 0-15.48a8 8 0 0 1 0 15.48Zm4.53-.16a10 10 0 0 0 0-15.16a8 8 0 0 1 0 15.16Z"/>`),
		g.Group(children),
	)
}