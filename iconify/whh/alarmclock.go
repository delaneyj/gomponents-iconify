package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Alarmclock(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M983.557 320h-3q-11 1-27-2q-85-145-238-211q-6-23-8.5-36.5t-2-28.5t7.5-23t22-13.5t40-5.5q104 0 177 73.5t73 176.5q0 67-41 70zm-23 256q0 91-35.5 174t-95.5 143t-143 95.5t-174 35.5t-174-35.5t-143-95.5t-95.5-143t-35.5-174t35.5-174t95.5-143t143-95.5t174-35.5t174 35.5t143 95.5t95.5 143t35.5 174zm-448-320q-87 0-160.5 43t-116.5 116.5t-43 160.5t43 160.5t116.5 116.5t160.5 43t160.5-43t116.5-116.5t43-160.5t-43-160.5t-116.5-116.5t-160.5-43zm-441 62q-15 3-27 2h-3q-41-3-41-70q0-103 73-176.5t177-73.5q25 0 40 5.5t22 13.5t7.5 23t-2 28.5t-8.5 36.5q-153 66-238 211z"/>`),
		g.Group(children),
	)
}