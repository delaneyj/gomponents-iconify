package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1025q-104 0-199-40.5t-163.5-109T40.5 712T0 513t40.5-199t109-163.5T313 41T512 0t199 41t163.5 109.5t109 163.5t40.5 199t-40.5 199t-109 163.5t-163.5 109t-199 40.5zm0-896q-104 0-192.5 51t-140 139.5t-51.5 193t51.5 193t140 140T512 897t192.5-51.5t140-140t51.5-193t-51.5-193t-140-139.5T512 129zm192 224L573 530q-5 20-22 33.5T512 577q-21 0-38-13t-23-32L256 289v-32h32l225 181l159-117h32v32z"/>`),
		g.Group(children),
	)
}