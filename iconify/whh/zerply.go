package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zerply(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M768 768q0 5-.5 14.5t-3.5 36t-8.5 50.5t-17.5 53.5t-28 51t-42.5 36T608 1024q-63 0-118-20t-95.5-48t-79-56t-86.5-48t-101-20q-77 0-102.5-5T0 800q0-48 54.5-131.5t122-159.5T328 337.5T448 192q-42 21-60.5 29.5T328 243t-72 13q-52 0-90-40t-38-88q0-26 6.5-49t16-36.5t19-24T185 4l7-4q1 3 3.5 7T207 22t19.5 19.5t27 15.5t34.5 7q76 0 140-10t97-22t64-22t51-10q23 0 37 5t19.5 16t6.5 19t1 24q0 35-55.5 112.5t-126 155t-155 175T242 662q57 20 140 64.5t148.5 75T640 832q79 0 124-58z"/>`),
		g.Group(children),
	)
}