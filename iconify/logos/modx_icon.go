package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ModxIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#00B5DE" d="M256.002 25.523H127.614l-17.788 27.843l92.81 57.233z"/><path fill="#00DECC" d="M109.826 53.366L23.976 0v128.388l30.163 17.015l148.497-34.804z"/><path fill="#FF5529" d="m230.48 136.896l-27.844-17.016l-57.233 92.037l85.077 53.366z"/><path fill="#FF9640" d="m202.636 119.88l-149.27 35.578L0 238.987h128.388l17.015-27.07z"/>`),
		g.Group(children),
	)
}