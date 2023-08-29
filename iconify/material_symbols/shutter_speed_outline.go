package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShutterSpeedOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M9 3V1h6v2H9Zm3 19q-1.875 0-3.513-.713T5.626 19.35Q4.4 18.125 3.7 16.487T3 13q0-1.85.713-3.488T5.65 6.65q1.225-1.225 2.863-1.938T12 4q1.575 0 3 .525T17.6 6l1.45-1.45l1.4 1.4l-1.4 1.45q.9 1.175 1.425 2.6T21 13q0 1.85-.7 3.488t-1.925 2.862q-1.225 1.225-2.863 1.938T12 22Zm0-9Zm0-2h6.7q-.45-1.55-1.538-2.725T14.6 6.5L12 11Zm-1.75 1l3.35-5.8q-1.475-.375-3.038-.063T7.65 7.5l2.6 4.5ZM5.1 14h5.15L6.9 8.2q-1.05 1.175-1.563 2.663T5.1 14Zm4.3 5.5L12 15H5.3q.45 1.55 1.538 2.725T9.4 19.5Zm1 .3q1.65.425 3.2.038t2.75-1.338l-2.6-4.5l-3.35 5.8Zm6.7-2q1.1-1.2 1.587-2.688T18.9 12h-5.15l3.35 5.8Z"/>`),
		g.Group(children),
	)
}