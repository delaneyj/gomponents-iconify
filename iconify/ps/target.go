package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Target(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0Q150 0 75 75T0 256t75 181t181 75t181-75t75-181t-75-181T256 0zm0 469q-88 0-150.5-62.5T43 256t62.5-150.5T256 43t150.5 62.5T469 256t-62.5 150.5T256 469zm0-362q-62 0-105.5 43.5T107 256t43.5 105.5T256 405t105.5-43.5T405 256t-43.5-105.5T256 107zm0 256q-45 0-76-31t-31-76t31-76t76-31t76 31t31 76t-31 76t-76 31zm43-107q0 18-12.5 30.5T256 299t-30.5-12.5T213 256t12.5-30.5T256 213t30.5 12.5T299 256z"/>`),
		g.Group(children),
	)
}