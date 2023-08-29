package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerSimpleSlashBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 88v80a12 12 0 0 1-24 0V88a12 12 0 0 1 24 0Zm-23.12 119.93a12 12 0 1 1-17.76 16.14L164 185.44V224a12 12 0 0 1-19.37 9.47L75.88 180H32a20 20 0 0 1-20-20V96a20 20 0 0 1 20-20h32.51L39.12 48.07a12 12 0 1 1 17.76-16.14ZM140 159l-53.67-59H36v56h44a12 12 0 0 1 7.37 2.53L140 199.46Zm-14.15-91.5l14.15-11v38.4a12 12 0 0 0 24 0V32a12 12 0 0 0-19.37-9.47L111.12 48.6a12 12 0 1 0 14.73 18.94ZM192 150.94a12 12 0 0 0 12-12V104a12 12 0 0 0-24 0v34.94a12 12 0 0 0 12 12Z"/>`),
		g.Group(children),
	)
}