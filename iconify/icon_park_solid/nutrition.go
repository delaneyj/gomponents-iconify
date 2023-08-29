package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nutrition(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSNutrition0"><g fill="none" stroke-width="4"><path fill="#fff" fill-rule="evenodd" stroke="#fff" stroke-linejoin="round" d="M24 42h5.955c.083-2.737.484-4.242 1.204-4.515C38.669 34.635 44 27.434 44 19H4c0 8.251 5.103 15.323 12.357 18.294c.758.31 1.325 1.88 1.699 4.706H24Z" clip-rule="evenodd"/><path stroke="#000" stroke-linecap="round" d="M14.443 26.023c.36 1.187.836 2.168 1.427 2.942a11.253 11.253 0 0 0 2.14 2.104"/><path stroke="#fff" stroke-linecap="round" d="M32.2 8.018a16.767 16.767 0 0 0-4.047-1.613M22.05 6c-7.123.823-12.906 6.098-14.55 13m33-.003a17.144 17.144 0 0 0-3.883-7.434"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSNutrition0)"/>`),
		g.Group(children),
	)
}