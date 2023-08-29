package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Table(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M0 93v508c0 25 20 45 45 45h628c25 0 44-20 44-45V93c0-25-19-45-44-45H45C20 48 0 68 0 93zm60 173V137h179v129H60zm209 0V137h180v129H269zm209 0V137h179v129H478zM60 426V297h179v129H60zm209 0V297h180v129H269zm209 0V297h179v129H478zM60 585V455h179v130H60zm209 0V455h180v130H269zm209 0V455h179v130H478z"/>`),
		g.Group(children),
	)
}