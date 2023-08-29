package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShieldStarBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M76.86 115.55a12 12 0 0 1 15.6-6.69l23.54 9.42V96a12 12 0 0 1 24 0v22.28l23.54-9.42a12 12 0 0 1 8.92 22.28L147 141.33l14.6 19.47a12 12 0 1 1-19.2 14.4L128 156l-14.4 19.2a12 12 0 1 1-19.2-14.4l14.6-19.47l-25.46-10.19a12 12 0 0 1-6.68-15.59ZM228 56v58.8c0 92.36-78.1 123-93.76 128.18a19.63 19.63 0 0 1-12.49 0C106.1 237.78 28 207.16 28 114.8V56a20 20 0 0 1 20-20h160a20 20 0 0 1 20 20Zm-24 4H52v54.8c0 73.56 60.53 99.53 76 105c15.48-5.43 76-31.4 76-104.95Z"/>`),
		g.Group(children),
	)
}