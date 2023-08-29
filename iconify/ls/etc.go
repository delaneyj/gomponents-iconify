package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Etc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M176 347c0-48-40-88-88-88c-49 0-88 40-88 88s39 88 88 88c48 0 88-40 88-88zm270 0c0-48-40-88-88-88s-88 40-88 88s40 88 88 88s88-40 88-88zm271 0c0-48-39-88-88-88c-48 0-88 40-88 88s40 88 88 88c49 0 88-40 88-88z"/>`),
		g.Group(children),
	)
}