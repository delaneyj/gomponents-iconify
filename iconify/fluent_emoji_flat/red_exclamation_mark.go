package fluent_emoji_flat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedExclamationMark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#F8312F" d="M13 20.512a2.501 2.501 0 0 0 5 .001V4.487a2.501 2.501 0 0 0-5 0v16.025Zm0 6.988c0 1.375 1.125 2.5 2.5 2.5s2.5-1.125 2.5-2.5s-1.125-2.5-2.5-2.5a2.508 2.508 0 0 0-2.5 2.5Z"/>`),
		g.Group(children),
	)
}