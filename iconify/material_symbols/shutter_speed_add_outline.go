package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShutterSpeedAddOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 13Zm1.75-1l1.375 2.4q-.2.175-.375.35t-.35.375L13.75 14l-3.35 5.8q.675.175 1.338.2t1.312-.075q.075.5.238 1t.412.925q-.425.075-.85.113T12 22q-1.875 0-3.512-.712T5.625 19.35Q4.4 18.125 3.7 16.487T3 13q0-1.85.713-3.488T5.65 6.65q1.225-1.225 2.863-1.938T12 4q1.575 0 3 .525T17.6 6l1.45-1.45l1.4 1.4l-1.4 1.45q.95 1.225 1.475 2.738T21 13.325q-.5-.15-1-.25t-1-.1q0-.25-.025-.487T18.9 12h-5.15ZM18 23v-3h-3v-2h3v-3h2v3h3v2h-3v3h-2ZM9 3V1h6v2H9Zm3 8h6.7q-.45-1.55-1.538-2.725T14.6 6.5L12 11Zm-1.75 1l3.35-5.8q-1.475-.375-3.038-.063T7.65 7.5l2.6 4.5ZM5.1 14h5.15L6.9 8.2q-1.05 1.175-1.563 2.663T5.1 14Zm4.3 5.5L12 15H5.3q.45 1.55 1.538 2.725T9.4 19.5Z"/>`),
		g.Group(children),
	)
}