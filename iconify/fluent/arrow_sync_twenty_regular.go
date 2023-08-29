package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowSyncTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M11.414 3.635a.5.5 0 0 0 0-.707L9.293.807a.5.5 0 0 0-.707.707l.997.997a7.5 7.5 0 0 0-4.075 13.495a.5.5 0 0 0 .6-.8A6.5 6.5 0 0 1 10.066 3.5c.024 0 .05-.002.073-.005L8.586 5.049a.5.5 0 0 0 .707.707l2.121-2.121ZM8.586 16.363a.5.5 0 0 0 0 .707l2.121 2.121a.5.5 0 0 0 .707-.707l-.997-.997a7.5 7.5 0 0 0 4.075-13.495a.5.5 0 1 0-.6.8a6.5 6.5 0 0 1-3.959 11.706a.502.502 0 0 0-.073.005l1.554-1.554a.5.5 0 1 0-.707-.707l-2.121 2.121Z"/>`),
		g.Group(children),
	)
}