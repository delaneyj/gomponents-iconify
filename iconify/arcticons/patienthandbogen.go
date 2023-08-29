package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Patienthandbogen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 30.937s1.011-2.375 4.606-2.375c4.27 0 8.506 2.953 14.894 1.54l-5.457-14.54c-5.104 0-8.506-3.017-10.496-3.017S24 13.668 24 13.668s-1.557-1.123-3.547-1.123s-5.393 3.017-10.496 3.017L4.5 30.102c6.388 1.413 10.625-1.54 14.894-1.54c3.595 0 4.606 2.375 4.606 2.375Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m41.6 30.384l1.731 3.554C38.204 36.731 34.617 33 29.296 33c-3.621 0-4.189 2.455-5.296 2.455c-1.107 0-1.675-2.455-5.296-2.455c-5.32 0-8.908 3.731-14.036.939L6.4 30.383m13.676-15.126c1.42 1.557 8.474-.53 8.474 1.757c0 2.833-7.631.843-7.631 3.515c0 2.335 6.524.893 6.524 2.964s-4.887.415-4.887 1.795c0 1.509 3.081.899 3.081 1.637c0 .433-1.942.979-1.942 1.54s.722.675.722.675M24 13.668v17.269"/>`),
		g.Group(children),
	)
}