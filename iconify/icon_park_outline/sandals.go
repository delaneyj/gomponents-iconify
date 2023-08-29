package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sandals(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="4"><path stroke-linejoin="bevel" d="m11.2 6.854l-.708.225c-3.487 1.104-5.536 4.71-4.898 8.312c.73 4.115 1.291 8.04 2.095 11.661c.785 3.537.447 6.689.25 10.032c-.192 3.266 2.409 5.843 5.672 6.07c4.623.321 8.182-4.082 7.297-8.63c-.684-3.516-1.202-7.318-.969-10.024c.25-2.904.277-7.29.23-11.251c-.055-4.616-4.568-7.788-8.968-6.395Zm23.61-1.131l-.77.144c-3.455.648-5.892 3.761-5.926 7.277C28.071 17.701 27.831 22.15 28 26c.156 3.553-.437 6.153-1.017 9.053c-.627 3.13 1.443 6.044 4.54 6.815c4.485 1.116 8.712-2.698 8.577-7.318c-.093-3.175.03-6.512.681-8.95c.79-2.96 1.547-7.903 2.129-12.314c.61-4.632-3.509-8.424-8.1-7.563Z"/><path stroke-linejoin="round" stroke-miterlimit="2" d="M8 29c1-9 6-14 6-14c1.636 2 4 8 6 14m20 1c.5-8-5-16-5-16c-1.636 2-7 5.278-7 11.5"/></g>`),
		g.Group(children),
	)
}