package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Happy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zM288 256q-40 0-68 37.5T192 384t28 90.5t68 37.5t68-37.5t28-90.5t-28-90.5t-68-37.5zm448 0q-40 0-68 37.5T640 384t28 90.5t68 37.5t68-37.5t28-90.5t-28-90.5t-68-37.5zm64 416q-13 0-22.5 9.5T768 704h-1q-53 41-102.5 52.5t-153 11.5t-153-11.5T256 704q0-13-9.5-22.5T224 672t-22.5 9.5T192 704q1 2 3 5q23 58 108 90.5T511 832h1q118 0 201.5-30T824 718l8-14q0-13-9.5-22.5T800 672z"/>`),
		g.Group(children),
	)
}