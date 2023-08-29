package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JoOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagJo1x10"><path fill-opacity=".7" d="M113.6 0H607v493.5H113.6z"/></clipPath></defs><g clip-path="url(#flagJo1x10)" transform="translate(-117.8) scale(1.0375)"><g fill-rule="evenodd" stroke-width="1pt"><path d="M0 0h987v164.5H0z"/><path fill="#fff" d="M0 164.5h987V329H0z"/><path fill="#090" d="M0 329h987v164.5H0z"/><path fill="red" d="m0 493.5l493.5-246.8L0 0v493.5z"/><path fill="#fff" d="m164.8 244l22 10.6h-24.5l5.5 24l-15.3-19.3l-15.3 19.2l5.5-23.9H118l22.1-10.7l-15.3-19.1l22.1 10.6l5.5-23.9l5.5 24l22-10.7z"/></g></g>`),
		g.Group(children),
	)
}