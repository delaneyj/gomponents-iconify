package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gauge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0zm320 96c0-26.9-16.5-49.9-40-59.3V88c0-13.3-10.7-24-24-24s-24 10.7-24 24v204.7c-23.5 9.5-40 32.5-40 59.3c0 35.3 28.7 64 64 64s64-28.7 64-64zM144 176a32 32 0 1 0 0-64a32 32 0 1 0 0 64zm-16 80a32 32 0 1 0-64 0a32 32 0 1 0 64 0zm288 32a32 32 0 1 0 0-64a32 32 0 1 0 0 64zm-16-144a32 32 0 1 0-64 0a32 32 0 1 0 64 0z"/>`),
		g.Group(children),
	)
}