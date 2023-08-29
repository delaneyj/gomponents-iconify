package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yinyang(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1025q-104 0-199-40.5t-163.5-109T40.5 712T0 512.5t40.5-199t109-163T313 41T512 0t199 41t163.5 109.5t109 163t40.5 199T983.5 712t-109 163.5t-163.5 109t-199 40.5zm0-768q-26 0-45 18.5t-19 45t18.5 45.5t45.5 19t45.5-19t18.5-45.5t-19-45t-45-18.5zm0 256q-80 0-136-56.5t-56-136T376 185t136-56q-104 0-192.5 51t-140 139.5t-51.5 193t51.5 193t140 140T512 897q80 0 136-56.5t56-136T648 569t-136-56zm.5 256q-26.5 0-45.5-19t-19-45.5t19-45t45.5-18.5t45 18.5t18.5 45t-18.5 45.5t-45 19z"/>`),
		g.Group(children),
	)
}