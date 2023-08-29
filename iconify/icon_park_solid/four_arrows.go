package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FourArrows(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFourArrows0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m24 24l-5-5m5-11v16V8Zm0 16l5-5l-5 5Zm0 0l-5 5m5 11V24v16Zm0-16l5 5l-5-5Z"/><path fill="#fff" d="M20 12a4 4 0 1 0 0-8a4 4 0 0 0 0 8ZM8 32a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/><path d="M8 24h32"/><path fill="#fff" d="M40 24a4 4 0 1 0 0-8a4 4 0 0 0 0 8ZM28 44a4 4 0 1 0 0-8a4 4 0 0 0 0 8Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFourArrows0)"/>`),
		g.Group(children),
	)
}