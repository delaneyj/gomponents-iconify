package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MailReadFortyEightFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="m6.05 18.573l17.313 10.253a1.25 1.25 0 0 0 1.274 0L41.95 18.574c.033.239.05.481.05.726v15.45A4.25 4.25 0 0 1 37.75 39h-27.5A4.25 4.25 0 0 1 6 34.75V19.3c0-.245.017-.488.05-.727ZM24.63 6.17l14.762 8.593a5.245 5.245 0 0 1 1.616 1.464L24 26.297L6.993 16.227a5.245 5.245 0 0 1 1.616-1.464L23.37 6.17a1.25 1.25 0 0 1 1.258 0Z"/>`),
		g.Group(children),
	)
}