package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bigace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm0-896q0 17-16.5 36.5T452 203t-61.5 43.5T320 296t-70.5 59t-61.5 70t-43.5 83.5T128 608q0 76 54 118t138 42q39 0 69-16.5t59-47.5q0 32-1 49.5t-6.5 45t-19.5 51t-37 46.5h256q-23-23-37-46.5t-19.5-51t-6.5-45t-1-49.5q29 31 59 47.5t69 16.5q83 0 137.5-42T896 608q0-59-21-112t-55-92.5t-75-74.5t-82-62.5t-75-51t-55-46t-21-41.5z"/>`),
		g.Group(children),
	)
}