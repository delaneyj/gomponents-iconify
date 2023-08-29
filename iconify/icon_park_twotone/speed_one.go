package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeedOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSpeedOne0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M30.297 18.779s-3.23 9.102-4.764 10.691a4 4 0 0 1-5.754-5.557c1.534-1.59 10.518-5.134 10.518-5.134Z"/><path stroke-linecap="round" d="M38.85 38.85A20.941 20.941 0 0 0 45 24c0-11.598-9.402-21-21-21S3 12.402 3 24c0 5.799 2.35 11.049 6.15 14.85M24 4v4m14.845 3.142l-3.108 2.517m6.785 13.574l-3.897-.9m-33.148.9l3.898-.9m-.22-15.191l3.108 2.517"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSpeedOne0)"/>`),
		g.Group(children),
	)
}