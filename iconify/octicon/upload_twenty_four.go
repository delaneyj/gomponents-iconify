package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4 20.25V18a.75.75 0 0 1 1.5 0v2.25c0 .138.112.25.25.25h12.5a.25.25 0 0 0 .25-.25V18a.75.75 0 0 1 1.5 0v2.25A1.75 1.75 0 0 1 18.25 22H5.75A1.75 1.75 0 0 1 4 20.25Z"/><path fill="currentColor" d="M5.22 9.53a.749.749 0 0 1 0-1.06l6.25-6.25a.749.749 0 0 1 1.06 0l6.25 6.25a.749.749 0 1 1-1.06 1.06l-4.97-4.969V16.75a.75.75 0 0 1-1.5 0V4.561L6.28 9.53a.749.749 0 0 1-1.06 0Z"/>`),
		g.Group(children),
	)
}