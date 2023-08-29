package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompassTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M236.3 452.9h39.4v-39.4h-39.4v39.4zm39.4-393.8h-39.4v39.4h39.4V59.1zM59.1 275.7h39.4v-39.4H59.1v39.4zM256 0C114.6 0 0 114.6 0 256s114.6 256 256 256s256-114.6 256-256S397.4 0 256 0zm0 472.6c-119.6 0-216.6-97-216.6-216.6S136.4 39.4 256 39.4s216.6 97 216.6 216.6s-97 216.6-216.6 216.6zM137.8 256c0 65.2 52.9 118.2 118.2 118.2c65.2 0 118.2-52.9 118.2-118.2V137.8H256c-65.2 0-118.2 53-118.2 118.2zM256 196.9c-32.6 0-59.1 26.4-59.1 59.1h-19.7c0-43.5 35.3-78.8 78.8-78.8v19.7zm157.5 39.4v39.4h39.4v-39.4h-39.4z"/>`),
		g.Group(children),
	)
}