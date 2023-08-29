package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserHappy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.5 23s1 1 2.5 1s2.5-1 2.5-1m4.5-5a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm-12 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Zm18-2c0 7.18-5.82 13-13 13S3 23.18 3 16S8.82 3 16 3s13 5.82 13 13Z"/>`),
		g.Group(children),
	)
}