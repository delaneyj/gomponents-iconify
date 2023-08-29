package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StackWallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 28.818L25.687 39.11l-3.363-1.94l-.014-.014l-14.433-8.338a7.39 7.39 0 0 1-2.359-2.51A7.584 7.584 0 0 1 4.5 22.502c0-1.22.614-2.23 1.505-2.83a5.32 5.32 0 0 1 .366-.21a3.35 3.35 0 0 1 3.228.096l3.806 2.196l3.355 1.932l.022-.008l8.905 5.137l8.883-5.132l8.93 5.132v.003Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.55 9.35c-1.117-.604-2.293-.648-3.528.075m1.663 19.387l17.812-10.28l-15.947-9.18m-3.528.073c-.044.024.044-.025 0 0l-13.9 7.81l-2.231 1.348l-1.522.875a3.35 3.35 0 0 1 3.228.096l3.828 2.18l3.354 1.94l8.906 5.138"/>`),
		g.Group(children),
	)
}