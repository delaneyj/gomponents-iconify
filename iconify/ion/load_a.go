package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LoadA(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<circle cx="256" cy="96" r="64" fill="currentColor"/><circle cx="96" cy="256" r="48" fill="currentColor"/><circle cx="368" cy="144" r="8" fill="currentColor"/><path d="M180.1 107.6c-19.9-20.1-52.2-20.1-72.1 0-19.9 20.1-19.9 52.7 0 72.8 19.9 20.1 52.2 20.1 72.1 0 19.9-20.1 19.9-52.7 0-72.8z" fill="currentColor"/><circle cx="416" cy="256" r="16" fill="currentColor"/><circle cx="369" cy="369" r="24" fill="currentColor"/><circle cx="256" cy="416" r="32" fill="currentColor"/><circle cx="144" cy="368" r="40" fill="currentColor"/>`),
		g.Group(children),
	)
}