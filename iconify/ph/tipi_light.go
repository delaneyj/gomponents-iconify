package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TipiLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M237.05 212.77L135.12 53.5l21.93-34.26A6 6 0 1 0 147 12.77l-19 29.6l-19-29.6a6 6 0 1 0-10 6.47l21.88 34.26L19 212.77a6 6 0 0 0 5 9.23h208a6 6 0 0 0 5.05-9.23ZM82.64 210L128 139.13L173.36 210Zm105 0l-54.55-85.23a6 6 0 0 0-10.1 0L68.4 210H35l93-145.37L221 210Z"/>`),
		g.Group(children),
	)
}