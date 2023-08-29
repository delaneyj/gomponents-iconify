package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Phpbbalt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 512H256L0 320L512 0l512 320zm0 64l224-160l-96 416l-192 192h-64l64-320H320l64 320h-64L128 832L32 416l224 160h512z"/>`),
		g.Group(children),
	)
}