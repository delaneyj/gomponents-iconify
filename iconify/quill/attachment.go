package quill

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Attachment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10.49 19.182l7.778-7.778c.976-.976 2.382-1.153 3.359-.177c.976.976.8 2.383-.177 3.359l-8.132 8.132c-1.414 1.414-4.243 1.414-6.01-.354c-1.768-1.768-1.768-4.596-.354-6.01l8.132-8.132a6.5 6.5 0 0 1 9.192 9.192l-4.596 4.596"/>`),
		g.Group(children),
	)
}