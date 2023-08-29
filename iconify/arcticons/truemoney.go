package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Truemoney(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.725 33.175H11.447M24.02 35.7l10.178.006c2.389 0 2.508-.48 3.042-1.404l6.17-17.235c.364-1.2-.373-2.478-2.357-2.344h-6.352c-1.904-.037-2.593.28-3.26 2.344L25.1 34.777c-.645 1.486-2.443 1.13-3.034-.63c-1.93-6.393-3.919-12.77-5.895-19.15c-.884-2.676-1.662-2.713-3.646-2.703l-5.953.038c-1.715-.038-2.403.785-1.923 2.664c1.814 5.827 3.546 13.857 4.946 16.873c.524 1.26 1.987 1.905 2.71.608l5.042-14.085h9.122c2.529-.077 2.96.269 3.484 2.326m13.457-3.651c.364-1.2-.373-2.478-2.357-2.344"/>`),
		g.Group(children),
	)
}