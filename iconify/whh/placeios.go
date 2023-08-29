package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Placeios(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M512 1024q-139 0-257-25.5t-186.5-70T0 832q0-48 59-90t161.5-68.5T448 642v-66h128v66q125 5 227.5 31.5T965 742t59 90q0 52-68.5 96.5t-186.5 70t-257 25.5zm64-319v63q0 33-20 80.5T512 896t-44-47.5t-20-80.5v-63q-164 7-274 43T64 832q0 35 60 64.5T287 943t225 17t225-17t163-46.5t60-64.5q0-48-110-84t-274-43zm-64-193q-106 0-181-75t-75-181t75-181T512 0t181 75t75 181t-75 181t-181 75zm24-434q-14-13-43-14.5t-64.5 11T368 112t-37.5 60.5t-11 65T335 280q33 34 133.5-67T536 78zm161 216q-12-12-37-3.5T611.5 323t-32 49t3 37t37 3.5t49-32t32-49T697 294z"/>`),
		g.Group(children),
	)
}