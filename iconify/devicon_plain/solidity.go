package devicon_plain

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Solidity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="currentColor" d="M43.322 0L22.756 36.576l20.566 36.561l20.564-36.561h41.143L84.465 0H43.322zm41.342 54.863L64.1 91.424H22.955L43.519 128h41.145l20.58-36.576l-20.58-36.561z"/>`),
		g.Group(children),
	)
}