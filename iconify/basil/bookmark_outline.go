package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookmarkOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.924 4.418a23.778 23.778 0 0 0-7.848 0a1.194 1.194 0 0 0-.973.94a35.64 35.64 0 0 0-.142 13.327l.163.913l3.67-3.487a1.75 1.75 0 0 1 2.411 0l3.67 3.487l.164-.913a35.64 35.64 0 0 0-.142-13.327a1.194 1.194 0 0 0-.973-.94ZM7.828 2.94a25.278 25.278 0 0 1 8.344 0a2.694 2.694 0 0 1 2.195 2.123c.923 4.579.973 9.29.149 13.888l-.222 1.235a1.323 1.323 0 0 1-2.213.726l-3.909-3.713a.25.25 0 0 0-.344 0l-3.909 3.713a1.323 1.323 0 0 1-2.213-.726l-.222-1.235a37.14 37.14 0 0 1 .148-13.888A2.694 2.694 0 0 1 7.828 2.94Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}