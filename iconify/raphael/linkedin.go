package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Linkedin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M27.25 3.125h-22a2 2 0 0 0-2 2v22a2 2 0 0 0 2 2h22a2 2 0 0 0 2-2v-22a2 2 0 0 0-2-2zM11.22 26.78h-4v-14h4v14zm-2-15.5c-1.384 0-2.5-1.118-2.5-2.5s1.116-2.5 2.5-2.5s2.5 1.12 2.5 2.5s-1.118 2.5-2.5 2.5zm16 15.5h-4v-8.5c0-.4-.404-1.054-.688-1.212c-.375-.21-1.26-.23-1.665-.034l-1.648.793v8.954h-4v-14h4v.615c1.582-.723 3.78-.652 5.27.184c1.58.885 2.73 2.863 2.73 4.7v8.5z"/>`),
		g.Group(children),
	)
}