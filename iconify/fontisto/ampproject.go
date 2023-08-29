package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ampproject(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.96.66L13.62 9.9h2.16c.6 0 .84.48.48 1.08L8.58 23.52a12.42 12.42 0 0 0 3.418.48H12c6.628-.005 11.998-5.379 11.998-12.008c0-5.209-3.317-9.644-7.955-11.306L15.959.66zm-5.58 13.38H8.22c-.6 0-.84-.48-.48-1.08L15.42.48A12.42 12.42 0 0 0 12.002 0H12C5.372.005.002 5.379.002 12.008c0 5.209 3.317 9.644 7.955 11.306l.084.026z"/>`),
		g.Group(children),
	)
}