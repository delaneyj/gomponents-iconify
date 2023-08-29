package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartRingOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.174 2.983l-.98.2a9 9 0 1 0 3.612 17.635a9.003 9.003 0 0 0 7.057-7.237l.176-.984l1.969.351l-.176.985c-.769 4.302-4.079 7.913-8.625 8.844c-5.951 1.22-11.764-2.617-12.983-8.569C.004 8.257 3.841 2.444 9.793 1.225l.98-.201l.4 1.96ZM13.31 1.03l.978.209a11.002 11.002 0 0 1 8.488 8.555l.201.98l-1.96.4l-.2-.98a9.002 9.002 0 0 0-6.946-7l-.978-.208l.416-1.956Zm2.61 10.168a4 4 0 1 0-7.838 1.605a4 4 0 0 0 7.838-1.605Zm-5.123-5.076a6 6 0 1 1 2.408 11.756a6 6 0 0 1-2.408-11.756Z"/>`),
		g.Group(children),
	)
}