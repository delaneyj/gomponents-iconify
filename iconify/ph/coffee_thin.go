package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoffeeThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M84 56V24a4 4 0 0 1 8 0v32a4 4 0 0 1-8 0Zm36 4a4 4 0 0 0 4-4V24a4 4 0 0 0-8 0v32a4 4 0 0 0 4 4Zm32 0a4 4 0 0 0 4-4V24a4 4 0 0 0-8 0v32a4 4 0 0 0 4 4Zm92 60v8a36 36 0 0 1-36 36h-.41a92.53 92.53 0 0 1-35.76 48H208a4 4 0 0 1 0 8H32a4 4 0 0 1 0-8h36.17A92.34 92.34 0 0 1 28 136V88a4 4 0 0 1 4-4h176a36 36 0 0 1 36 36Zm-40 16V92H36v44a84.28 84.28 0 0 0 48.21 76h71.58A84.28 84.28 0 0 0 204 136Zm32-16a28 28 0 0 0-24-27.71V136a91.75 91.75 0 0 1-2.2 19.94A28 28 0 0 0 236 128Z"/>`),
		g.Group(children),
	)
}