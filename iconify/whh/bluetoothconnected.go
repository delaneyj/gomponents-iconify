package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bluetoothconnected(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024 612q0 12-7 21q-13 14-33 0l-136-93q-16-9-16-26.5t16.5-30T899 448t59-40l26-17q21-14 33 0q7 9 7 21v200zM848 349L599 512l279 182q18 18 18 43.5T878 781q-13 14-35 17l-347 203q-19 21-46 22q-29 2-48-17t-17-48V651l-152 99q-18 18-43.5 18T146 750t-18-43.5t18-43.5l232-151l-232-151q-18-18-18-43.5t18-43.5t43.5-18t43.5 18l152 98V65q-2-28 17-47t48-17q27 1 46 21l382 225q18 18 18 43.5T878 334q-11 11-30 15zM512 854l205-121l-205-133v254zm0-684v254l205-133zM176 540L40 633q-20 14-33 0q-7-9-7-21V412q0-13 7-21q12-14 33 0q99 65 135 93q17 12 17 29.5T176 540z"/>`),
		g.Group(children),
	)
}