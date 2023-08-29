package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BottleDroplet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M96 0C82.7 0 72 10.7 72 24s10.7 24 24 24c4.4 0 8 3.6 8 8v64.9c0 12.2-7.2 23.1-17.2 30.1C53.7 174.1 32 212.5 32 256v192c0 35.3 28.7 64 64 64h128c35.3 0 64-28.7 64-64V256c0-43.5-21.7-81.9-54.8-105c-10-7-17.2-17.9-17.2-30.1V56c0-4.4 3.6-8 8-8c13.3 0 24-10.7 24-24S237.3 0 224 0H96zm64 382c-26.5 0-48-20.1-48-45c0-16.8 22.1-48.1 36.3-66.4c6-7.8 17.5-7.8 23.5 0c14.1 18.3 36.2 49.6 36.2 66.4c0 24.9-21.5 45-48 45z"/>`),
		g.Group(children),
	)
}