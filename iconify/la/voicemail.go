package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Voicemail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 9c-3.855 0-7 3.145-7 7s3.145 7 7 7h16c3.855 0 7-3.145 7-7s-3.145-7-7-7s-7 3.145-7 7c0 1.957.82 3.727 2.125 5h-6.25A6.975 6.975 0 0 0 15 16c0-3.855-3.145-7-7-7zm0 2c2.773 0 5 2.227 5 5s-2.227 5-5 5s-5-2.227-5-5s2.227-5 5-5zm16 0c2.773 0 5 2.227 5 5s-2.227 5-5 5s-5-2.227-5-5s2.227-5 5-5z"/>`),
		g.Group(children),
	)
}