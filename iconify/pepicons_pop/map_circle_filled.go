package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MapCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopMapCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M13 13a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0-3a.5.5 0 1 1 0 1a.5.5 0 0 1 0-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8 10.286C8 13.199 10.805 18 13 18s5-4.801 5-7.714C18 7.38 15.777 5 13 5s-5 2.381-5 5.286Zm8 0C16 12.24 13.803 16 13 16c-.803 0-3-3.76-3-5.714C10 8.456 11.36 7 13 7s3 1.456 3 3.286Z" clip-rule="evenodd"/><path d="M16.73 13a1 1 0 1 1 0-2a3 3 0 0 1 2.89 2.197l1.39 5A3 3 0 0 1 18.118 22H7.88a3 3 0 0 1-2.89-3.803l1.389-5A3 3 0 0 1 9.27 11a1 1 0 0 1 0 2a1 1 0 0 0-.963.732l-1.39 5A1 1 0 0 0 7.882 20H18.12a1 1 0 0 0 .963-1.268l-1.389-5A1 1 0 0 0 16.73 13Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopMapCircleFilled0)"/></g>`),
		g.Group(children),
	)
}