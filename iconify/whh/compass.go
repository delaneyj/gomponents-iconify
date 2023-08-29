package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-104 0-199-40.5t-163.5-109T40.5 711T0 512t40.5-199t109-163.5T313 40.5T512 0t199 40.5t163.5 109t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm0-896q-104 0-192.5 51t-140 139.5t-51.5 193t51.5 193t140 140T512 896t192.5-51.5t140-140t51.5-193t-51.5-193t-140-139.5T512 128zm0 448q-35 0-54-30L224 416l288 32q35 0 54 30l234 130z"/>`),
		g.Group(children),
	)
}