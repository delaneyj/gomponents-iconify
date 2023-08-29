package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trident(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M17.156 16.438v41.656L192.78 234.22l-51.436 51.405l-11.97 12L140.782 309l106.69 106.688l-24.532 24.53l125.75 53.844l-53.875-125.718l-23.407 23.406l-94.72-94.72l48.44-48.436l135.78 135.75l-23.97 23.937l125.72 53.876l-53.844-125.72l-23.968 23.97l-135.78-135.75l48.467-48.47l94.72 94.72l-23.375 23.406l125.72 53.844l-53.876-125.72l-24.533 24.533L309.5 140.28l-11.406-11.374l-11.97 11.97l-51.468 51.436L58.812 16.438H17.157z"/>`),
		g.Group(children),
	)
}