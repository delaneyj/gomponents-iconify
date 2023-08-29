package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bintray(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 476.07h400.36V512H0v-35.93zm286.155-305.403v35.93h52.612L218.145 327.218V69.293l37.213 37.213l25.665-25.664L200.18 0l-80.842 80.842l25.665 25.664l37.213-38.496v259.208L61.594 206.597h52.612v-35.93H0v114.205h35.93v-53.895l164.25 164.251l164.251-164.25v53.894h35.93V170.667H286.155z"/>`),
		g.Group(children),
	)
}