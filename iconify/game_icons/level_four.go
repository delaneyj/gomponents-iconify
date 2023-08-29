package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LevelFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M16 16c48 96 24 120 72 168c0 0 44.864 0 69.53-21.688A135.765 135.765 0 0 0 120.25 256a135.765 135.765 0 0 0 271.5 0a135.765 135.765 0 0 0-36.375-92.906C380.22 183.997 424 184 424 184c48-48 24-72 72-168c-96 48-120 24-168 72c0 0 .004 44.6 21.5 69.313a135.765 135.765 0 0 0-186.72-.344C183.987 132.19 184 88 184 88c-48-48-72-24-168-72zm72 312c-48 48-24 72-72 168c96-48 120-24 168-72c0 0 0-48-24-72s-72-24-72-24zm336 0s-48 0-72 24s-24 72-24 72c48 48 72 24 168 72c-48-96-24-120-72-168z"/>`),
		g.Group(children),
	)
}