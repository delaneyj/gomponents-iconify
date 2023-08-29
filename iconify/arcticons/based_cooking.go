package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BasedCooking(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="19.5" cy="20.337" r="14" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="19.43" cy="19.928" r="4.083" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.628 28.831c3.347 2.827 11.01 8.916 11.654 10.445s-.181 2.744-1.868 2.291s-8.165-7.912-11.143-11.204m-1.065-12.462a3.825 3.825 0 0 1 .037 3.142c-1.082 1.961-2.352 4.668-4.202 6.563a4.36 4.36 0 0 1-3.83.874a3.586 3.586 0 0 0-2.547.428a3.693 3.693 0 0 1-3.068.335a2.626 2.626 0 0 1-1.71-2.343c.145-1.377-1.286-2.176-2.045-3.16c-.942-1.223-.172-4.95.335-7.326c.367-1.721 2.483-2.143 4.072-3.57c1.282-1.151 3.618-2.428 4.982-1.376a8.805 8.805 0 0 0 2.659 1.525a3.906 3.906 0 0 1 2.175 1.338c1.084 1.388 2.559 2.136 3.142 3.57Z"/>`),
		g.Group(children),
	)
}