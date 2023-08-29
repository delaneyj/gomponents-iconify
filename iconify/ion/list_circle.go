package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 48C141.31 48 48 141.31 48 256s93.31 208 208 208s208-93.31 208-208S370.69 48 256 48Zm-88 302a24 24 0 1 1 24-24a24 24 0 0 1-24 24Zm0-71a24 24 0 1 1 24-24a24 24 0 0 1-24 24Zm0-73a24 24 0 1 1 24-24a24 24 0 0 1-24 24Zm184 135H224a16 16 0 0 1 0-32h128a16 16 0 0 1 0 32Zm0-71H224a16 16 0 0 1 0-32h128a16 16 0 0 1 0 32Zm0-72H224a16 16 0 0 1 0-32h128a16 16 0 0 1 0 32Z"/>`),
		g.Group(children),
	)
}