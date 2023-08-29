package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dumpling(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#fcea2b" d="M57.09 18.09a6.431 6.431 0 0 0-2.467-1.533a4.371 4.371 0 0 0-3.793-.224a4.477 4.477 0 0 0-6.104-2.962a4.46 4.46 0 0 0-6.223-.2a4.445 4.445 0 0 0-6.532.574a4.472 4.472 0 0 0-6.716 2.333a4.434 4.434 0 0 0-5.252 4.424a4.432 4.432 0 0 0 .055.545a4.405 4.405 0 0 0-3.98 5.207a4.472 4.472 0 0 0-2.332 6.716a4.445 4.445 0 0 0-.575 6.532a4.445 4.445 0 0 0 .575 6.532A4.487 4.487 0 0 0 15.7 52.62a6.486 6.486 0 0 0 11.07 4.98l30.32-30.32a6.519 6.519 0 0 0 0-9.192z"/><path fill="none" stroke="#f1b31c" stroke-linecap="round" stroke-width="2" d="m14.25 33.21l3.393.613M20.56 21.28l2.823 2.014m9.097-9.314l.593 3.396"/><path fill="none" stroke="#f1b31c" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.001" d="M50.4 8.855A12.77 14.13 0 0 1 39-4.115a12.77 14.13 0 0 1 9.473-14.75" paint-order="stroke fill markers" transform="matrix(.6322 .7748 -.7436 .6686 0 0)"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M57.59 18.33a6.431 6.431 0 0 0-2.467-1.533a4.371 4.371 0 0 0-3.793-.224a4.477 4.477 0 0 0-6.104-2.962a4.46 4.46 0 0 0-6.223-.2a4.445 4.445 0 0 0-6.532.574a4.472 4.472 0 0 0-6.716 2.333a4.434 4.434 0 0 0-5.252 4.424a4.432 4.432 0 0 0 .055.545a4.405 4.405 0 0 0-3.98 5.207a4.472 4.472 0 0 0-2.332 6.716a4.445 4.445 0 0 0-.575 6.532a4.445 4.445 0 0 0 .575 6.532A4.487 4.487 0 0 0 16.2 52.86a6.486 6.486 0 0 0 11.07 4.98l30.32-30.32a6.519 6.519 0 0 0 0-9.192z"/>`),
		g.Group(children),
	)
}