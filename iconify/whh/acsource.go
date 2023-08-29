package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Acsource(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm0-896q-104 0-192.5 51.5t-140 140T128 512t51.5 192.5t140 140T512 896t192.5-51.5t140-140T896 512t-51.5-192.5t-140-140T512 128zm128 576q-58 0-97.5-42T480 512q-21-49-47.5-88.5T384 384q-15 0-25 18.5t-21 56t-18 53.5h-64q0-27 7-57.5t21-62t40-52t60-20.5q58 0 97.5 42T544 512q21 49 47.5 88.5T640 640q15 0 25-18.5t21-56t18-53.5h64q0 27-7 57.5t-21 62t-40 52t-60 20.5z"/>`),
		g.Group(children),
	)
}