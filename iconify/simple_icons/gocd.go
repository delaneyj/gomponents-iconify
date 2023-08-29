package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gocd(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18.043 13.237l-8.907 5.935a1.47 1.47 0 0 1-.823.25a1.449 1.449 0 0 1-.696-.173a1.48 1.48 0 0 1-.784-1.308V12a1.482 1.482 0 1 1 2.957 0v3.163L14.537 12L7.478 7.304A1.49 1.49 0 1 1 9.13 4.823l8.913 5.94a1.492 1.492 0 0 1 0 2.474M12 0a12 12 0 1 0 12 12A12.012 12.012 0 0 0 12 0"/>`),
		g.Group(children),
	)
}