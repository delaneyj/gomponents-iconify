package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckBoxWithCheck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#427687" d="M116 4H12c-4.4 0-8 3.6-8 8v104c0 4.4 3.6 8 8 8h104c4.4 0 8-3.6 8-8V12c0-4.4-3.6-8-8-8z"/><path fill="#8CAFBF" d="M109.7 4H11.5C7.4 4 4 7.4 4 11.5v97.9c0 4.2 3.4 7.5 7.5 7.5h98.1c4.2 0 7.5-3.4 7.5-7.5V11.5c.2-4.1-3.3-7.5-7.4-7.5z"/><path fill="#B4E1ED" d="M39.7 12.9c0-2.3-1.6-3-10.8-2.7c-7.7.3-11.5 1.2-13.8 4s-2.9 8.5-3 15.3c0 4.8 0 9.3 2.5 9.3c3.4 0 3.4-7.9 6.2-12.3c5.4-8.7 18.9-10.6 18.9-13.6z" opacity=".5"/><path fill="#FBF9F9" d="M38.8 105.5a8.862 8.862 0 0 1-8.7-6.2L15 53.8c-1.46-4.81 1.26-9.9 6.07-11.36c4.65-1.41 9.59 1.08 11.23 5.66l9.8 29.5L89.2 18c3.12-3.95 8.85-4.62 12.8-1.5s4.62 8.85 1.5 12.8L45.9 102a9.086 9.086 0 0 1-7.1 3.5z"/>`),
		g.Group(children),
	)
}