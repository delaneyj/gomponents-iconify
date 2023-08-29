package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WebhooksLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M114.59 188A52 52 0 0 1 12 176a51.54 51.54 0 0 1 10.41-31.21a12 12 0 1 1 19.18 14.41A27.74 27.74 0 0 0 36 176a28 28 0 0 0 56 0a12 12 0 0 1 12-12h88a12 12 0 0 1 0 24Zm-56.88-1.78a12 12 0 0 0 16.51-3.93l43-69.94a12 12 0 0 0-3.93-16.51a28 28 0 1 1 40.34-35.06a12 12 0 1 0 22-9.59a52 52 0 1 0-84.42 57.62l-37.47 60.9a12 12 0 0 0 3.97 16.51ZM192 124a51.89 51.89 0 0 0-16.3 2.62L138.22 65.7a12 12 0 0 0-20.44 12.58l43 69.93a12 12 0 0 0 16.52 3.93A28 28 0 1 1 192 204a12 12 0 0 0 0 24a52 52 0 0 0 0-104Z"/>`),
		g.Group(children),
	)
}