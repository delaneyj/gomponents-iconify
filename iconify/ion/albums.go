package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Albums(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M368 96H144a16 16 0 0 1 0-32h224a16 16 0 0 1 0 32Zm32 48H112a16 16 0 0 1 0-32h288a16 16 0 0 1 0 32Zm19.13 304H92.87A44.92 44.92 0 0 1 48 403.13V204.87A44.92 44.92 0 0 1 92.87 160h326.26A44.92 44.92 0 0 1 464 204.87v198.26A44.92 44.92 0 0 1 419.13 448Z"/>`),
		g.Group(children),
	)
}