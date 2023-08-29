package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Forgejo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M16.777 0a2.9 2.9 0 1 1-2.529 4.322H12.91a4.266 4.266 0 0 0-4.265 4.195v2.118a7.076 7.076 0 0 1 4.147-1.42l.118-.002h1.338a2.9 2.9 0 0 1 5.43 1.422a2.9 2.9 0 0 1-5.43 1.422H12.91a4.266 4.266 0 0 0-4.265 4.195v2.319A2.9 2.9 0 0 1 7.222 24A2.9 2.9 0 0 1 5.8 18.57V8.589a7.109 7.109 0 0 1 6.991-7.108l.118-.001h1.338A2.9 2.9 0 0 1 16.778 0ZM7.223 19.905a1.194 1.194 0 1 0 0 2.389a1.194 1.194 0 0 0 0-2.389Zm9.554-10.464a1.194 1.194 0 1 0 0 2.389a1.194 1.194 0 0 0 0-2.39Zm0-7.735a1.194 1.194 0 1 0 0 2.389a1.194 1.194 0 0 0 0-2.389Z"/>`),
		g.Group(children),
	)
}