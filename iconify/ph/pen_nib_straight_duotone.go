package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenNibStraightDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M215.17 127.43L184 72H72l-31.17 55.43a8 8 0 0 0 .73 8.29L128 248l86.43-112.28a8 8 0 0 0 .74-8.29ZM128 152a20 20 0 1 1 20-20a20 20 0 0 1-20 20Z" opacity=".2"/><path d="M222.33 123.89c-.06-.13-.12-.26-.19-.38L192 69.9V32a16 16 0 0 0-16-16H80a16 16 0 0 0-16 16v37.92l-30.14 53.59c-.07.12-.13.25-.2.38a15.94 15.94 0 0 0 1.46 16.57l.11.14l86.44 112.28a8 8 0 0 0 12.67 0l86.43-112.28l.11-.14a15.92 15.92 0 0 0 1.45-16.57ZM176 32v32H80V32Zm-48 112a12 12 0 1 1 12-12a12 12 0 0 1-12 12Zm8 80.5v-65.67a28 28 0 1 0-16 0v65.66L48 131l28.69-51h102.63L208 131Z"/></g>`),
		g.Group(children),
	)
}