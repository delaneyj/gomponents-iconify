package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PsOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagPs1x10"><path fill-opacity=".7" d="M237.1 0h493.5v493.5H237.1z"/></clipPath></defs><g clip-path="url(#flagPs1x10)" transform="translate(-246) scale(1.0375)"><g fill-rule="evenodd" stroke-width="1pt"><path d="M0 0h987v164.5H0z"/><path fill="#fff" d="M0 164.5h987V329H0z"/><path fill="#090" d="M0 329h987v164.5H0z"/><path fill="red" d="m0 493.5l493.5-246.8L0 0v493.5z"/></g></g>`),
		g.Group(children),
	)
}