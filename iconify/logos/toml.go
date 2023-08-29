package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Toml(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#9C4221" d="M198.472 0v28.764h25.888v198.472h-25.888V256H256V0z"/><path d="M64.719 83.416V51.775h126.562v31.641H143.82v138.067h-31.64V83.416z"/><path fill="#9C4221" d="M57.528 0v28.764H31.64v198.472h25.888V256H0V0z"/>`),
		g.Group(children),
	)
}