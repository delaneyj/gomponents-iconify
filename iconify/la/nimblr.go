package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nimblr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M7 1v18h.025c-.008.166-.025.331-.025.5c0 5.238 4.262 9.5 9.5 9.5s9.5-4.262 9.5-9.5s-4.262-9.5-9.5-9.5a9.467 9.467 0 0 0-7.066 3.172L7 1zm9.5 11c4.136 0 7.5 3.364 7.5 7.5S20.636 27 16.5 27S9 23.636 9 19.5s3.364-7.5 7.5-7.5zm-3 5a1.5 1.5 0 0 0 0 3a1.5 1.5 0 0 0 0-3zm6 0a1.5 1.5 0 0 0 0 3a1.5 1.5 0 0 0 0-3z"/>`),
		g.Group(children),
	)
}