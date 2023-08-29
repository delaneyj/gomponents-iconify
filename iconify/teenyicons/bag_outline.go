package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BagOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="m2.401 6.39l-.497-.056l.497.056Zm-.778 7l.497.055l-.497-.055Zm11.754 0l-.497.055l.497-.055Zm-.778-7l.497-.056l-.497.056ZM1.904 6.334l-.778 7l.994.11l.778-7l-.994-.11ZM2.617 15h9.766v-1H2.617v1Zm11.257-1.666l-.778-7l-.994.11l.778 7l.993-.11ZM11.604 5H3.396v1h8.21V5Zm1.492 1.334A1.5 1.5 0 0 0 11.605 5v1a.5.5 0 0 1 .497.445l.994-.11ZM12.383 15a1.5 1.5 0 0 0 1.49-1.666l-.993.11a.5.5 0 0 1-.497.556v1ZM1.126 13.334A1.5 1.5 0 0 0 2.617 15v-1a.5.5 0 0 1-.497-.555l-.994-.11Zm1.772-6.89A.5.5 0 0 1 3.395 6V5a1.5 1.5 0 0 0-1.49 1.334l.993.11ZM5 4v-.5H4V4h1Zm5-.5V4h1v-.5h-1ZM7.5 1A2.5 2.5 0 0 1 10 3.5h1A3.5 3.5 0 0 0 7.5 0v1ZM5 3.5A2.5 2.5 0 0 1 7.5 1V0A3.5 3.5 0 0 0 4 3.5h1Z"/>`),
		g.Group(children),
	)
}