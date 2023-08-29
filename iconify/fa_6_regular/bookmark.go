package fa_6_regular

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 48C0 21.5 21.5 0 48 0v441.4l130.1-92.9c8.3-6 19.6-6 27.9 0l130 92.9V48H48V0h288c26.5 0 48 21.5 48 48v440c0 9-5 17.2-13 21.3s-17.6 3.4-24.9-1.8L192 397.5l-154.1 110c-7.3 5.2-16.9 5.9-24.9 1.8S0 497 0 488V48z"/>`),
		g.Group(children),
	)
}