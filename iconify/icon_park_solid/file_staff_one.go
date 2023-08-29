package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileStaffOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSFileStaffOne0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M40 23v-9L31 4H10a2 2 0 0 0-2 2v36a2 2 0 0 0 2 2h10"/><circle cx="34" cy="32" r="4" fill="#fff"/><path d="M42 44a8 8 0 1 0-16 0m4-40v10h10"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSFileStaffOne0)"/>`),
		g.Group(children),
	)
}