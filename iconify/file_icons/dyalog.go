package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dyalog(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M.713 0v136.303h75.483V75.426h116.99c63.329 0 115.82 19.655 151.877 56.74c31.551 32.606 48.924 76.51 48.924 123.779c0 86.979-62.817 180.559-200.802 180.559H0V512h193.185c181.346 0 276.199-128.774 276.199-256.056c0-66.924-24.933-129.598-70.234-176.268C348.572 27.557 277.311 0 193.186 0z"/>`),
		g.Group(children),
	)
}