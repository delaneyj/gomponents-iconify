package basil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SkipNextOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.607 10.748c.82.634.82 1.87 0 2.505a25.758 25.758 0 0 1-7.143 3.9l-.466.166c-1.023.364-2.1-.329-2.238-1.381c-.34-2.59-.34-5.286 0-7.876c.138-1.052 1.215-1.745 2.238-1.381l.466.165a25.76 25.76 0 0 1 7.143 3.902Zm-.918 1.318a.084.084 0 0 0 0-.132A24.257 24.257 0 0 0 6.962 8.26l-.466-.166a.194.194 0 0 0-.249.163a29.063 29.063 0 0 0 0 7.486c.017.13.15.198.25.163l.465-.166c2.423-.86 4.694-2.1 6.727-3.674ZM18 6.25a.75.75 0 0 1 .75.75v10a.75.75 0 0 1-1.5 0V7a.75.75 0 0 1 .75-.75Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}