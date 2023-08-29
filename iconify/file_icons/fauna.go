package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fauna(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m0 0l256.343 204.562l43.203-57.428l116.124-22.433l-60.675 92.345l-.213 125.226L46.186 512l84.86-136.026l80.537-48.089l-83.806 21.038l-72.39-122.18l146.593 40.525l-160.27-79.189z"/>`),
		g.Group(children),
	)
}