package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SwitchCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M21.125 5H19.2l-1.304-2.301C17.67 2.3 17.18 2 16.5 2h-7c-.709 0-1.169.3-1.396.699L6.8 5H4.875A4.874 4.874 0 0 0 0 9.875v8.25A4.874 4.874 0 0 0 4.875 23h16.25A4.874 4.874 0 0 0 26 18.125v-8.25A4.874 4.874 0 0 0 21.125 5zM13 21a7 7 0 0 1-7-7H4l3-4l3 4H8c0 2.757 2.243 5 5 5c1.268 0 2.415-.49 3.297-1.271l1.214 1.619A6.968 6.968 0 0 1 13 21zm6-3l-3-4h2c0-2.757-2.243-5-5-5c-1.268 0-2.415.49-3.297 1.271L8.489 8.652A6.968 6.968 0 0 1 13 7a7 7 0 0 1 7 7h2l-3 4z"/>`),
		g.Group(children),
	)
}