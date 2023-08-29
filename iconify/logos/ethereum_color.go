package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EthereumColor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#5A9DED" d="M256 241.65L128 426.8V322.457l128-80.807Zm-67.84 65.08l-41.09 25.858v33.473l41.09-59.33Z"/><path fill="#53D3E0" d="M0 241.65L128 426.8V322.457L0 241.65Zm67.84 65.08l41.09 25.858v33.473l-41.09-59.33Z"/><path fill="#D995D1" d="m144.238 156.998l84.775 47.215l3.005 9.187l-10.094 8.812l-77.686-43.358z"/><path fill="#A6E276" d="m111.818 156.998l-84.775 47.215l-3.005 9.187l10.094 8.812l77.686-43.358z"/><path fill="#FF9C92" d="m128 0l128 219.203l-128 80.608V0Zm18.904 70.317V266.03l83.622-52.631l-83.622-143.083Z"/><path fill="#FFE94D" d="M128 0L0 219.203l128 80.608V0Zm-18.904 70.317V266.03L25.474 213.4l83.622-143.083Z"/>`),
		g.Group(children),
	)
}