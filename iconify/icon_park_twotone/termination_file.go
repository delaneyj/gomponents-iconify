package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TerminationFile(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTTerminationFile0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="round" d="M40 16.398V6a2 2 0 0 0-2-2H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h10"/><path d="M16 14h13m-13 7h5"/><path fill="#555" stroke-linejoin="round" d="M34 44c-5.523 0-10-4.477-10-10s4.477-10 10-10s10 4.477 10 10s-4.477 10-10 10Z"/><path d="m27 27l14 14"/><path stroke-linejoin="round" d="M24 34c0-5.523 4.477-10 10-10m0 20c5.523 0 10-4.477 10-10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTTerminationFile0)"/>`),
		g.Group(children),
	)
}