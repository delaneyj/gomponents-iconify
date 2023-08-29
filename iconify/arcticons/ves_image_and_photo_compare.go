package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VesImageAndPhotoCompare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M27.436 7.303h11.242c.983 0 1.78.797 1.78 1.78v29.356c0 .983-.797 1.79-1.78 1.78H27.407V23.71l9.162 11.005v-23.97h-9.162V7.316m-6.873 16.342l-9.2 11.056h9.2V23.658Z"/><path d="M20.663 43.5v-3.28H9.322a1.78 1.78 0 0 1-1.78-1.78V9.081c0-.982.797-1.77 1.78-1.78h11.341V4.5l2.818.001v39h-2.818Z"/></g>`),
		g.Group(children),
	)
}