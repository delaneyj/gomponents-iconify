package topcoat

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Refresh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 42 42"),
		g.Raw(`<path fill="currentColor" d="M26.731 17.5H39.5V4.91l-4.335 4.385c-.34-.399-.651-.79-1.021-1.17c-7.48-7.5-19.572-7.5-27.042 0c-7.479 7.5-7.465 19.661.015 27.161c7.471 7.5 19.597 7.5 27.078 0c0 0 2.505-2.939 3.334-4.51l-4.129-1.84c-.609 1.068-2.39 3.148-2.39 3.148c-5.719 5.74-14.979 5.74-20.689 0c-5.72-5.729-5.719-15.031 0-20.761c5.71-5.74 14.971-5.74 20.689 0c.381.38.729.574 1.061.984l-5.34 5.193z"/>`),
		g.Group(children),
	)
}