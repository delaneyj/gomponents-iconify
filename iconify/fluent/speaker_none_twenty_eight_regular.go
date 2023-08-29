package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerNoneTwentyEightRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M14.395 3.902c.798-.748 2.105-.182 2.105.912v18.37c0 1.095-1.306 1.66-2.105.912L9.458 19.47a1.75 1.75 0 0 0-1.196-.473H5.25A3.25 3.25 0 0 1 2 15.747v-3.492a3.25 3.25 0 0 1 3.25-3.25h3.011c.445 0 .873-.169 1.197-.473l4.937-4.63zM15 5.392l-4.516 4.234a3.25 3.25 0 0 1-2.223.88H5.25a1.75 1.75 0 0 0-1.75 1.75v3.491c0 .967.784 1.75 1.75 1.75h3.012c.825 0 1.62.314 2.222.879L15 22.607V5.391z" fill="currentColor"/><path d="M19.782 10.722a.75.75 0 0 0-1.064 1.056l2.218 2.235l-2.215 2.206a.75.75 0 1 0 1.058 1.062l2.218-2.207l2.225 2.208a.75.75 0 0 0 1.056-1.064l-2.22-2.205l2.224-2.234a.75.75 0 1 0-1.063-1.058l-2.223 2.231l-2.214-2.23z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}