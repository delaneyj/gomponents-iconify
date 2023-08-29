package cif

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 301 201"),
		g.Raw(`<defs><path id="cifMr0" d="M.724.5h299.777v200H.724z"/></defs><g fill="none" fill-rule="evenodd"><path fill="#006233" fill-rule="nonzero" d="M.5.5h299.777v200H.5z"/><ellipse cx="150.389" cy="75.5" fill="#FFC400" fill-rule="nonzero" rx="64.535" ry="64.583"/><mask id="cifMr1" fill="#fff"><use href="#cifMr0"/></mask><ellipse cx="150.389" cy="58.833" fill="#006233" fill-rule="nonzero" mask="url(#cifMr1)" rx="62.453" ry="62.5"/><path fill="#FFC400" fill-rule="nonzero" d="m126.63 51.108l14.683 10.676l-5.608 17.275l14.684-10.676l14.683 10.676l-5.609-17.275l14.684-10.676h-18.15l-5.608-17.275l-5.609 17.275z"/></g>`),
		g.Group(children),
	)
}