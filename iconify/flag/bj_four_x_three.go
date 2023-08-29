package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BjFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagBj4x30"><path fill="gray" d="M67.6-154h666v666h-666z"/></clipPath></defs><g clip-path="url(#flagBj4x30)" transform="matrix(.961 0 0 .7207 -65 111)"><g fill-rule="evenodd" stroke-width="1pt"><path fill="#319400" d="M0-154h333v666H0z"/><path fill="#ffd600" d="M333-154h666v333H333z"/><path fill="#de2110" d="M333 179h666v333H333z"/></g></g>`),
		g.Group(children),
	)
}