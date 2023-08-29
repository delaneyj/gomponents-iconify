package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightArrow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#427687" d="M116 4H12c-4.4 0-8 3.6-8 8v104c0 4.4 3.6 8 8 8h104c4.4 0 8-3.6 8-8V12c0-4.4-3.6-8-8-8z"/><path fill="#8CAFBF" d="M109.7 4H11.5C7.4 4 4 7.4 4 11.5v97.9c0 4.2 3.4 7.5 7.5 7.5h98.1c4.2 0 7.5-3.4 7.5-7.5V11.5c.2-4.1-3.3-7.5-7.4-7.5z"/><path fill="#FAFAFA" d="m80.6 37l24 24c1.8 1.6 1.8 4.4 0 6l-24 24c-2.6 2.3-6.6.4-6.6-3V74c0-1.1-.9-2-2-2H26c-2.2 0-4-1.8-4-4v-8c0-2.2 1.8-4 4-4h46c1.1 0 2-.9 2-2V40c0-3.4 4.1-5.3 6.6-3z"/><path fill="#B4E1ED" d="M39.7 12.9c0-2.3-1.6-3-10.8-2.7c-7.7.3-11.5 1.2-13.8 4s-2.9 8.5-3 15.3c0 4.8 0 9.3 2.5 9.3c3.4 0 3.4-7.9 6.2-12.3c5.4-8.7 18.9-10.6 18.9-13.6z" opacity=".5"/>`),
		g.Group(children),
	)
}