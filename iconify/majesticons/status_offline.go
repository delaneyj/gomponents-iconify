package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatusOffline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill-rule="evenodd" clip-rule="evenodd" d="M3.707 2.293a1 1 0 0 0-1.414 1.414L15.535 16.95l2.829 2.828l1.929 1.93a1 1 0 0 0 1.414-1.415l-1.253-1.254c3.607-4.321 3.382-10.76-.676-14.817a1 1 0 1 0-1.414 1.414a9.001 9.001 0 0 1 .668 11.982l-1.425-1.425a7.002 7.002 0 0 0-.657-9.143a1 1 0 1 0-1.414 1.414a5.002 5.002 0 0 1 .636 6.294l-1.465-1.465a3 3 0 0 0-4-4l-7-7zM3.75 8.4a1 1 0 0 0-1.834-.8C.161 11.624.928 16.485 4.222 19.778a1 1 0 0 0 1.414-1.414A9.004 9.004 0 0 1 3.749 8.4zm3.32 2.766a1 1 0 0 0-1.972-.332A6.992 6.992 0 0 0 7.05 16.95a1 1 0 1 0 1.414-1.414a4.993 4.993 0 0 1-1.394-4.37z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}