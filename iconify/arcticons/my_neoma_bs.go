package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MyNeomaBs(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.856 7.49c7.39 7.831 8.247 22.796.03 33.049c19.591-3.406 19.39-28.317-.03-33.048Zm-9.891-.029C-.34 12.864-.34 35.748 19.084 40.421h0c-7.901-9.835-8.24-22.91-.119-32.96Z"/>`),
		g.Group(children),
	)
}