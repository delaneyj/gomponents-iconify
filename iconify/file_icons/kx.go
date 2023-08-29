package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kx(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M512 435.35L408.464 292.07l89.843-125.244h-95.52l-42.634 70.533l-43.535-70.533H156.973L82.16 263.188V76.65H0v358.7h82.16V306.795l76.15 128.555h147.955l53.965-83.507l54.914 83.507H512zm-281.55-30.392l-78.153-119.901l82.829-97.858l75.815 105.205l-80.49 112.554z"/>`),
		g.Group(children),
	)
}