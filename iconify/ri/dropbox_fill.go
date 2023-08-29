package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DropboxFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m17.285 10.668l5.215 3.323l-5.252 3.346L12 13.993l-5.248 3.344L1.5 13.99l5.215-3.323L1.5 7.346L6.752 4L12 7.343L17.248 4L22.5 7.346l-5.215 3.322Zm-.074 0L12 7.348l-5.211 3.32L12 13.988l5.211-3.32ZM6.786 18.446l5.252-3.346l5.252 3.346l-5.252 3.346l-5.252-3.346Z"/>`),
		g.Group(children),
	)
}