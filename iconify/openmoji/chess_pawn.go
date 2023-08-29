package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChessPawn(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#3f3f3f" d="M53.533 60.1H16.928L27.58 48.157l3.22-12.794l10.352 1.716l3.835 11.003L53.533 60.1z"/><path d="M43.531 45.449L53.533 60.1H40.586"/><path fill="#3f3f3f" d="M20.662 29.337h32.256v3.726H20.662z"/><circle cx="36" cy="17" r="10" fill="#3f3f3f"/><path d="M41.381 8.58a9.989 9.989 0 0 1-13.8 13.8a9.994 9.994 0 1 0 13.8-13.8Z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="17" r="10"/><path d="M21.534 29.337a2.019 2.019 0 0 0 0 4m9.266 2.309c0 10.967-11.857 23.574-13.872 24.454H36m5.25-23.203C42.083 47.55 53.133 59.254 55.072 60.1H36m16.047-26.763a2.019 2.019 0 0 0 0-4m-30.513 0h30.513m-30.513 4h30.513"/></g>`),
		g.Group(children),
	)
}