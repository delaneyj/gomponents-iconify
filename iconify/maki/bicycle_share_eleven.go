package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BicycleShareEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M7 1a1 1 0 1 1 1 1a1 1 0 0 1-1-1zM1.973 7.1a1.5 1.5 0 0 1 1.654.408a.5.5 0 0 0 .749-.663A2.519 2.519 0 0 0 2.363 6a2.5 2.5 0 1 0 2.008 4.158a.5.5 0 0 0-.748-.658a1.5 1.5 0 1 1-1.65-2.4zM7.1 4.8a.5.5 0 0 0 .4.2h2a.5.5 0 1 0 .014-1H7.75L6.4 2.2a.5.5 0 0 0-.386-.2a.506.506 0 0 0-.368.147l-2 2a.5.5 0 0 0 0 .707L5 6.207V8.5a.5.5 0 0 0 1 .015V6a.505.505 0 0 0-.144-.353L5.1 4.895l1.165-1.2zm2.287 1.362A2.526 2.526 0 0 0 8.643 6a2.525 2.525 0 0 0-2.014.838a.5.5 0 0 0 .75.664a1.5 1.5 0 1 1 0 1.992a.5.5 0 0 0-.78.627a.596.596 0 0 0 .034.037a2.5 2.5 0 1 0 2.752-4z" fill="currentColor"/>`),
		g.Group(children),
	)
}