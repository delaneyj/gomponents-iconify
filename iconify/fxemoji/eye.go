package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eye(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#E5E4DF" d="M256 0c141.4 0 256 114.6 256 256S397.4 512 256 512S0 397.3 0 256C0 114.6 114.6 0 256 0z"/><path fill="#00B1FF" d="M256 80c97.2 0 176 78.8 176 176s-78.8 176-176 176S80 353.2 80 256S158.8 80 256 80z"/><path fill="#2B3B47" d="M256 160c53 0 96 43 96 96s-43 96-96 96s-96-43-96-96s43-96 96-96z"/><path fill="#D4EDF6" d="M327.9 160c22.1 0 40 17.9 40 40s-17.9 40-40 40s-40-17.9-40-40s18-40 40-40z"/><path fill="#D6DBDE" d="M349.2 233.7c-6.5-27.3-24.5-50-48.7-62.7c-7.7 7.3-12.6 17.5-12.6 29c0 22.1 17.9 40 40 40c7.9-.1 15.2-2.4 21.3-6.3z"/>`),
		g.Group(children),
	)
}