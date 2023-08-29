package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Corn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23 16c-.337 0-.67.02-1 .05V16c0-7.885-2.956-14-5.5-14S11 8.115 11 16v5H6v3a6.007 6.007 0 0 0 6 6h3a11.012 11.012 0 0 0 11-11v-3Zm-10 0h2v-2h-1.927a25.894 25.894 0 0 1 .565-4H16V8h-1.8c.81-2.406 1.815-3.772 2.3-3.986c.66.292 2.284 2.703 3.065 6.986H18v2h1.844A29.163 29.163 0 0 1 20 16v.426A11.055 11.055 0 0 0 13.79 21H13Zm-1 12a4.005 4.005 0 0 1-4-4v-1h4.764A10.903 10.903 0 0 0 12 27Zm12-9a9.01 9.01 0 0 1-9 9h-1v-1a9.01 9.01 0 0 1 9-9h1Z"/>`),
		g.Group(children),
	)
}