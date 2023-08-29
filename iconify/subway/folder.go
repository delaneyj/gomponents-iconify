package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M490.7 96H277.3l-10.7-42.7c-1.5-9.2-12-21.3-21.3-21.3H96c-9.3 0-19.8 12.2-21.3 21.3L64 96H21.3C12.1 96 0 108.1 0 117.3v341.3C0 468 12.1 480 21.3 480h469.3c9.3 0 21.3-12 21.3-21.3V117.3C512 108.1 500 96 490.7 96z"/>`),
		g.Group(children),
	)
}