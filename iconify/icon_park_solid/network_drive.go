package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NetworkDrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSNetworkDrive0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M40.518 36.316A9.21 9.21 0 0 0 44 26c-1.213-3.83-4.93-5.929-8.947-5.925h-2.321a14.737 14.737 0 1 0-25.31 13.429"/><path fill="#fff" d="M14 35h20v6H14z"/><path d="M34 28H22m-6 0h-2"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSNetworkDrive0)"/>`),
		g.Group(children),
	)
}