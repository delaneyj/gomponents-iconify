package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WonSign(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M62.4 53.9c-5.6-16.8-23.8-25.8-40.5-20.3S-3.9 57.4 1.6 74.1l50 149.9H32c-17.7 0-32 14.3-32 32s14.3 32 32 32h40.9l56.7 170.1c4.5 13.5 17.4 22.4 31.6 21.9s26.4-10.4 29.8-24.2L233 288h46l42 167.8c3.4 13.8 15.6 23.7 29.8 24.2s27.1-8.4 31.6-21.9L439.1 288H480c17.7 0 32-14.3 32-32s-14.3-32-32-32h-19.6l50-149.9c5.6-16.8-3.5-34.9-20.2-40.5s-34.9 3.5-40.5 20.2L392.9 224H329L287 56.2C283.5 42 270.7 32 256 32s-27.5 10-31 24.2L183 224h-64L62.4 53.9zm78 234.1H167l-11.4 45.6l-15.2-45.6zM249 224l7-28.1l7 28.1h-14zm96 64h26.6l-15.2 45.6L345 288z"/>`),
		g.Group(children),
	)
}