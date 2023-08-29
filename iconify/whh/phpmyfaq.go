package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phpmyfaq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M384 832V704q58 0 107.5-26.5t81-69.5t49.5-95t18-104q0-90-68.5-153.5T405 192q-72 0-142 36t-107 92L71 167q61-75 155.5-121T416 0q96 0 177 47t128 128t47 177q0 82-33 169t-85.5 155T526 788t-142 44zm-64-67v63q-87-10-150-48T70.5 678.5T17 535T0 352q0-26 7-57l121 121q1 92 20 164.5T211 704t109 61zm316 67q38 0 57 28t7.5 68t-47 68t-73.5 28t-57-28t-7.5-68t47-68t73.5-28z"/>`),
		g.Group(children),
	)
}