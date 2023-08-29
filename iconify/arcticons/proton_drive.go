package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProtonDrive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<defs><path id="arcticonsProtonDrive0" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 36.103V11.897a4.035 4.035 0 0 1 4.034-4.035h7.567a3.586 3.586 0 0 1 2.122.695l2.532 1.86a3.587 3.587 0 0 0 2.123.695h16.587a4.035 4.035 0 0 1 4.035 4.035v20.956a4.035 4.035 0 0 1-4.035 4.035H8.535A4.035 4.035 0 0 1 4.5 36.103Z"/></defs><use href="#arcticonsProtonDrive0" stroke-linecap="round" stroke-linejoin="round"/><use href="#arcticonsProtonDrive0" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.207 40.138v-21.05a3.138 3.138 0 0 0-3.156-3.139l-15.276.088a3.138 3.138 0 0 1-1.837-.58l-3.402-2.419a3.138 3.138 0 0 0-1.818-.58H4.5"/>`),
		g.Group(children),
	)
}