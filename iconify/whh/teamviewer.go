package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Teamviewer(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-384-896q-104 0-192.5 51.5t-140 140t-51.5 192.5t51.5 192.5t140 140t192.5 51.5t192.5-51.5t140-140t51.5-192.5t-51.5-192.5t-140-140t-192.5-51.5zm128 448h-256v128l-192-192l192-192v128h256V320l192 192l-192 192V576z"/>`),
		g.Group(children),
	)
}