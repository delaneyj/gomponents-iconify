package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wrench(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M15.671 12.779L8.475 6.611A4.5 4.5 0 0 0 3.193.193l2.6 2.6a1.003 1.003 0 0 1 0 1.414L4.207 5.793a1.003 1.003 0 0 1-1.414 0l-2.6-2.6a4.5 4.5 0 0 0 6.418 5.282l6.168 7.196a.914.914 0 0 0 1.358.052l1.586-1.586a.914.914 0 0 0-.052-1.358z"/>`),
		g.Group(children),
	)
}