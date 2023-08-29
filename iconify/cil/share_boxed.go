package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareBoxed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M464 248v216H48V48h216V16H16v480h480V216l-32 32z"/><path fill="currentColor" d="M106.12 171.135A96.274 96.274 0 0 0 96 214.364v216.181h47.782l41.181-147.564a66.953 66.953 0 0 1 64.283-48.8H304V320h38.6l152.027-151.1L342.656 16H304v88h-93.818a114.4 114.4 0 0 0-104.062 67.135ZM336 136V54.7l113.373 114.058L336 281.441v-79.259h-86.754a99.055 99.055 0 0 0-95.105 72.2L128 368.051V214.364a64.576 64.576 0 0 1 6.879-29.2l.292-.614A82.356 82.356 0 0 1 210.182 136Z"/>`),
		g.Group(children),
	)
}