package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Myspace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M436 247c-59 0-106-47-106-106S377 34 436 34s107 48 107 107s-48 106-107 106zm-194-18c-43 0-78-35-78-78s35-78 78-78c44 0 79 35 79 78s-35 78-79 78zm340 183c0 3 0 7-1 10h1v252H291V542H136v-66H0V319c0-48 39-87 87-87c34 0 62 18 77 46c20-21 47-35 78-35c43 0 80 25 97 62c26-24 60-38 97-38c81 0 146 65 146 145zM87 220c-35 0-64-28-64-64c0-35 29-64 64-64c36 0 64 29 64 64c0 36-28 64-64 64z"/>`),
		g.Group(children),
	)
}