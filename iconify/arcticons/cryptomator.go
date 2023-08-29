package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cryptomator(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m42.8 31.84l.7-2.12l-1.3-2.45l-1.12-.27l-2.62-3.63c0-2.73-1.6-3.36-3.67-3.3V25l3.92 3.07l-.79.89l.26 2.93l1.93.81m-1.29-4.65l2.37-1.09M24 22.9h0a2.42 2.42 0 0 1 .78 4.71l1.06 5.12l-1.84.51l-1.86-.51l1.06-5.12a2.42 2.42 0 0 1 .8-4.71Zm3.26 14.7c.14 1.48.42 3 .73 3h6.83c2.23 0 2.26-8.88.16-8.88h-1.17m-28.61.12l-.7-2.12l1.3-2.45L6.92 27l2.62-3.63c0-2.73 1.6-3.36 3.67-3.3V25l-3.92 3.05l.79.89l-.26 2.93l-1.93.81m1.4-4.63l-2.37-1.09M20.86 37.6c-.14 1.48-.42 3-.74 3h-6.83c-2.22 0-2.26-8.88-.16-8.88h1.18m15.62-15.87h-3v-.76a1.52 1.52 0 0 1 3 0Zm-8.81 0h-3v-.76a1.52 1.52 0 0 1 3 0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.31 19.59H14.69a1.48 1.48 0 0 0-1.48 1.48v6.24a10.79 10.79 0 0 0 21.58 0v-6.24a1.48 1.48 0 0 0-1.48-1.48ZM24 7.31h0A10.79 10.79 0 0 1 34.79 18.1v0a1.49 1.49 0 0 1-1.49 1.49H14.69a1.49 1.49 0 0 1-1.49-1.49v0A10.79 10.79 0 0 1 24 7.31Z"/>`),
		g.Group(children),
	)
}