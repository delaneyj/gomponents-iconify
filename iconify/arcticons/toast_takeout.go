package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ToastTakeout(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.057 5.5c11.175 0 17.593 5.739 17.744 10.533c.043 1.372-.893 2.259-2.252 2.826c2.856 2.384 2.818 13.974 2.705 17.712c-.113 3.738-1.623 5.474-4.04 5.361s-6.455-1.925-13.968-1.963c-7.513-.038-11.1 2.605-14.158 2.53c-3.058-.076-4.38-5.626-4.38-11.78s.68-10.042 3.587-12.836c-.867-.988-2.42-2.785-2.19-4.53c.49-3.726 5.46-7.853 16.952-7.853Z"/>`),
		g.Group(children),
	)
}