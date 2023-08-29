package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NestWifiGaleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M4.6 10h14.8l-.275-4H4.875L4.6 10Zm-.45 7h15.7l-.325-5H4.475l-.325 5ZM5 20l-.325-1H4.15q-.875 0-1.463-.625t-.537-1.5l.725-11q.05-.8.625-1.338T4.875 4h14.25q.8 0 1.375.537t.625 1.338l.725 11q.05.875-.537 1.5T19.85 19h-.525L19 20H5Z"/>`),
		g.Group(children),
	)
}