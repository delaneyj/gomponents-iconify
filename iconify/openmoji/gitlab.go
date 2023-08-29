package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gitlab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#f1b31c" d="m12.231 33.089l-4.22 11.054L35.645 64.14L12.231 33.089zm23.816 31.352l27.935-20.7l-3.718-9.948l-24.217 30.648z"/><path fill="#ea5a47" d="m25.496 32.888l-6.231-19.796l-7.536 19.997l13.767-.201zm20.8.402h14.169l-7.436-20.097l-6.733 20.097zm-20.298-.201l9.948 30.548l10.853-30.246l-20.801-.302z"/><path fill="#e27022" d="m12.131 32.888l13.867.201l9.948 30.649l-23.815-30.85zm34.467.201l-9.747 29.845l23.112-29.543l-13.365-.302z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="m8.28 43.382l9.731-29.513a1.1 1.1 0 0 1 2.096.018l5.968 19.206l20.216-.001l5.718-19.233a1.1 1.1 0 0 1 2.094-.021l9.635 28.825a1.23 1.23 0 0 1-.425 1.371l-27.255 20.62L8.725 44.762a1.23 1.23 0 0 1-.445-1.38Z"/><path stroke-miterlimit="10" d="M11.673 33.092h14.952m18.961 0h14.953"/><path stroke-linejoin="round" d="M60.337 33.079L36.245 64.617L11.79 33.079"/><path stroke-linejoin="round" d="M46.782 33.582c-.26.994-10.537 31.035-10.537 31.035c.08-.304-9.935-30.726-9.935-30.726"/></g>`),
		g.Group(children),
	)
}