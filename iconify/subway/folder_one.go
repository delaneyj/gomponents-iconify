package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m0 192l36.6 237.7c0 8 10.3 18.3 18.3 18.3h402.3c8 0 18.3-10.3 18.3-18.3L512 192H0zm475.4-54.9c0-8-10.3-18.3-18.3-18.3H274.3l-9.1-36.6c-1.3-7.8-10.3-18.3-18.3-18.3h-128c-7.9 0-17 10.4-18.3 18.3l-9.1 36.6H54.9c-7.9 0-18.3 10.3-18.3 18.3v18.3h438.9v-18.3z"/>`),
		g.Group(children),
	)
}