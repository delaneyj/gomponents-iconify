package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagnetOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 3a2 2 0 0 1 2 2m0 4v4a3 3 0 0 0 5.552 1.578M15 11V5a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v8a7.99 7.99 0 0 1-.424 2.577m-1.463 2.584A8 8 0 0 1 4 13V5c0-.297.065-.58.181-.833M4 8h4m7 0h4M3 3l18 18"/>`),
		g.Group(children),
	)
}