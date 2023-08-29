package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GooglePhotos(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.75 13.51h0A9.25 9.25 0 0 1 24 22.76V24h0H6.74a1.24 1.24 0 0 1-1.24-1.24v0a9.25 9.25 0 0 1 9.25-9.25Zm18.5 20.98h0A9.25 9.25 0 0 1 24 25.24V24h17.26a1.24 1.24 0 0 1 1.24 1.24h0a9.25 9.25 0 0 1-9.25 9.25Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.5 33.25h0A9.25 9.25 0 0 1 22.75 24H24v17.25a1.24 1.24 0 0 1-1.25 1.25h0a9.25 9.25 0 0 1-9.25-9.25Zm21-18.5h0A9.25 9.25 0 0 1 25.25 24H24h0V6.74a1.24 1.24 0 0 1 1.25-1.24h0a9.25 9.25 0 0 1 9.25 9.25Z"/>`),
		g.Group(children),
	)
}