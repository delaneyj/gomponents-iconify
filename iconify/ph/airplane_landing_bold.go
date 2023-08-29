package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AirplaneLandingBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M252 216a12 12 0 0 1-12 12H96a12 12 0 0 1 0-24h144a12 12 0 0 1 12 12Zm-31.24-24.45L44.14 142.09A44.13 44.13 0 0 1 12 99.72V48a20 20 0 0 1 26.32-19l5.48 1.83a12 12 0 0 1 7.49 7.3l9.91 27.46l22.8 6.5V48a20 20 0 0 1 26.32-19l5.48 1.83a12 12 0 0 1 7.27 6.74l21.75 51.85l59 16.49A44.12 44.12 0 0 1 236 148.32V180a12 12 0 0 1-15.24 11.55ZM212 148.32a20.05 20.05 0 0 0-14.65-19.27L132.77 111a12 12 0 0 1-7.84-6.91L108 63.71V88a12 12 0 0 1-15.29 11.53L48.71 87a12 12 0 0 1-8-7.46L36 66.48v33.24A20.07 20.07 0 0 0 50.61 119L212 164.18Z"/>`),
		g.Group(children),
	)
}