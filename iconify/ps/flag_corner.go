package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagCorner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M48 384h107v64h330l-72-171l72-170H325V43H48V21q0-9-6-15T27 0Q17 0 11 6T5 21v491h43V384zm277-235h96l-55 128l55 128H197v-21h128V149zM48 235V85h235v256H48V235z"/>`),
		g.Group(children),
	)
}