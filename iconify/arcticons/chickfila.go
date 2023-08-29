package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Chickfila(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.951 41.705c7.028-31.638 29.394-25.349 21.526-7.236c-6.03 13.884-18.627 7.853-9.313-3.63m13.24-4.246l-.548 5.867c.97-.841 2.118-1.413 3.378-1.872c4.791-1.744-.61-3.464-2.83-3.995ZM14.283 15.614c-2.492-2.737-6.423-.846-3.036 2.944c2.619 2.93 6.18.507 3.036-2.945Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20.538 11.893c-2.55-5.727-8.161-4.103-5.16 2.648c2.157 4.848 7.794 3.267 5.16-2.648Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.44 12.12c-1.329 7.384-8.462 6.4-7.053-1.346c1.553-8.54 8.322-5.697 7.054 1.347Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.682 11.619c-7.685 6.07-1.162 13.056 6.118 6.734c6.828-5.93 1.984-13.134-6.118-6.734Zm-9.805 14.826c-1.515 2.66 1.123 3.463 2.131 1.55c1.232-2.337-.689-4.08-2.13-1.55Z"/>`),
		g.Group(children),
	)
}