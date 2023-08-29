package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftArrowCurvingRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#427687" d="M116 4H12c-4.4 0-8 3.6-8 8v104c0 4.4 3.6 8 8 8h104c4.4 0 8-3.6 8-8V12c0-4.4-3.6-8-8-8z"/><path fill="#8CAFBF" d="M109.7 4H11.5C7.4 4 4 7.4 4 11.5v97.9c0 4.2 3.4 7.5 7.5 7.5h98.1c4.2 0 7.5-3.4 7.5-7.5V11.5c.2-4.1-3.3-7.5-7.4-7.5z"/><path fill="#FAFAFA" d="M44 22h32c2.2 0 4 1.8 4 4v8c0 2.2-1.8 4-4 4H44c-6.5 0-12 7.3-12 16s5.5 16 12 16h34c1.1 0 2-.9 2-2V54c0-3.4 4.1-5.3 6.6-3l24 24c1.8 1.6 1.8 4.4 0 6l-24 24c-2.6 2.3-6.6.4-6.6-3V88c0-1.1-.9-2-2-2H44c-15.5 0-28-14.3-28-32s12.5-32 28-32z"/><path fill="#B4E1ED" d="M39.7 12.9c0-2.3-1.6-3-10.8-2.7c-7.7.3-11.5 1.2-13.8 4s-2.9 8.5-3 15.3c0 4.8 0 9.3 2.5 9.3c3.4 0 3.4-7.9 6.2-12.3c5.4-8.7 18.9-10.6 18.9-13.6z" opacity=".5"/>`),
		g.Group(children),
	)
}