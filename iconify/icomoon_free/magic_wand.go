package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MagicWand(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4 3L2 1H1v1l2 2zm1-3h1v2H5zm4 5h2v1H9zm1-3V1H9L7 3l1 1zM0 5h2v1H0zm5 4h1v2H5zM1 9v1h1l2-2l-1-1zm14.781 4.781L5.842 3.842a.752.752 0 0 0-1.061 0l-.939.939a.752.752 0 0 0 0 1.061l9.939 9.939a.752.752 0 0 0 1.061 0l.939-.939a.752.752 0 0 0 0-1.061zM7.5 8.5l-3-3l1-1l3 3l-1 1z"/>`),
		g.Group(children),
	)
}