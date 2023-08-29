package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hackerrank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15.998 3C14.225 3 5.535 7.984 4.658 9.504c-.878 1.52-.878 11.479 0 12.994C5.54 24.017 14.23 29 15.998 29c1.764 0 10.454-4.98 11.338-6.498c.887-1.522.887-11.488 0-13.004v-.002C26.445 7.98 17.756 3 15.998 3zm-.002 2.012c1.697.353 8.422 4.205 9.6 5.498c.54 1.64.539 9.334 0 10.978c-1.17 1.291-7.9 5.148-9.6 5.5c-1.698-.35-8.424-4.205-9.596-5.498c-.535-1.648-.535-9.334 0-10.98c1.17-1.293 7.897-5.148 9.596-5.498zM13 9l-2 2h1v10h2v-4h4v4h-1l2 2l2-2h-1v-9h-2v3h-4v-4h1l-2-2z"/>`),
		g.Group(children),
	)
}