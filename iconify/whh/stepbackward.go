package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stepbackward(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896 1025q-53 0-90.5-37.5T768 897V128q0-53 37.5-90.5T896 0t90.5 37.5T1024 128v769q0 53-37.5 90.5T896 1025zm-446-11L15 551Q0 535 0 513t15-38L450 12q25-27 62 12v977q-37 40-62 13z"/>`),
		g.Group(children),
	)
}