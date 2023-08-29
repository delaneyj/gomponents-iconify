package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ControllerJumpToStart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14.959 4.571L7.756 9.52s-.279.201-.279.481s.279.479.279.479l7.203 4.951c.572.38 1.041.099 1.041-.626V5.196c0-.727-.469-1.008-1.041-.625zM6 4H5c-.553 0-1 .048-1 .6v10.8c0 .552.447.6 1 .6h1c.553 0 1-.048 1-.6V4.6c0-.552-.447-.6-1-.6z"/>`),
		g.Group(children),
	)
}