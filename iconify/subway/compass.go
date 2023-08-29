package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Compass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0zm0 472.6c-119.6 0-216.6-97-216.6-216.6S136.4 39.4 256 39.4s216.6 97 216.6 216.6s-97 216.6-216.6 216.6zm-137.8-78.8l187.1-88.6l88.6-187.1l-187.1 88.6l-88.6 187.1zm167.3-108.3l-118.2 59.1l59.1-118.2l59.1 59.1z"/>`),
		g.Group(children),
	)
}