package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckReadBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4.565 12.407a.75.75 0 1 0-1.13.986l1.13-.986ZM7.143 16.5l-.565.493a.75.75 0 0 0 1.13 0l-.565-.493Zm8.422-8.507a.75.75 0 1 0-1.13-.986l1.13.986Zm-5.059 3.514a.75.75 0 0 0 1.13.986l-1.13-.986Zm-.834 3.236a.75.75 0 1 0-1.13-.986l1.13.986Zm-6.237-1.35l3.143 3.6l1.13-.986l-3.143-3.6l-1.13.986Zm4.273 3.6l1.964-2.25l-1.13-.986l-1.964 2.25l1.13.986Zm3.928-4.5l1.965-2.25l-1.13-.986l-1.965 2.25l1.13.986Zm1.965-2.25l1.964-2.25l-1.13-.986l-1.964 2.25l1.13.986Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20 7.563l-4.286 4.5M11 16l.429.563l2.143-2.25"/></g>`),
		g.Group(children),
	)
}