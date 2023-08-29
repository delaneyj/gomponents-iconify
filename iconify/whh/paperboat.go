package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Paperboat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1011 337L802 653q-12 21-34 35.5T726 703H300q-20 0-42-14.5T224 653L15 337q-18-33-13-57t32-24h197L469 18q18-18 43.5-18T556 18l238 238h198q27 0 32 24t-13 57z"/>`),
		g.Group(children),
	)
}