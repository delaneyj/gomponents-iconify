package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GestureRanslation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9.865 2.186a2.57 2.57 0 0 1 5.092.068a2.57 2.57 0 0 1 3.158 2.502v.83a2.567 2.567 0 0 1 3.134 2.504V15a8 8 0 0 1-8 8h-1.35a8 8 0 0 1-7.738-5.968l-1.269-4.829A2.296 2.296 0 0 1 6.698 9.96V4.755a2.57 2.57 0 0 1 2.57-2.569h.597Zm-.028 2.57a.57.57 0 1 0-1.139 0v8.748c0 1.072-1.432 1.374-1.89.463l-1.44-2.495a.296.296 0 0 0-.541.223l1.268 4.83A6 6 0 0 0 11.898 21h1.35a6 6 0 0 0 6-6V8.09a.567.567 0 1 0-1.134 0V12h-2V4.755a.57.57 0 1 0-1.138 0V12h-2V2.57a.57.57 0 0 0-1.139 0V12h-2V4.755Z"/>`),
		g.Group(children),
	)
}