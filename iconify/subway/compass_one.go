package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0zM42.7 256c0-117.8 95.5-213.3 213.3-213.3c117.8 0 213.3 95.5 213.3 213.3c0 144-426.6 144-426.6 0zm256 85.3L352 96L213.3 341.3h85.4z"/>`),
		g.Group(children),
	)
}