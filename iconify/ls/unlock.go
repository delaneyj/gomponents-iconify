package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M468 332h67c9 0 16 8 16 17v352c0 9-7 16-16 16H16c-9 0-16-7-16-16V349c0-9 7-17 16-17h368V176c0-51-42-94-94-94h-31c-51 0-93 43-93 94v59H81v-67C81 76 140 0 231 0h86c92 0 151 76 151 168v164zM228 615h94l-23-101c16-8 27-25 27-44c0-29-23-51-51-51s-50 22-50 51c0 19 11 36 27 44z"/>`),
		g.Group(children),
	)
}