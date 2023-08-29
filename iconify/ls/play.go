package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Play(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 49v605c0 11 5 20 15 26c4 3 11 4 16 4s9-1 14-4l524-302c10-7 16-16 16-27c0-10-6-19-16-26L45 23C27 10 0 26 0 49z"/>`),
		g.Group(children),
	)
}