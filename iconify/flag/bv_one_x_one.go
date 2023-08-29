package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BvOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagBv1x10"><path fill-opacity=".7" d="M0 0h512v512H0z"/></clipPath></defs><g fill-rule="evenodd" stroke-width="1pt" clip-path="url(#flagBv1x10)"><path fill="#fff" d="M-68 0h699.7v512H-68z"/><path fill="#d72828" d="M-93-77.8h218.7v276.2H-93zM249.4-.6h381v199h-381zM-67.6 320h190.4v190.3H-67.5zm319.6 2.1h378.3v188.2H252z"/><path fill="#003897" d="M156.7-25.4H221v535.7h-64.5z"/><path fill="#003897" d="M-67.5 224.8h697.8v63.5H-67.5z"/></g>`),
		g.Group(children),
	)
}