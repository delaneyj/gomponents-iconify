package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RubyOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M14.5 14.5v.5h.5v-.5h-.5Zm0-14h.5V0h-.5v.5Zm-6 0V0h-.207l-.147.146L8.5.5Zm-8 8l-.354-.354L0 8.293V8.5h.5Zm0 6H0v.5h.5v-.5Zm4-4l-.224.447l.019.01l.02.007l.185-.464Zm0-6V4a.5.5 0 0 0-.5.5h.5Zm6 0l.464-.186l-.008-.019l-.009-.019l-.447.224Zm4.5 10V.5h-1v14h1ZM14.5 0h-6v1h6V0ZM8.146.146l-8 8l.708.708l8-8l-.708-.708ZM0 8.5v6h1v-6H0ZM.5 15h14v-1H.5v1ZM14.146.146l-14 14l.708.708l14-14l-.708-.708ZM5 10.5v-6H4v6h1ZM4.5 5h6V4h-6v1Zm-.186 5.964l10 4l.372-.928l-10-4l-.372.928Zm5.722-6.278l4 10l.928-.372l-4-10l-.928.372ZM8.053.724l2 4l.894-.448l-2-4l-.894.448ZM.276 8.947l4 2l.448-.894l-4-2l-.448.894Z"/>`),
		g.Group(children),
	)
}