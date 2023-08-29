package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KwOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagKw1x10"><path fill-opacity=".7" d="M0 0h496v496H0z"/></clipPath></defs><g fill-rule="evenodd" stroke-width="1pt" clip-path="url(#flagKw1x10)" transform="scale(1.0321)"><path fill="#fff" d="M0 165.3h992.1v165.4H0z"/><path fill="#f31830" d="M0 330.7h992.1v165.4H0z"/><path fill="#00d941" d="M0 0h992.1v165.4H0z"/><path d="M0 0v496l247.5-165.3l.5-165.5L0 0z"/></g>`),
		g.Group(children),
	)
}