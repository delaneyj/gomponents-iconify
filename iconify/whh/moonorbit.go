package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Moonorbit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-105 0-202-42q28-21 46-51q76 29 156 29q91 0 174-35.5T829 829t95.5-143T960 512t-35.5-174T829 195T686 99.5T512 64T338 99.5T195 195T99.5 338T64 512q0 80 29 156q-30 18-51 46Q0 617 0 512q0-104 40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm0-768q106 0 181 75t75 181t-75 181t-181 75t-181-75t-75-181t75-181t181-75zM192 704q53 0 90.5 37.5T320 832t-37.5 90.5T192 960t-90.5-37.5T64 832t37.5-90.5T192 704z"/>`),
		g.Group(children),
	)
}