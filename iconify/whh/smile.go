package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Smile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zM288 256q-40 0-68 37.5T192 384t28 90.5t68 37.5t68-37.5t28-90.5t-28-90.5t-68-37.5zm448 0q-40 0-68 37.5T640 384t28 90.5t68 37.5t68-37.5t28-90.5t-28-90.5t-68-37.5zm64 384H224q-13 0-22.5 9.5T192 672q40 113 117.5 168.5T511 896t202.5-55.5T832 672q0-13-9.5-22.5T800 640zM511 832q-166 0-223-114q0-6 5.5-10t13.5-4h410q8 0 13.5 4t5.5 10q-57 114-225 114z"/>`),
		g.Group(children),
	)
}