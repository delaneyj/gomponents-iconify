package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Libreoffice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21.818 0a.795.795 0 0 0-.74.469a.78.78 0 0 0 .172.849l6.646 6.661a.811.811 0 0 0 .849.177a.792.792 0 0 0 .484-.708V.771A.807.807 0 0 0 28.448 0zM3.547 0a.787.787 0 0 0-.776.786v30.427c0 .432.349.781.776.786h24.896a.789.789 0 0 0 .786-.786V11.619a.792.792 0 0 0-.229-.557L18.234.244a.799.799 0 0 0-.563-.245zm.786 1.578h13.005l10.313 10.359v18.484H4.333z"/>`),
		g.Group(children),
	)
}