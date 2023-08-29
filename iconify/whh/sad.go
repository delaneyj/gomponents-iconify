package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sad(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zM288 256q-40 0-68 37.5T192 384t28 90.5t68 37.5t68-37.5t28-90.5t-28-90.5t-68-37.5zm448 0q-40 0-68 37.5T640 384t28 90.5t68 37.5t68-37.5t28-90.5t-28-90.5t-68-37.5zm93 507q-22-58-107.5-90.5T512 640t-209.5 32.5T195 763l-3 5q0 13 9.5 22.5T224 800t22.5-9.5T256 768q53-41 102.5-52.5T512 704t153.5 11.5T768 768q0 13 9.5 22.5T800 800t22.5-9.5T832 768q-1-2-3-5z"/>`),
		g.Group(children),
	)
}