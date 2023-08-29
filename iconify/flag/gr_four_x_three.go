package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GrFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#0d5eaf" fill-rule="evenodd" d="M0 0h640v53.3H0z"/><path fill="#fff" fill-rule="evenodd" d="M0 53.3h640v53.4H0z"/><path fill="#0d5eaf" fill-rule="evenodd" d="M0 106.7h640V160H0z"/><path fill="#fff" fill-rule="evenodd" d="M0 160h640v53.3H0z"/><path fill="#0d5eaf" d="M0 0h266.7v266.7H0z"/><path fill="#0d5eaf" fill-rule="evenodd" d="M0 213.3h640v53.4H0z"/><path fill="#fff" fill-rule="evenodd" d="M0 266.7h640V320H0z"/><path fill="#0d5eaf" fill-rule="evenodd" d="M0 320h640v53.3H0z"/><path fill="#fff" fill-rule="evenodd" d="M0 373.3h640v53.4H0z"/><g fill="#fff" fill-rule="evenodd" stroke-width="1.3"><path d="M106.7 0H160v266.7h-53.3z"/><path d="M0 106.7h266.7V160H0z"/></g><path fill="#0d5eaf" d="M0 426.7h640V480H0z"/>`),
		g.Group(children),
	)
}