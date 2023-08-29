package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stopwatch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.414 8.902a.993.993 0 0 0 .293-.195l.5-.5a.999.999 0 1 0-1.414-1.414l-.5.5l-.115.173a8.969 8.969 0 0 0-5.189-2.41L13 5V4h1c.55 0 1-.45 1-1s-.45-1-1-1h-4c-.55 0-1 .45-1 1s.45 1 1 1h1v1l.012.057A9 9 0 0 0 12 23a9 9 0 0 0 7.414-14.098zM12 21c-3.859 0-7-3.14-7-7s3.141-7 7-7s7 3.14 7 7s-3.141 7-7 7zm1-8v-2c0-.55-.45-1-1-1s-1 .45-1 1v3c0 .55.45 1 1 1h3c.55 0 1-.45 1-1s-.45-1-1-1h-2zm-1-5c-3.312 0-6 2.688-6 6s2.688 6 6 6s6-2.688 6-6s-2.688-6-6-6zm0 11c-2.757 0-5-2.243-5-5s2.243-5 5-5s5 2.243 5 5s-2.243 5-5 5z"/>`),
		g.Group(children),
	)
}