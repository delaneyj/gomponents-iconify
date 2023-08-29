package fa

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Truck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1536 1536"),
		g.Raw(`<path fill="currentColor" d="M576 1152q0-52-38-90t-90-38t-90 38t-38 90t38 90t90 38t90-38t38-90zM192 640h384V384H418q-13 0-22 9L201 588q-9 9-9 22v30zm1280 512q0-52-38-90t-90-38t-90 38t-38 90t38 90t90 38t90-38t38-90zM1728 64v1024q0 15-4 26.5t-13.5 18.5t-16.5 11.5t-23.5 6t-22.5 2t-25.5 0t-22.5-.5q0 106-75 181t-181 75t-181-75t-75-181H704q0 106-75 181t-181 75t-181-75t-75-181h-64q-3 0-22.5.5t-25.5 0t-22.5-2t-23.5-6t-16.5-11.5T4 1114.5T0 1088q0-26 19-45t45-19V704q0-8-.5-35t0-38t2.5-34.5t6.5-37t14-30.5t22.5-30l198-198q19-19 50.5-32t58.5-13h160V64q0-26 19-45t45-19h1024q26 0 45 19t19 45z"/>`),
		g.Group(children),
	)
}