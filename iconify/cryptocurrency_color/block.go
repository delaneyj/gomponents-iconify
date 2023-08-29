package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Block(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#101341" fill-rule="nonzero"/><g fill="#FFF"><path d="M11.022 7H21.25l5.25 9l-5.25 9H10.931l5.16-9l-5.07-9zm5.43 3.166L19.803 16l-3.35 5.834h2.988L22.789 16l-3.35-5.834h-2.986z"/><path d="M12.113 11.026L9.211 16l2.876 4.93l-1.839 3.209L5.5 16l4.789-8.211z" opacity=".5"/></g></g>`),
		g.Group(children),
	)
}