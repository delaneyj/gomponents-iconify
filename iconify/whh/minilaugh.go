package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Minilaugh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 512q0 73-35.5 141.5T889 776t-162.5 87T512 896t-214.5-33T135 776T35.5 653T0 512q0-27 18.5-45.5T64 448h896q27 0 45.5 18.5T1024 512zM512 768q74 0 138-17t108.5-45.5t75-61T881 576H144q17 35 47.5 68t75 61.5T375 751t137 17zm256-448q-53 0-90.5-47T640 160t37.5-113T768 0t90.5 47T896 160t-37.5 113t-90.5 47zm-512 0q-53 0-90.5-47T128 160t37.5-113T256 0t90.5 47T384 160t-37.5 113t-90.5 47z"/>`),
		g.Group(children),
	)
}