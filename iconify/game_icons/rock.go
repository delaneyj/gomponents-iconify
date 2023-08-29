package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M228.813 23L68.75 72.28L39.5 182.095l47.53-21.22l10.44-4.655l2.5 11.155l8.75 39.125l6.405 28.53l-21.75-19.53l-15.72-14.125l-28.218 32.344l140.657 136l9.656-40.69l7.53-31.874l10.407 31.063l54.72 163.592l159.936-26.31l45.75-202.938l-84.563-148.718L228.814 23zm-57.688 49.875l-27.813 39.906l-3.25 73.44l-27.187-88.94l58.25-24.405zm17.844 93.406l113.124 155.25L407 355.407l-107.375-.844l-110.656-128v-60.28zM79.312 330.25l140.125 153.125l-5.563-65.875l-134.563-87.25z"/>`),
		g.Group(children),
	)
}