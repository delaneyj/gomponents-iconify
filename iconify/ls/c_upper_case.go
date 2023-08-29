package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CUpperCase(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="m721 160l-63 44C600 125 505 73 397 73s-204 52-262 131c-38 51-60 115-60 184s22 132 60 183c58 79 154 130 262 130s203-51 261-130l60 41c-58 80-148 139-254 157c-155 28-305-37-391-155c-33-45-57-99-67-158c-19-106 10-211 70-294C135 83 225 25 330 7c156-28 306 36 391 153z"/>`),
		g.Group(children),
	)
}