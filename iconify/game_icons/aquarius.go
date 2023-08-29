package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aquarius(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m301.313 68.906l-88.875 69.438l-35.157-68.156l-141.218 93.406v85.72l108.626-76.69l39.437 67l93.03-65.06l34.658 69.28l93.343-68.094l67.97 78.563l1.28-112.75L426.5 70.438L337.437 137l-36.125-68.094zm0 194.125l-88.875 69.44l-35.157-68.126L36.063 357.72v85.717L144.69 366.75l39.437 67l93.03-65.063l34.658 69.282l93.343-68.064l67.97 78.53l1.28-112.748l-47.906-71.094l-89.063 66.53l-36.125-68.093z"/>`),
		g.Group(children),
	)
}