package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UploadToCloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15.213 6.639c-.276 0-.546.025-.809.068C13.748 4.562 11.716 3 9.309 3c-2.939 0-5.32 2.328-5.32 5.199c0 .256.02.508.057.756a3.567 3.567 0 0 0-.429-.027C1.619 8.928 0 10.51 0 12.463S1.619 16 3.617 16H8v-4H5.5L10 7l4.5 5H12v4h3.213C17.856 16 20 13.904 20 11.32c0-2.586-2.144-4.681-4.787-4.681z"/>`),
		g.Group(children),
	)
}