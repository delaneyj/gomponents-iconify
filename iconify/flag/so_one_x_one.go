package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagSo1x10"><path fill-opacity=".7" d="M177.2 0h708.6v708.7H177.2z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagSo1x10)" transform="translate(-128) scale(.72249)"><path fill="#40a6ff" d="M0 0h1063v708.7H0z"/><path fill="#fff" d="m643 527.6l-114.3-74l-113.6 74.8l42.3-121.5l-113.5-75l140.4-1l43.5-121.1l44.5 120.8l140.3.1l-112.9 75.7L643 527.6z"/></g>`),
		g.Group(children),
	)
}