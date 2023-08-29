package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Play(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="fePlay0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="fePlay1" fill="currentColor" fill-rule="nonzero"><path id="fePlay2" d="M6 5.912c0-.155.037-.307.107-.443c.23-.44.75-.599 1.163-.354l10.29 6.088c.14.083.255.206.332.355c.23.44.08.995-.332 1.239L7.27 18.885a.814.814 0 0 1-.415.115C6.383 19 6 18.592 6 18.089V5.912Z"/></g></g>`),
		g.Group(children),
	)
}