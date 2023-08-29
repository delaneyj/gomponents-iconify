package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChristmasTreeOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTChristmasTreeOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M24 36v8"/><path fill="#555" d="M14 15L24 4l10 11l-3 3l8 8l-5 1.158L42 36H6l8-8.842L9 26l8-8l-3-3Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTChristmasTreeOne0)"/>`),
		g.Group(children),
	)
}