package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImageSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M456 64H56a24 24 0 0 0-24 24v336a24 24 0 0 0 24 24h400a24 24 0 0 0 24-24V88a24 24 0 0 0-24-24Zm-124.38 64.2a48 48 0 1 1-43.42 43.42a48 48 0 0 1 43.42-43.42ZM76 416a12 12 0 0 1-12-12v-87.63L192.64 202l96.95 96.75L172.37 416Zm372-12a12 12 0 0 1-12 12H217.63l149.53-149.53L448 333.84Z"/>`),
		g.Group(children),
	)
}