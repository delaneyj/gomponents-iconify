package vs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeTimes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1792 1792"),
		g.Raw(`<path fill="currentColor" d="m1264 421l193 195l194-195q17-17 37.5-23.5t38.5-2t34.5 14.5t27 27t14.5 35.5t-2.5 39T1777 549l-195 193l195 193q17 17 23.5 37.5t2.5 38.5t-14.5 34.5t-27 27t-34.5 14.5t-38.5-3t-37.5-24l-194-192l-193 192q-15 15-33 22t-34 6t-31.5-8.5t-27-19t-19-27t-8.5-31.5t6-34t22-33l192-193l-192-193q-17-17-24-37.5t-3-38.5t14.5-35t27-27.5t34.5-15t38.5 2T1264 421zM920 88v1306q0 32-21.5 54.5T846 1471q-32 0-56-24l-400-399H76q-32 0-54-22T0 972V510q0-32 22-54t54-22h314L790 34q25-25 56-25t52.5 23.5T920 88z"/>`),
		g.Group(children),
	)
}