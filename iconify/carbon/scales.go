package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Scales(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20 16a5 5 0 0 0 10 0a1 1 0 0 0-.105-.447l-3.999-7.997a.891.891 0 0 0-.045-.081A1 1 0 0 0 25 7h-6.178A3.015 3.015 0 0 0 17 5.184V2h-2v3.184A3.015 3.015 0 0 0 13.178 7H7a1 1 0 0 0-.894.553l-4 8A1 1 0 0 0 2 16a5 5 0 0 0 10 0a1 1 0 0 0-.105-.447L8.617 9h4.56A3.015 3.015 0 0 0 15 10.815V28H6v2h20v-2h-9V10.816A3.015 3.015 0 0 0 18.822 9h4.56l-3.277 6.553A1 1 0 0 0 20 16ZM7 19a2.996 2.996 0 0 1-2.815-2h5.63A2.996 2.996 0 0 1 7 19Zm2.382-4H4.618L7 10.236ZM16 9a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm9 10a2.996 2.996 0 0 1-2.815-2h5.63A2.996 2.996 0 0 1 25 19Zm0-8.764L27.382 15h-4.764Z"/>`),
		g.Group(children),
	)
}