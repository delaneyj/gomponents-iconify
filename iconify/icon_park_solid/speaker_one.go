package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSSpeakerOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#fff" d="M4 15h11v18H4z"/><path d="m15 15l18-7v32l-18-7m25-16h2m-3 8h5m-4 8h2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSSpeakerOne0)"/>`),
		g.Group(children),
	)
}