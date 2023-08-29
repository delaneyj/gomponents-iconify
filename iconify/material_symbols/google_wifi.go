package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoogleWifi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m5 20l-.325-1H4.15q-.875 0-1.463-.625t-.537-1.5l.325-4.925l19.05.05l.325 4.875q.05.875-.537 1.5T19.85 19h-.525L19 20H5ZM2.6 10l.275-4.125q.05-.8.625-1.338T4.875 4h14.25q.8 0 1.375.537t.625 1.338L21.4 10H2.6Z"/>`),
		g.Group(children),
	)
}